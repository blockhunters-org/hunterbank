package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/blockhunters-org/hunterbank/testutil/keeper"
	"github.com/blockhunters-org/hunterbank/x/loan/keeper"
	"github.com/blockhunters-org/hunterbank/x/loan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LoanKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
