package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"recipes/x/recipes/types"
)

func (k msgServer) CreateRecipe(goCtx context.Context, msg *types.MsgCreateRecipe) (*types.MsgCreateRecipeResponse, error) {
	// Get the context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create variable of type Recipe
	var recipe = types.Recipe{
		Creator: msg.Creator,
		Data:    msg.Data,
		Meta:    msg.Meta,
	}

	// Add a recipe to the store and get back the ID
	id := k.AppendRecipe(ctx, recipe)

	// Return the ID of the recipe
	return &types.MsgCreateRecipeResponse{Id: id}, nil
}
