package testutil

import (
	"testing"
	"time"

	"github.com/onomyprotocol/tm-load-test/customclient/testutil/network"
	sdk "github.com/onomyprotocol/tm-load-test/customclient/types"
	"github.com/onomyprotocol/tm-load-test/customclient/x/gov/types"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 1
	suite.Run(t, NewIntegrationTestSuite(cfg))

	genesisState := types.DefaultGenesisState()
	genesisState.DepositParams = types.NewDepositParams(sdk.NewCoins(sdk.NewCoin(cfg.BondDenom, types.DefaultMinDepositTokens)), time.Duration(15)*time.Second)
	genesisState.VotingParams = types.NewVotingParams(time.Duration(5) * time.Second)
	bz, err := cfg.Codec.MarshalJSON(genesisState)
	require.NoError(t, err)
	cfg.GenesisState["gov"] = bz
	suite.Run(t, NewDepositTestSuite(cfg))
}
