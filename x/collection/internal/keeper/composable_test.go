package keeper

import (
	"testing"

	"github.com/line/link/x/collection/internal/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_Attach(t *testing.T) {
	ctx := cacheKeeper()
	prepareCollectionTokens(ctx, t)

	require.EqualError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID1, defaultTokenID1), types.ErrCannotAttachToItself(types.DefaultCodespace, defaultTokenID1).Error())
	require.EqualError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID1, defaultTokenID6), types.ErrTokenNotExist(types.DefaultCodespace, defaultContractID, defaultTokenID6).Error())
	require.EqualError(t, keeper.Attach(ctx, wrongContractID, addr1, defaultTokenID1, defaultTokenID2), types.ErrTokenNotExist(types.DefaultCodespace, wrongContractID, defaultTokenID2).Error())
	require.EqualError(t, keeper.Attach(ctx, defaultContractID, addr2, defaultTokenID1, defaultTokenID2), types.ErrTokenNotOwnedBy(types.DefaultCodespace, defaultTokenID2, addr2).Error())

	require.EqualError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID6, defaultTokenID1), types.ErrTokenNotExist(types.DefaultCodespace, defaultContractID, defaultTokenID6).Error())
	require.EqualError(t, keeper.Attach(ctx, defaultContractID, addr2, defaultTokenID1, defaultTokenID5), types.ErrTokenNotOwnedBy(types.DefaultCodespace, defaultTokenID1, addr2).Error())

	require.NoError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID1, defaultTokenID2))
	require.EqualError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID3, defaultTokenID2), types.ErrTokenAlreadyAChild(types.DefaultCodespace, defaultTokenID2).Error())
	require.EqualError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID2, defaultTokenID1), types.ErrCannotAttachToADescendant(types.DefaultCodespace, defaultTokenID1, defaultTokenID2).Error())
}

func TestKeeper_AttachFrom(t *testing.T) {
	ctx := cacheKeeper()
	prepareCollectionTokens(ctx, t)

	require.EqualError(t, keeper.AttachFrom(ctx, defaultContractID, addr1, addr2, defaultTokenID1, defaultTokenID2), types.ErrCollectionNotApproved(types.DefaultCodespace, addr1.String(), addr2.String(), defaultContractID).Error())
	prepareProxy(ctx, t)
	require.NoError(t, keeper.AttachFrom(ctx, defaultContractID, addr1, addr2, defaultTokenID1, defaultTokenID2))
}

func TestKeeper_Detach(t *testing.T) {
	ctx := cacheKeeper()
	prepareCollectionTokens(ctx, t)

	require.EqualError(t, keeper.Detach(ctx, defaultContractID, addr1, defaultTokenID6), types.ErrTokenNotExist(types.DefaultCodespace, defaultContractID, defaultTokenID6).Error())
	require.EqualError(t, keeper.Detach(ctx, wrongContractID, addr1, defaultTokenID1), types.ErrTokenNotExist(types.DefaultCodespace, wrongContractID, defaultTokenID1).Error())
	require.EqualError(t, keeper.Detach(ctx, defaultContractID, addr2, defaultTokenID1), types.ErrTokenNotOwnedBy(types.DefaultCodespace, defaultTokenID1, addr2).Error())
	require.EqualError(t, keeper.Detach(ctx, defaultContractID, addr1, defaultTokenID1), types.ErrTokenNotAChild(types.DefaultCodespace, defaultTokenID1).Error())
	require.EqualError(t, keeper.Detach(ctx, defaultContractID, addr1, defaultTokenIDFT), types.ErrTokenNotNFT(types.DefaultCodespace, defaultTokenIDFT).Error())

	require.NoError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID1, defaultTokenID2))
	require.NoError(t, keeper.Detach(ctx, defaultContractID, addr1, defaultTokenID2))
}

func TestKeeper_DetachFrom(t *testing.T) {
	ctx := cacheKeeper()
	prepareCollectionTokens(ctx, t)

	require.EqualError(t, keeper.DetachFrom(ctx, defaultContractID, addr1, addr2, defaultTokenID2), types.ErrCollectionNotApproved(types.DefaultCodespace, addr1.String(), addr2.String(), defaultContractID).Error())
	prepareProxy(ctx, t)
	require.NoError(t, keeper.Attach(ctx, defaultContractID, addr2, defaultTokenID1, defaultTokenID2))
	require.NoError(t, keeper.DetachFrom(ctx, defaultContractID, addr1, addr2, defaultTokenID2))
}

func TestKeeper_RootOf(t *testing.T) {
	ctx := cacheKeeper()
	prepareCollectionTokens(ctx, t)

	_, err := keeper.RootOf(ctx, defaultContractID, defaultTokenID6)
	require.EqualError(t, err, types.ErrTokenNotExist(types.DefaultCodespace, defaultContractID, defaultTokenID6).Error())

	_, err = keeper.RootOf(ctx, wrongContractID, defaultTokenID1)
	require.EqualError(t, err, types.ErrTokenNotExist(types.DefaultCodespace, wrongContractID, defaultTokenID1).Error())

	_, err = keeper.RootOf(ctx, defaultContractID, defaultTokenIDFT)
	require.EqualError(t, err, types.ErrTokenNotNFT(types.DefaultCodespace, defaultTokenIDFT).Error())

	nft, err := keeper.RootOf(ctx, defaultContractID, defaultTokenID1)
	require.NoError(t, err)
	require.Equal(t, nft.GetContractID(), defaultContractID)
	require.Equal(t, nft.GetTokenID(), defaultTokenID1)

	require.NoError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID1, defaultTokenID2))
	require.NoError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID2, defaultTokenID3))

	nft, err = keeper.RootOf(ctx, defaultContractID, defaultTokenID3)
	require.NoError(t, err)
	require.Equal(t, nft.GetContractID(), defaultContractID)
	require.Equal(t, nft.GetTokenID(), defaultTokenID1)
}

func TestKeeper_ParentOf(t *testing.T) {
	ctx := cacheKeeper()
	prepareCollectionTokens(ctx, t)

	_, err := keeper.ParentOf(ctx, defaultContractID, defaultTokenID6)
	require.EqualError(t, err, types.ErrTokenNotExist(types.DefaultCodespace, defaultContractID, defaultTokenID6).Error())

	_, err = keeper.ParentOf(ctx, wrongContractID, defaultTokenID1)
	require.EqualError(t, err, types.ErrTokenNotExist(types.DefaultCodespace, wrongContractID, defaultTokenID1).Error())

	_, err = keeper.ParentOf(ctx, defaultContractID, defaultTokenIDFT)
	require.EqualError(t, err, types.ErrTokenNotNFT(types.DefaultCodespace, defaultTokenIDFT).Error())

	nft, err := keeper.ParentOf(ctx, defaultContractID, defaultTokenID1)
	require.NoError(t, err)
	require.Equal(t, nft, nil)

	require.NoError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID1, defaultTokenID2))
	require.NoError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID2, defaultTokenID3))

	nft, err = keeper.ParentOf(ctx, defaultContractID, defaultTokenID3)
	require.NoError(t, err)
	require.Equal(t, nft.GetContractID(), defaultContractID)
	require.Equal(t, nft.GetTokenID(), defaultTokenID2)
}

func TestKeeper_ChildrenOf(t *testing.T) {
	ctx := cacheKeeper()
	prepareCollectionTokens(ctx, t)

	_, err := keeper.ChildrenOf(ctx, defaultContractID, defaultTokenID6)
	require.EqualError(t, err, types.ErrTokenNotExist(types.DefaultCodespace, defaultContractID, defaultTokenID6).Error())

	_, err = keeper.ChildrenOf(ctx, wrongContractID, defaultTokenID1)
	require.EqualError(t, err, types.ErrTokenNotExist(types.DefaultCodespace, wrongContractID, defaultTokenID1).Error())

	_, err = keeper.ChildrenOf(ctx, defaultContractID, defaultTokenIDFT)
	require.EqualError(t, err, types.ErrTokenNotNFT(types.DefaultCodespace, defaultTokenIDFT).Error())

	tokens, err := keeper.ChildrenOf(ctx, defaultContractID, defaultTokenID1)
	require.NoError(t, err)
	require.Equal(t, len(tokens), 0)

	require.NoError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID1, defaultTokenID2))
	require.NoError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID2, defaultTokenID3))
	require.NoError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID1, defaultTokenID4))

	tokens, err = keeper.ChildrenOf(ctx, defaultContractID, defaultTokenID1)
	require.NoError(t, err)
	require.Equal(t, len(tokens), 2)
	require.Equal(t, tokens[0].GetTokenID(), defaultTokenID2)
	require.Equal(t, tokens[1].GetTokenID(), defaultTokenID4)

	tokens, err = keeper.ChildrenOf(ctx, defaultContractID, defaultTokenID2)
	require.NoError(t, err)
	require.Equal(t, len(tokens), 1)
	require.Equal(t, tokens[0].GetTokenID(), defaultTokenID3)
}

func TestAttachDetachScenario(t *testing.T) {
	ctx := cacheKeeper()
	prepareCollectionTokens(ctx, t)

	//
	// attach success cases
	//

	// attach token1 <- token2 (basic case) : success
	require.NoError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID1, defaultTokenID2))

	// attach token2 <- token3 (attach to a child): success
	require.NoError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID2, defaultTokenID3))

	// attach token1 <- token4 (attach to a root): success
	require.NoError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID1, defaultTokenID4))

	// verify the relations

	// root of token1 is token1
	rootOfToken1, err1 := keeper.RootOf(ctx, defaultContractID, defaultTokenID1)
	require.NoError(t, err1)
	require.Equal(t, rootOfToken1.GetTokenID(), defaultTokenID1)

	// root of token2 is token1
	rootOfToken2, err2 := keeper.RootOf(ctx, defaultContractID, defaultTokenID2)
	require.NoError(t, err2)
	require.Equal(t, rootOfToken2.GetTokenID(), defaultTokenID1)

	// root of token3 is token1
	rootOfToken3, err3 := keeper.RootOf(ctx, defaultContractID, defaultTokenID3)
	require.NoError(t, err3)
	require.Equal(t, rootOfToken3.GetTokenID(), defaultTokenID1)

	// root of token4 is token1
	rootOfToken4, err4 := keeper.RootOf(ctx, defaultContractID, defaultTokenID4)
	require.NoError(t, err4)
	require.Equal(t, rootOfToken4.GetTokenID(), defaultTokenID1)

	// parent of token1 is nil
	parentOfToken1, err5 := keeper.ParentOf(ctx, defaultContractID, defaultTokenID1)
	require.NoError(t, err5)
	require.Nil(t, parentOfToken1)

	// parent of token2 is token1
	parentOfToken2, err6 := keeper.ParentOf(ctx, defaultContractID, defaultTokenID2)
	require.NoError(t, err6)
	require.Equal(t, parentOfToken2.GetTokenID(), defaultTokenID1)

	// parent of token3 is token2
	parentOfToken3, err7 := keeper.ParentOf(ctx, defaultContractID, defaultTokenID3)
	require.NoError(t, err7)
	require.Equal(t, parentOfToken3.GetTokenID(), defaultTokenID2)

	// parent of token4 is token1
	parentOfToken4, err8 := keeper.ParentOf(ctx, defaultContractID, defaultTokenID4)
	require.NoError(t, err8)
	require.Equal(t, parentOfToken4.GetTokenID(), defaultTokenID1)

	// children of token1 are token2, token4
	childrenOfToken1, err9 := keeper.ChildrenOf(ctx, defaultContractID, defaultTokenID1)
	require.NoError(t, err9)
	require.Equal(t, len(childrenOfToken1), 2)
	require.True(t, (childrenOfToken1[0].GetTokenID() == defaultTokenID2 && childrenOfToken1[1].GetTokenID() == defaultTokenID4) || (childrenOfToken1[0].GetTokenID() == defaultTokenID4 && childrenOfToken1[1].GetTokenID() == defaultTokenID2))

	// child of token2 is token3
	childrenOfToken2, err10 := keeper.ChildrenOf(ctx, defaultContractID, defaultTokenID2)
	require.NoError(t, err10)
	require.Equal(t, len(childrenOfToken2), 1)
	require.True(t, childrenOfToken2[0].GetTokenID() == defaultTokenID3)

	// child of token3 is empty
	childrenOfToken3, err11 := keeper.ChildrenOf(ctx, defaultContractID, defaultTokenID3)
	require.NoError(t, err11)
	require.Equal(t, len(childrenOfToken3), 0)

	// child of token4 is empty
	childrenOfToken4, err12 := keeper.ChildrenOf(ctx, defaultContractID, defaultTokenID4)
	require.NoError(t, err12)
	require.Equal(t, len(childrenOfToken4), 0)

	// query failure cases
	_, err := keeper.ParentOf(ctx, defaultContractID, defaultTokenIDFT)
	require.EqualError(t, err, types.ErrTokenNotNFT(types.DefaultCodespace, defaultTokenIDFT).Error())

	//
	// attach error cases
	//

	// attach non-root token : failure
	require.EqualError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID1, defaultTokenID2), types.ErrTokenAlreadyAChild(types.DefaultCodespace, defaultTokenID2).Error())

	// attach non-exist token : failure
	require.EqualError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID1, defaultTokenID8), types.ErrTokenNotExist(types.DefaultCodespace, defaultContractID, defaultTokenID8).Error())
	require.EqualError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID8, defaultTokenID1), types.ErrTokenNotExist(types.DefaultCodespace, defaultContractID, defaultTokenID8).Error())

	// attach non-mine token : failure
	require.EqualError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID1, defaultTokenID5), types.ErrTokenNotOwnedBy(types.DefaultCodespace, defaultTokenID5, addr1).Error())
	require.EqualError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID5, defaultTokenID1), types.ErrTokenNotOwnedBy(types.DefaultCodespace, defaultTokenID5, addr1).Error())

	// attach to itself : failure
	require.EqualError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID1, defaultTokenID1), types.ErrCannotAttachToItself(types.DefaultCodespace, defaultTokenID1).Error())

	// attach to a descendant : failure
	require.EqualError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID2, defaultTokenID1), types.ErrCannotAttachToADescendant(types.DefaultCodespace, defaultTokenID1, defaultTokenID2).Error())
	require.EqualError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID3, defaultTokenID1), types.ErrCannotAttachToADescendant(types.DefaultCodespace, defaultTokenID1, defaultTokenID3).Error())
	require.EqualError(t, keeper.Attach(ctx, defaultContractID, addr1, defaultTokenID4, defaultTokenID1), types.ErrCannotAttachToADescendant(types.DefaultCodespace, defaultTokenID1, defaultTokenID4).Error())

	//
	// detach error cases
	//

	// detach not a child : failure
	require.EqualError(t, keeper.Detach(ctx, defaultContractID, addr1, defaultTokenID1), types.ErrTokenNotAChild(types.DefaultCodespace, defaultTokenID1).Error())

	// detach non-mine token : failure
	require.EqualError(t, keeper.Detach(ctx, defaultContractID, addr1, defaultTokenID5), types.ErrTokenNotOwnedBy(types.DefaultCodespace, defaultTokenID5, addr1).Error())

	// detach non-exist token : failure
	require.EqualError(t, keeper.Detach(ctx, defaultContractID, addr1, defaultTokenID8), types.ErrTokenNotExist(types.DefaultCodespace, defaultContractID, defaultTokenID8).Error())

	//
	// detach success cases
	//

	// detach single child
	require.NoError(t, keeper.Detach(ctx, defaultContractID, addr1, defaultTokenID4))

	// detach a child having child
	require.NoError(t, keeper.Detach(ctx, defaultContractID, addr1, defaultTokenID2))

	// detach child
	require.NoError(t, keeper.Detach(ctx, defaultContractID, addr1, defaultTokenID3))

	//
	// verify the relations
	//
	// parent of token2 is nil
	parentOfToken2, err6 = keeper.ParentOf(ctx, defaultContractID, defaultTokenID2)
	require.NoError(t, err6)
	require.Nil(t, parentOfToken2)

	// parent of token3 is nil
	parentOfToken3, err7 = keeper.ParentOf(ctx, defaultContractID, defaultTokenID3)
	require.NoError(t, err7)
	require.Nil(t, parentOfToken3)

	// parent of token4 is nil
	parentOfToken4, err8 = keeper.ParentOf(ctx, defaultContractID, defaultTokenID4)
	require.NoError(t, err8)
	require.Nil(t, parentOfToken4)

	// children of token1 is empty
	childrenOfToken1, err1 = keeper.ChildrenOf(ctx, defaultContractID, defaultTokenID1)
	require.NoError(t, err1)
	require.Equal(t, len(childrenOfToken1), 0)

	// owner of token3 is addr1
	token3, err13 := keeper.GetToken(ctx, defaultContractID, defaultTokenID3)
	require.NoError(t, err13)

	require.Equal(t, (token3.(types.NFT)).GetOwner(), addr1)
}
