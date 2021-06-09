package cachekv_test

import (
	"crypto/rand"
	"sort"
	"testing"

	"github.com/line/lfb-sdk/codec"
	"github.com/line/lfb-sdk/testutil/store"
	types2 "github.com/line/lfb-sdk/x/auth/types"
	"github.com/line/tm-db/v2/memdb"

	"github.com/line/lfb-sdk/store/cachekv"
	"github.com/line/lfb-sdk/store/dbadapter"
)

func benchmarkCacheKVStoreIterator(numKVs int, b *testing.B) {
	mem := dbadapter.Store{DB: memdb.NewDB()}
	cdc := codec.NewProtoCodec(store.CreateTestInterfaceRegistry())
	cstore := cachekv.NewStore(mem)
	keys := make([]string, numKVs)

	for i := 0; i < numKVs; i++ {
		key := make([]byte, 32)
		value := store.ValFmt(i)

		_, _ = rand.Read(key)

		keys[i] = string(key)
		cstore.Set(key, value, types2.GetAccountMarshalFunc(cdc))
	}

	sort.Strings(keys)

	for n := 0; n < b.N; n++ {
		iter := cstore.Iterator([]byte(keys[0]), []byte(keys[numKVs-1]))

		for _ = iter.Key(); iter.Valid(); iter.Next() {
		}

		iter.Close()
	}
}

func BenchmarkCacheKVStoreIterator500(b *testing.B)    { benchmarkCacheKVStoreIterator(500, b) }
func BenchmarkCacheKVStoreIterator1000(b *testing.B)   { benchmarkCacheKVStoreIterator(1000, b) }
func BenchmarkCacheKVStoreIterator10000(b *testing.B)  { benchmarkCacheKVStoreIterator(10000, b) }
func BenchmarkCacheKVStoreIterator50000(b *testing.B)  { benchmarkCacheKVStoreIterator(50000, b) }
func BenchmarkCacheKVStoreIterator100000(b *testing.B) { benchmarkCacheKVStoreIterator(100000, b) }
