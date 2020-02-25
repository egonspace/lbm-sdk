package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/line/link/x/collection/internal/types"
)

var ChildExists = []byte{1}

type ComposeKeeper interface {
	Attach(ctx sdk.Context, from sdk.AccAddress, symbol string, toTokenID string, tokenID string) sdk.Error
	AttachFrom(ctx sdk.Context, proxy sdk.AccAddress, from sdk.AccAddress, symbol string, toTokenID string, tokenID string) sdk.Error
	Detach(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, symbol string, tokenID string) sdk.Error
	DetachFrom(ctx sdk.Context, proxy sdk.AccAddress, from sdk.AccAddress, to sdk.AccAddress, symbol string, tokenID string) sdk.Error
	RootOf(ctx sdk.Context, symbol string, tokenID string) (types.NFT, sdk.Error)
	ParentOf(ctx sdk.Context, symbol string, tokenID string) (types.NFT, sdk.Error)
	ChildrenOf(ctx sdk.Context, symbol string, tokenID string) (types.Tokens, sdk.Error)
}

func (k Keeper) Attach(ctx sdk.Context, from sdk.AccAddress, symbol string, toTokenID string, tokenID string) sdk.Error {
	if err := k.attach(ctx, from, symbol, toTokenID, tokenID); err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAttachToken,
			sdk.NewAttribute(types.AttributeKeyFrom, from.String()),
			sdk.NewAttribute(types.AttributeKeySymbol, symbol),
			sdk.NewAttribute(types.AttributeKeyToTokenID, toTokenID),
			sdk.NewAttribute(types.AttributeKeyTokenID, tokenID),
		),
	})

	return nil
}

func (k Keeper) AttachFrom(ctx sdk.Context, proxy sdk.AccAddress, from sdk.AccAddress, symbol string, toTokenID string, tokenID string) sdk.Error {
	if !k.IsApproved(ctx, symbol, proxy, from) {
		return types.ErrCollectionNotApproved(types.DefaultCodespace, proxy.String(), from.String(), symbol)
	}

	if err := k.attach(ctx, from, symbol, toTokenID, tokenID); err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAttachFrom,
			sdk.NewAttribute(types.AttributeKeyProxy, proxy.String()),
			sdk.NewAttribute(types.AttributeKeyFrom, from.String()),
			sdk.NewAttribute(types.AttributeKeySymbol, symbol),
			sdk.NewAttribute(types.AttributeKeyToTokenID, toTokenID),
			sdk.NewAttribute(types.AttributeKeyTokenID, tokenID),
		),
	})

	return nil
}

func (k Keeper) attach(ctx sdk.Context, from sdk.AccAddress, symbol string, parentID string, childID string) sdk.Error {
	store := ctx.KVStore(k.storeKey)

	if parentID == childID {
		return types.ErrCannotAttachToItself(types.DefaultCodespace, childID)
	}

	childToken, err := k.GetNFT(ctx, symbol, childID)
	if err != nil {
		return err
	}

	if !from.Equals(childToken.GetOwner()) {
		return types.ErrTokenNotOwnedBy(types.DefaultCodespace, childID, from)
	}

	toToken, err := k.GetNFT(ctx, symbol, parentID)
	if err != nil {
		return err
	}

	if !from.Equals(toToken.GetOwner()) {
		return types.ErrTokenNotOwnedBy(types.DefaultCodespace, parentID, from)
	}

	// verify token should be a root
	childToParentKey := types.TokenChildToParentKey(symbol, childID)
	if store.Has(childToParentKey) {
		return types.ErrTokenAlreadyAChild(types.DefaultCodespace, childID)
	}

	// verify no circulation(toToken must not be a descendant of token)
	rootOfToToken, err := k.RootOf(ctx, symbol, parentID)
	if err != nil {
		return err
	}
	parentToken, err := k.GetNFT(ctx, symbol, parentID)
	if err != nil {
		return err
	}
	if rootOfToToken.GetTokenID() == childID {
		return types.ErrCannotAttachToADescendant(types.DefaultCodespace, childID, parentID)
	}

	parentToChildKey := types.TokenParentToChildKey(symbol, parentID, childID)
	if store.Has(parentToChildKey) {
		panic("token is already a child of some other")
	}

	if !from.Equals(parentToken.GetOwner()) {
		if err := k.moveNFToken(ctx, symbol, from, parentToken.GetOwner(), childToken); err != nil {
			return err
		}
	}

	store.Set(childToParentKey, k.mustEncodeTokenID(parentID))
	store.Set(parentToChildKey, k.mustEncodeTokenID(childID))

	return nil
}

func (k Keeper) Detach(ctx sdk.Context, from sdk.AccAddress, symbol string, tokenID string) sdk.Error {
	parentTokenID, err := k.detach(ctx, from, symbol, tokenID)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDetachToken,
			sdk.NewAttribute(types.AttributeKeyFrom, from.String()),
			sdk.NewAttribute(types.AttributeKeySymbol, symbol),
			sdk.NewAttribute(types.AttributeKeyFromTokenID, parentTokenID),
			sdk.NewAttribute(types.AttributeKeyTokenID, tokenID),
		),
	})
	return nil
}

//nolint:dupl
func (k Keeper) DetachFrom(ctx sdk.Context, proxy sdk.AccAddress, from sdk.AccAddress, symbol string, tokenID string) sdk.Error {
	if !k.IsApproved(ctx, symbol, proxy, from) {
		return types.ErrCollectionNotApproved(types.DefaultCodespace, proxy.String(), from.String(), symbol)
	}

	parentTokenID, err := k.detach(ctx, from, symbol, tokenID)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDetachFrom,
			sdk.NewAttribute(types.AttributeKeyProxy, proxy.String()),
			sdk.NewAttribute(types.AttributeKeyFrom, from.String()),
			sdk.NewAttribute(types.AttributeKeySymbol, symbol),
			sdk.NewAttribute(types.AttributeKeyFromTokenID, parentTokenID),
			sdk.NewAttribute(types.AttributeKeyTokenID, tokenID),
		),
	})

	return nil
}

func (k Keeper) detach(ctx sdk.Context, from sdk.AccAddress, symbol string, childID string) (string, sdk.Error) {
	store := ctx.KVStore(k.storeKey)

	childToken, err := k.GetNFT(ctx, symbol, childID)
	if err != nil {
		return "", err
	}

	if !from.Equals(childToken.GetOwner()) {
		return "", types.ErrTokenNotOwnedBy(types.DefaultCodespace, childID, from)
	}

	childToParentKey := types.TokenChildToParentKey(symbol, childID)
	if !store.Has(childToParentKey) {
		return "", types.ErrTokenNotAChild(types.DefaultCodespace, childID)
	}

	bz := store.Get(childToParentKey)
	parentID := k.mustDecodeTokenID(bz)

	_, err = k.GetNFT(ctx, symbol, parentID)
	if err != nil {
		return "", err
	}

	parentToChildKey := types.TokenParentToChildKey(symbol, parentID, childID)
	if !store.Has(parentToChildKey) {
		panic("token is not a child of some other")
	}

	store.Delete(childToParentKey)
	store.Delete(parentToChildKey)

	return parentID, nil
}

func (k Keeper) RootOf(ctx sdk.Context, symbol string, tokenID string) (types.NFT, sdk.Error) {
	store := ctx.KVStore(k.storeKey)

	token, err := k.GetNFT(ctx, symbol, tokenID)
	if err != nil {
		return nil, err
	}

	for childToParentKey := types.TokenChildToParentKey(symbol, token.GetTokenID()); store.Has(childToParentKey); {
		bz := store.Get(childToParentKey)
		parentTokenID := k.mustDecodeTokenID(bz)

		token, err = k.GetNFT(ctx, symbol, parentTokenID)
		if err != nil {
			return nil, err
		}
		childToParentKey = types.TokenChildToParentKey(symbol, token.GetTokenID())
	}

	return token, nil
}

func (k Keeper) ParentOf(ctx sdk.Context, symbol string, tokenID string) (types.NFT, sdk.Error) {
	store := ctx.KVStore(k.storeKey)

	token, err := k.GetNFT(ctx, symbol, tokenID)
	if err != nil {
		return nil, err
	}
	childToParentKey := types.TokenChildToParentKey(symbol, token.GetTokenID())
	if store.Has(childToParentKey) {
		bz := store.Get(childToParentKey)
		parentTokenID := k.mustDecodeTokenID(bz)

		return k.GetNFT(ctx, symbol, parentTokenID)
	}
	return nil, nil
}

func (k Keeper) ChildrenOf(ctx sdk.Context, symbol string, tokenID string) (types.Tokens, sdk.Error) {
	_, err := k.GetNFT(ctx, symbol, tokenID)
	if err != nil {
		return nil, err
	}
	tokens := k.getChildren(ctx, symbol, tokenID)

	return tokens, nil
}

func (k Keeper) mustEncodeTokenID(tokenID string) []byte {
	return k.cdc.MustMarshalBinaryBare(tokenID)
}

func (k Keeper) mustDecodeTokenID(tokenIDByte []byte) (tokenID string) {
	k.cdc.MustUnmarshalBinaryBare(tokenIDByte, &tokenID)
	return tokenID
}

func (k Keeper) getChildren(ctx sdk.Context, symbol, parentID string) (tokens types.Tokens) {
	getToken := func(tokenID string) bool {
		token, err := k.GetNFT(ctx, symbol, tokenID)
		if err != nil {
			panic(err)
		}
		tokens = append(tokens, token)
		return false
	}

	k.iterateChildren(ctx, symbol, parentID, getToken)

	return tokens
}

func (k Keeper) iterateChildren(ctx sdk.Context, symbol, parentID string, process func(string) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.TokenParentToChildSubKey(symbol, parentID))
	defer iter.Close()
	for {
		if !iter.Valid() {
			return
		}
		val := iter.Value()
		tokenID := k.mustDecodeTokenID(val)
		if process(tokenID) {
			return
		}
		iter.Next()
	}
}
