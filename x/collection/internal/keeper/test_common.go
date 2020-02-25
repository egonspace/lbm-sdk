package keeper

import (
	"github.com/line/link/x/collection/internal/types"
	"github.com/line/link/x/iam"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestKeeper() (sdk.Context, store.CommitMultiStore, Keeper) {
	keyIam := sdk.NewKVStoreKey(iam.StoreKey)
	keyCollection := sdk.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(keyCollection, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyIam, sdk.StoreTypeIAVL, db)
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	cdc := codec.New()
	types.RegisterCodec(cdc)
	iam.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	cdc.Seal()

	// add keepers
	iamKeeper := iam.NewKeeper(cdc, keyIam)
	keeper := NewKeeper(cdc, iamKeeper.WithPrefix(types.ModuleName), keyCollection)
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())

	return ctx, ms, keeper
}
