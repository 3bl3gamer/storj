package satellitedb

import (
	"context"
	"database/sql"
	"time"

	"github.com/zeebo/errs"
	"storj.io/storj/satellite/marketing"
	dbx "storj.io/storj/satellite/satellitedb/dbx"
)

type offers struct {
	db dbx.Methods
}

// ListAllOffers returns all offers from the db
func (offers *offers) ListAllOffers(ctx context.Context) ([]marketing.Offer, error) {
	offersDbx, err := offers.db.All_Offer(ctx)
	if err != nil {
		return nil, marketing.OffersErr.Wrap(err)
	}

	return offersFromDbx(offersDbx)
}

func (offers *offers) GetOfferByStatusAndType(ctx context.Context, offerStatus marketing.OfferStatus, offerType marketing.OfferType) (*marketing.Offer, error) {
	if offerStatus == 0 || offerType == 0 {
		return nil, errs.New("offer status or type can't be nil")
	}

	offer, err := offers.db.Get_Offer_By_Status_And_Type(ctx, dbx.Offer_Status(int(offerStatus)), dbx.Offer_Type(int(offerType)))
	if err == sql.ErrNoRows {
		return nil, marketing.OffersErr.New("not found %v", offerStatus)
	}
	if err != nil {
		return nil, marketing.OffersErr.Wrap(err)
	}

	return convertDBOffer(offer)
}

func (offers *offers) GetValidOffer(ctx context.Context, offerStatus marketing.OfferStatus, offerType marketing.OfferType) (*marketing.Offer, error) {
	if offerStatus == 0 || offerType == 0 {
		return nil, errs.New("offer status or type can't be nil")
	}

	offer, err := offers.db.Get_Offer_By_Status_And_Type_And_ExpiresAt_GreaterOrEqual(ctx, dbx.Offer_Status(int(offerStatus)), dbx.Offer_Type(int(offerType)), dbx.Offer_ExpiresAt(time.Now()))
	if err == sql.ErrNoRows {
		return nil, marketing.OffersErr.New("offer not found %i", offerStatus)
	}
	if err != nil {
		return nil, marketing.OffersErr.Wrap(err)
	}

	return convertDBOffer(offer)
}

func (offers *offers) 

// Create insert a new offer into the db
func (offers *offers) Create(ctx context.Context, offer *marketing.Offer) (*marketing.Offer, error) {
	createdOffer, err := offers.db.Create_Offer(ctx,
		dbx.Offer_Name(offer.Name),
		dbx.Offer_Description(offer.Description),
		dbx.Offer_Type(int(offer.Type)),
		dbx.Offer_Credit_In_Cents(offer.Credit),
		dbx.Offer_AwardCreditDurationDays(offer.AwardCreditDurationDays),
		dbx.Offer_InviteeCreditDurationDays(offer.InviteeCreditDurationDays),
		dbx.Offer_RedeemableCap(offer.RedeemableCap),
		dbx.Offer_NumRedeemed(offer.NumRedeemed),
		dbx.Offer_OfferDurationDays(offer.OfferDurationDays),
		dbx.Offer_Status(int(offer.Status)),
	)

	if err != nil {
		return nil, marketing.OffersErr.Wrap(err)
	}

	return convertDBOffer(createdOffer)
}

// Update modifies an existing offer
func (offers *offers) Update(ctx context.Context, offer *marketing.Offer) error {
	updateFields := dbx.Offer_Update_Fields{
		Name:                      dbx.Offer_Name(offer.Name),
		Description:               dbx.Offer_Description(offer.Description),
		Type:                      dbx.Offer_Type(int(offer.Type)),
		Credit:                    dbx.Offer_Credit_In_Cents(offer.Credit),
		AwardCreditDurationDays:   dbx.Offer_AwardCreditDurationDays(offer.AwardCreditDurationDays),
		InviteeCreditDurationDays: dbx.Offer_InviteeCreditDurationDays(offer.InviteeCreditDurationDays),
		RedeemableCap:             dbx.Offer_RedeemableCap(offer.RedeemableCap),
		NumRedeemed:               dbx.Offer_NumRedeemed(offer.NumRedeemed),
		OfferDurationDays:         dbx.Offer_OfferDurationDays(offer.OfferDurationDays),
		Status:                    dbx.Offer_Status(int(offer.Status)),
	}

	offerId := dbx.Offer_Id(offer.ID)

	_, err := offers.db.Update_Offer_By_Id(ctx, offerId, updateFields)

	return marketing.OffersErr.Wrap(err)

}

// Delete is a method for deleting offer by Id from the database.
func (offers *offers) Delete(ctx context.Context, id int) error {
	_, err := offers.db.Delete_Offer_By_Id(ctx, dbx.Offer_Id(id))

	return marketing.OffersErr.Wrap(err)
}

func offersFromDbx(offersDbx []*dbx.Offer) ([]marketing.Offer, error) {
	var offers []marketing.Offer
	var errors []error

	for _, offerDbx := range offersDbx {

		offer, err := convertDBOffer(offerDbx)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		offers = append(offers, *offer)
	}

	return offers, errs.Combine(errors...)
}

func convertDBOffer(offerDbx *dbx.Offer) (*marketing.Offer, error) {
	if offerDbx == nil {
		return nil, marketing.OffersErr.New("offerDbx parameter is nil")
	}

	o := marketing.Offer{
		ID:                        offerDbx.Id,
		Name:                      offerDbx.Name,
		Description:               offerDbx.Description,
		Credit:                    offerDbx.Credit,
		RedeemableCap:             offerDbx.RedeemableCap,
		NumRedeemed:               offerDbx.NumRedeemed,
		OfferDurationDays:         offerDbx.OfferDurationDays,
		AwardCreditDurationDays:   offerDbx.AwardCreditDurationDays,
		InviteeCreditDurationDays: offerDbx.InviteeCreditDurationDays,
		CreatedAt:                 offerDbx.CreatedAt,
		Status:                    marketing.OfferStatus(offerDbx.Status),
		Type:                      marketing.OfferType(offerDbx.Type),
	}

	return &o, nil
}