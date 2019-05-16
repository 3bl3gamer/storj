// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #ifndef UPLINK_HEADERS
//   #define UPLINK_HEADERS
//   #include "c/headers/main.h"
// #endif
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/gogo/protobuf/proto"

	"storj.io/storj/lib/uplink/ext/pb"

	"github.com/zeebo/errs"

	"storj.io/storj/internal/memory"
	"storj.io/storj/pkg/storj"
)

var (
	// cgo types
	cStringType = reflect.TypeOf(C.CString(""))
	cBoolType   = reflect.TypeOf(C.bool(false))
	cIntType    = reflect.TypeOf(C.int(0))
	cUintType   = reflect.TypeOf(C.uint(0))
	// NB: C.uchar is uint8
	cUcharType = reflect.TypeOf(C.uchar('0'))
	// NB: C.long is int64
	cLongType  = reflect.TypeOf(C.long(0))
	cUlongType = reflect.TypeOf(C.ulong(0))

	// our types
	memorySizeType          = reflect.TypeOf(memory.Size(0))
	cipherSuiteType         = reflect.TypeOf(storj.CipherSuite(0))
	redundancyAlgorithmType = reflect.TypeOf(storj.RedundancyAlgorithm(0))
	keyPtrType              = reflect.TypeOf(new(C.Key))
	goValueType             = reflect.TypeOf(C.struct_GoValue{})
	cGoUintptrType           = reflect.TypeOf(C.GoUintptr(0))

	ErrConvert  = errs.Class("struct conversion error")
	ErrSnapshot = errs.Class("unable to snapshot value")

	structRefMap = newMapping()
)

//export GetIDVersion
func GetIDVersion(number C.uint, cErr **C.char) (cIDVersion C.gvIDVersion) {
	goIDVersion, err := storj.GetIDVersion(storj.IDVersionNumber(number))
	if err != nil {
		*cErr = C.CString(err.Error())
		return cIDVersion
	}

	return C.gvIDVersion{
		Ptr:  C.IDVersionRef(structRefMap.Add(goIDVersion)),
		Type: C.IDVersionType,
	}
}

// Create pointer to a go struct for C to interact with
func cPointerFromGoStruct(s interface{}) C.GoUintptr {
	return C.GoUintptr(reflect.ValueOf(s).Pointer())
}

func goPointerFromCGoUintptr(p C.GoUintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p))
}

type GoValue struct {
	ptr      uintptr
	_type    uint
	snapshot []byte
	size     uintptr
}

func (val GoValue) Snapshot() (data []byte, _ error) {
	// TODO: do this using reflect?
	// TODO: use mapping instead of uintptr
	switch val._type {
	case C.IDVersionType:
		idVersion := (*storj.IDVersion)(unsafe.Pointer(val.ptr))
		idVersionPb := pb.IDVersion{
			Number: uint32(idVersion.Number),
		}
		return proto.Marshal(&idVersionPb)
	default:
		// TODO: rename `ErrConvert` to `ErrValue` or something and change message accordingly
		return nil, ErrSnapshot.New("type %s", val._type)
	}
}

//export Unpack
func Unpack(cValue *C.struct_GoValue, cErr **C.char) {
	// TODO: use mapping instead
	value := new(GoValue)
	//value := GoValue{}
	err := CToGoStruct(*cValue, value)
	if err != nil {
		*cErr = C.CString(err.Error())
		return
	}
	data, err := value.Snapshot()
	if err != nil {
		*cErr = C.CString(err.Error())
		return
	}

	value.size = uintptr(len(data))

	value.snapshot = data
	//value.snapshot = CMalloc(uintptr(len(data)))
	//copy(*(*[]byte)(unsafe.Pointer(value.snapshot)), data)
	//snapshotBytes := C.GoBytes(unsafe.Pointer(fromValue.Interface().(*C.Key)), 32)

	if err := GoToCStruct(*value, cValue); err != nil {
		*cErr = C.CString(err.Error())
		return
	}
}

func CMalloc(size uintptr) uintptr {
	CMem := C.malloc(C.ulong(size))
	return uintptr(CMem)
}

func GoToCStruct(fromVar, toPtr interface{}) error {
	fromValue := reflect.ValueOf(fromVar)
	fromKind := fromValue.Kind()
	toPtrValue := reflect.ValueOf(toPtr)

	conversionFunc := fromValue.MethodByName("GoToC")
	if conversionFunc.IsValid() {
		return conversionFunc.Call([]reflect.Value{toPtrValue})[0].Interface().(error)
	}

	toValue := reflect.Indirect(toPtrValue)

	switch fromKind {
	case reflect.String:
		toValue.Set(reflect.ValueOf(C.CString(fromValue.String())))
		return nil
	case reflect.Bool:
		toValue.Set(reflect.ValueOf(C.bool(fromValue.Bool())))
		return nil
	case reflect.Int:
		toValue.Set(reflect.ValueOf(C.int(fromValue.Int())))
		return nil
	case reflect.Uint:
		toValue.Set(reflect.ValueOf(C.uint(fromValue.Uint())))
		return nil
	case reflect.Uint8:
		toValue.Set(reflect.ValueOf(C.uint(fromValue.Uint())))
		return nil
	case reflect.Struct:
		for i := 0; i < fromValue.NumField(); i++ {
			fromFieldValue := fromValue.Field(i)
			fromField := fromValue.Type().Field(i)
			toField := toValue.FieldByName(fromField.Name)
			if toField.CanSet() {
				toFieldPtr := reflect.New(toField.Type())
				toFieldValue := toFieldPtr.Interface()

				// initialize new C value pointer
				if err := GoToCStruct(fromFieldValue.Interface(), toFieldValue); err != nil {
					return err
				}

				// set "to" field value to modified value
				toValue.FieldByName(fromField.Name).Set(reflect.Indirect(toFieldPtr))
			}
		}
		return nil
	default:
		return ErrConvert.New("unsupported kind %s", fromKind)
	}
}

func CToGoStruct(fromVar, toPtr interface{}) error {
	fromValue := reflect.ValueOf(fromVar)
	fromType := fromValue.Type()
	toPtrValue := reflect.ValueOf(toPtr)
	toValue := reflect.Indirect(toPtrValue)

	conversionFunc := toPtrValue.MethodByName("CToGo")
	if conversionFunc.IsValid() {
		result := conversionFunc.Call([]reflect.Value{fromValue})[0].Interface()
		if err, ok := result.(error); ok && err != nil {
			return err
		}
		return nil
	}

	switch fromType {
	case cStringType:
		toValue.Set(reflect.ValueOf(C.GoString(fromValue.Interface().(*C.char))))
		return nil
	case keyPtrType:
		key := new(storj.Key)
		from := C.GoBytes(unsafe.Pointer(fromValue.Interface().(*C.Key)), 32)
		copy(key[:], from)
		toValue.Set(reflect.ValueOf(key))
		return nil
	case cBoolType:
		toValue.Set(reflect.ValueOf(fromValue.Bool()))
		return nil
	case cIntType:
		switch toValue.Kind() {
		case reflect.Int32:
			toValue.Set(reflect.ValueOf(int32(fromValue.Int())))
		default:
			toValue.Set(reflect.ValueOf(int(fromValue.Int())))
		}
		return nil
	case cUintType:
		toValue.Set(reflect.ValueOf(uint(fromValue.Uint())))
		return nil
	case cUcharType:
		switch toValue.Type() {
		case cipherSuiteType:
			toValue.Set(reflect.ValueOf(storj.CipherSuite(fromValue.Uint())))
		case redundancyAlgorithmType:
			toValue.Set(reflect.ValueOf(storj.RedundancyAlgorithm(fromValue.Uint())))
		default:
			toValue.Set(reflect.ValueOf(uint8(fromValue.Uint())))
		}
		return nil
	case cLongType:
		switch toValue.Type() {
		case memorySizeType:
			// TODO: can casting be done with reflection?
			toValue.Set(reflect.ValueOf(memory.Size(fromValue.Int())))
		default:
			toValue.Set(reflect.ValueOf(int64(fromValue.Int())))
		}
		return nil
	case cUlongType:
		switch fromType {
		case cGoUintptrType:
			// TODO: can casting be done with reflection?
			idVersion, ok := structRefMap.Get(token(uintptr(fromValue.Uint()))).(storj.IDVersion)
			if !ok {
				return ErrConvert.New("")
			}
			toValue.Set(reflect.ValueOf(idVersion))
		default:
			toValue.Set(reflect.ValueOf(uint64(fromValue.Uint())))
		}
		return nil
	case goValueType:
		fromSize := uintptr(fromValue.FieldByName("Size").Uint())
		data := C.GoBytes(unsafe.Pointer(fromValue.FieldByName("Snapshot").Pointer()), C.int(fromSize))

		goValue := GoValue{
			ptr:      uintptr(fromValue.FieldByName("Ptr").Uint()),
			_type:    uint(fromValue.FieldByName("Type").Uint()),
			size:     fromSize,
			snapshot: data,
		}
		reflect.Indirect(toValue).Set(reflect.ValueOf(goValue))
		return nil
	default:
		if fromType.Kind() == reflect.Struct {
			for i := 0; i < fromValue.NumField(); i++ {
				fromFieldValue := fromValue.Field(i)
				fromField := fromValue.Type().Field(i)
				toField := toValue.FieldByName(fromField.Name)

				if toField.CanSet() {
					toFieldPtr := reflect.New(toField.Type())
					toFieldValue := toFieldPtr.Interface()

					// initialize new Go value pointer
					if err := CToGoStruct(fromFieldValue.Interface(), toFieldValue); err != nil {
						return err
					}

					// set "to" field value to modified value
					toValue.FieldByName(fromField.Name).Set(reflect.Indirect(toFieldPtr))
				}
			}
			return nil
		}

		return ErrConvert.New("unsupported type %s", fromType)
	}
}
