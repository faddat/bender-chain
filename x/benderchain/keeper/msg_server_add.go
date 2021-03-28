package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/faddat/bender-chain/x/benderchain/types"
)

func (k msgServer) CreateAdd(goCtx context.Context, msg *types.MsgCreateAdd) (*types.MsgCreateAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendAdd(
		ctx,
		msg.Creator,
		msg.Post,
		msg.Title,
		msg.Body,
		msg.Ipfs,
	)

	return &types.MsgCreateAddResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateAdd(goCtx context.Context, msg *types.MsgUpdateAdd) (*types.MsgUpdateAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var add = types.Add{
		Creator: msg.Creator,
		Id:      msg.Id,
		Post:    msg.Post,
		Title:   msg.Title,
		Body:    msg.Body,
		Ipfs:    msg.Ipfs,
	}

	// Checks that the element exists
	if !k.HasAdd(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetAddOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetAdd(ctx, add)

	return &types.MsgUpdateAddResponse{}, nil
}

func (k msgServer) DeleteAdd(goCtx context.Context, msg *types.MsgDeleteAdd) (*types.MsgDeleteAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasAdd(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetAddOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAdd(ctx, msg.Id)

	return &types.MsgDeleteAddResponse{}, nil
}
