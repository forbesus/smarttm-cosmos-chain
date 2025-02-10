package smarttm_test

import (
	"testing"

	keepertest "smarttm/testutil/keeper"
	"smarttm/testutil/nullify"
	smarttm "smarttm/x/smarttm/module"
	"smarttm/x/smarttm/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SmarttmKeeper(t)
	smarttm.InitGenesis(ctx, k, genesisState)
	got := smarttm.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
