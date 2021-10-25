package simulation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/onomyprotocol/tm-load-test/customclient/simapp"
	sdk "github.com/onomyprotocol/tm-load-test/customclient/types"
	"github.com/onomyprotocol/tm-load-test/customclient/types/kv"
	"github.com/onomyprotocol/tm-load-test/customclient/x/authz"
	"github.com/onomyprotocol/tm-load-test/customclient/x/authz/keeper"
	"github.com/onomyprotocol/tm-load-test/customclient/x/authz/simulation"
	banktypes "github.com/onomyprotocol/tm-load-test/customclient/x/bank/types"
)

func TestDecodeStore(t *testing.T) {
	cdc := simapp.MakeTestEncodingConfig().Codec
	dec := simulation.NewDecodeStore(cdc)

	grant, _ := authz.NewGrant(banktypes.NewSendAuthorization(sdk.NewCoins(sdk.NewInt64Coin("foo", 123))), time.Now().UTC())
	grantBz, err := cdc.Marshal(&grant)
	require.NoError(t, err)
	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{Key: []byte(keeper.GrantKey), Value: grantBz},
			{Key: []byte{0x99}, Value: []byte{0x99}},
		},
	}

	tests := []struct {
		name        string
		expectErr   bool
		expectedLog string
	}{
		{"Grant", false, fmt.Sprintf("%v\n%v", grant, grant)},
		{"other", true, ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectErr {
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			} else {
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
