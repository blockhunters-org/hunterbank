package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/loan module sentinel errors
var (
	ErrWrongLoanState = sdkerrors.Register(ModuleName, 2, "wrong loan state")
	ErrDeadline       = sdkerrors.Register(ModuleName, 3, "deadline")
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
)
