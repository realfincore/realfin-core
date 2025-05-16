package keeper

import (
	"realfin/x/realestate/types"
)

var _ types.QueryServer = Keeper{}
