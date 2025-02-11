package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "smarttm/testutil/keeper"
	"smarttm/x/token/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TokenKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
