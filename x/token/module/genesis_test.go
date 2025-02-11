package token_test

import (
	"testing"

	keepertest "smarttm/testutil/keeper"
	"smarttm/testutil/nullify"
	token "smarttm/x/token/module"
	"smarttm/x/token/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TokenKeeper(t)
	token.InitGenesis(ctx, k, genesisState)
	got := token.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
