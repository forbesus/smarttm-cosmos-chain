package keeper

import (
	"smarttm/x/token/types"
)

var _ types.QueryServer = Keeper{}
