package keeper

import (
	"realfin/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
