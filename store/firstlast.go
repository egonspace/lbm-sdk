package store

import (
	"bytes"

	"github.com/line/lfb-sdk/store/types"
	sdk "github.com/line/lfb-sdk/types"
	sdkkv "github.com/line/lfb-sdk/types/kv"
)

// Gets the first item.
func First(st KVStore, start, end []byte) (kv sdkkv.Pair, ok bool) {
	iter := st.Iterator(start, end)
	if !iter.Valid() {
		return kv, false
	}
	defer iter.Close()

	return sdkkv.Pair{Key: iter.Key(), Value: iter.Value()}, true
}

// Gets the last item.  `end` is exclusive.
func Last(st KVStore, start, end []byte) (kv sdkkv.Pair, ok bool) {
	iter := st.ReverseIterator(end, start)
	if !iter.Valid() {
		if v := st.Get(start, types.GetBytesUnmarshalFunc()).([]byte); v != nil {
			return sdkkv.Pair{Key: sdk.CopyBytes(start), Value: sdk.CopyBytes(v)}, true
		}
		return kv, false
	}
	defer iter.Close()

	if bytes.Equal(iter.Key(), end) {
		// Skip this one, end is exclusive.
		iter.Next()
		if !iter.Valid() {
			return kv, false
		}
	}

	return sdkkv.Pair{Key: iter.Key(), Value: iter.Value()}, true
}
