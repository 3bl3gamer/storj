package satellitedb_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"storj.io/storj/internal/testcontext"
	"storj.io/storj/satellite/satellitedb"
)

func TestAddrtoNetworkConversion(t *testing.T) {
	ctx := testcontext.New(t)

	ip := "8.8.8.8:28967"
	network, err := satellitedb.GetNetwork(ctx, ip)
	require.Equal(t, "8.8.8.0", network)
	require.NoError(t, err)

	ipv6 := "[fc00::1:200]:28967"
	network, err = satellitedb.GetNetwork(ctx, ipv6)
	require.Equal(t, "fc00::", network)
	require.NoError(t, err)
}
