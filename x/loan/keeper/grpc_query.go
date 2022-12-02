package keeper

import (
	"github.com/blockhunters-org/hunterbank/x/loan/types"
)

var _ types.QueryServer = Keeper{}
