package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"recipes/x/recipes/types"
)

func (k msgServer) CreateData(goCtx context.Context, msg *types.MsgCreateData) (*types.MsgCreateDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateDataResponse{}, nil
}
