package block

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	cdc "github.com/line/link/client/rpc/link/block/codec"
	lct "github.com/line/link/client/rpc/link/block/context"
	ltp "github.com/line/link/client/rpc/link/block/proxy"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"sync"
)

type Util struct {
	lcdc    cdc.Codec
	ltmtl   ltp.Tendermint
	lcliCtx lct.CLIContext
}

func NewBlockUtil(ctx context.CLIContext) *Util {
	return &Util{lcdc: cdc.NewCodecWrapper(ctx.Codec), lcliCtx: lct.NewCLIContextWrapper(ctx), ltmtl: ltp.NewTendermintLiteWrapper()}
}

func (u *Util) LatestBlockHeight() (int64, error) {
	node, err := u.lcliCtx.GetNode()
	if err != nil {
		return -1, err
	}

	status, err := node.Status()
	if err != nil {
		return -1, err
	}

	height := status.SyncInfo.LatestBlockHeight
	return height, nil
}

func (u *Util) Indent(res interface{}) ([]byte, error) {
	if u.lcliCtx.Indent() {
		return u.lcdc.MarshalJSONIndent(res)
	}
	return u.lcdc.MarshalJSON(res)
}

func (u *Util) IndentJSON(res interface{}) ([]byte, error) {
	if u.lcliCtx.Indent() {
		return u.lcdc.MarshalJSONIndent(res)
	}
	return u.lcdc.MarshalJSON(res)
}

// inject translated transactions to block data
func (u *Util) InjectByteToJsonTxs(blockResponse []byte, byteTxs [][]byte) (block map[string]interface{}, err error) {
	// load block response as a map
	if err := json.Unmarshal(blockResponse, &block); err != nil {
		return nil, err
	}
	var totalTxJSON []map[string]interface{}
	// load translated txs as a map
	for _, byteTx := range byteTxs {
		var txJSON map[string]interface{}
		if err := json.Unmarshal(byteTx, &txJSON); err != nil {
			return nil, err
		}
		// generate a slice to inject
		totalTxJSON = append(totalTxJSON, txJSON)
	}

	// inject the translated transactions
	block["block"].(map[string]interface{})["data"].(map[string]interface{})["txs"] = totalTxJSON
	return
}

func (u *Util) ValidateBlock(rb *ctypes.ResultBlock) (err error) {
	if !u.lcliCtx.TrustNode() {
		check, err := u.lcliCtx.Verify(rb.Block.Height)
		if err != nil {
			return err
		}
		err = u.ltmtl.ValidateBlockMeta(rb.BlockMeta, check)
		if err != nil {
			return err
		}

		err = u.ltmtl.ValidateBlock(rb.Block, check)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *Util) fetchByBlockHeights(latestBlockHeight *int64, fromBlockHeight *int64, fetchSize *int8) (blockWithRxResultsWrapper *cdc.HasMoreResponseWrapper, err error) {
	if *fromBlockHeight > *latestBlockHeight {
		return nil,
			fmt.Errorf("latestBlockHeight(%d) less than fromBlockHeight(%d)", *latestBlockHeight, *fromBlockHeight)
	}

	fbh := NewFetchInfo(latestBlockHeight, fromBlockHeight, fetchSize)
	blockWithTxResults := make([]*cdc.FetchResult, fbh.fetchItemCnt)
	blockFetchErrors := make([]error, fbh.fetchItemCnt)
	var blockConcurrency sync.WaitGroup
	blockConcurrency.Add(int(fbh.fetchItemCnt))
	for idx, blockHeightCursor := range fbh.fetchItemRange {
		go func(idx int, blockHeightCursor int64) {
			defer blockConcurrency.Done()
			defer func() {
				if err := recover(); err != nil {
					blockFetchErrors[idx] = fmt.Errorf("an error occurred while fetching a block by blockHeight(%d), err(%s)", blockHeightCursor, err)
				}
			}()
			fetchResult, err := u.fetchBlock(blockHeightCursor)
			if err != nil {
				panic(err)
			}
			blockWithTxResults[idx] = fetchResult
		}(idx, blockHeightCursor)
	}
	blockConcurrency.Wait()
	for _, blockFetchErr := range blockFetchErrors {
		if blockFetchErr != nil {
			return nil, blockFetchErr
		}
	}
	blockWithRxResultsWrapper = &cdc.HasMoreResponseWrapper{Items: blockWithTxResults, HasMore: fbh.hasMore}
	return
}

func (u *Util) fetchBlock(fetchBlockHeight int64) (*cdc.FetchResult, error) {
	client, err := u.lcliCtx.GetNode()
	if err != nil {
		panic(err)
	}

	block, err := client.Block(&fetchBlockHeight)
	if err != nil {
		return nil, err
	}
	fetchTxsCnt := len(block.Block.Txs)
	txResults := make([]*ctypes.ResultTx, fetchTxsCnt)
	txFetchErrors := make([]error, fetchTxsCnt)
	err = u.ValidateBlock(block)
	if err != nil {
		return nil, err
	}
	var txConcurrency sync.WaitGroup
	txConcurrency.Add(fetchTxsCnt)
	for idx, tx := range block.Block.Txs {
		go func(idx int, txHash []byte) {
			defer txConcurrency.Done()
			defer func() {
				if err := recover(); err != nil {
					txFetchErrors[idx] = fmt.Errorf("an error occurred while fetching a tx by hash(%x), err(%s)", txHash, err)
				}
			}()
			txResult, err := client.Tx(txHash, u.lcliCtx.TrustNode())
			if err != nil {
				panic(err)
			}
			txResults[idx] = txResult
		}(idx, tx.Hash())
	}
	txConcurrency.Wait()

	for _, txFetchErr := range txFetchErrors {
		if txFetchErr != nil {
			return nil, txFetchErr
		}
	}
	return &cdc.FetchResult{Block: block, TxResults: txResults}, nil
}
