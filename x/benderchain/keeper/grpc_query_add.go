package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/faddat/bender-chain/x/benderchain/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AddAll(c context.Context, req *types.QueryAllAddRequest) (*types.QueryAllAddResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var adds []*types.Add
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	addStore := prefix.NewStore(store, types.KeyPrefix(types.AddKey))

	pageRes, err := query.Paginate(addStore, req.Pagination, func(key []byte, value []byte) error {
		var add types.Add
		if err := k.cdc.UnmarshalBinaryBare(value, &add); err != nil {
			return err
		}

		adds = append(adds, &add)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAddResponse{Add: adds, Pagination: pageRes}, nil
}

func (k Keeper) Add(c context.Context, req *types.QueryGetAddRequest) (*types.QueryGetAddResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var add types.Add
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasAdd(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAddIDBytes(req.Id)), &add)

	return &types.QueryGetAddResponse{Add: &add}, nil
}
