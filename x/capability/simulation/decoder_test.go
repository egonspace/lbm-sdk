package simulation_test

import (
	"fmt"
	"testing"

	types2 "github.com/line/lfb-sdk/store/types"
	"github.com/line/lfb-sdk/types/kv"
	"github.com/stretchr/testify/require"

	"github.com/line/lfb-sdk/simapp"
	sdk "github.com/line/lfb-sdk/types"
	"github.com/line/lfb-sdk/x/capability/simulation"
	"github.com/line/lfb-sdk/x/capability/types"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := simapp.MakeCodecs()
	dec := simulation.NewDecodeStore(cdc)

	capOwners := types.CapabilityOwners{
		Owners: []types.Owner{{Module: "transfer", Name: "ports/transfer"}},
	}

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{
				Key:   types.KeyIndex,
				Value: sdk.Uint64ToBigEndian(10),
			},
			{
				Key:   types.KeyPrefixIndexCapability,
				Value: cdc.MustMarshalBinaryBare(&capOwners),
			},
			{
				Key:   []byte{0x99},
				Value: []byte{0x99},
			},
		},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Index", "Index A: 10\nIndex B: 10\n"},
		{"CapabilityOwners", fmt.Sprintf("CapabilityOwners A: %v\nCapabilityOwners B: %v\n", capOwners, capOwners)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			var value interface{}
			if i == len(tests) - 1 {
				require.Panics(t, func () {dec.Unmarshal(kvPairs.Pairs[i].Key)(kvPairs.Pairs[i].Value)}, tt.name)
				value = nil
			} else {
				value = dec.Unmarshal(kvPairs.Pairs[i].Key)(kvPairs.Pairs[i].Value)
			}
			pair := types2.KOPair{
				Key:   kvPairs.Pairs[i].Key,
				Value: value,
			}
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec.LogPair(pair, pair) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec.LogPair(pair, pair), tt.name)
			}
		})
	}
}
