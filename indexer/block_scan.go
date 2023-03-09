package indexer

import (
	"context"
	"fmt"
	"time"

	xelisDaemon "github.com/xelis-project/xelis-go-sdk/daemon"
)

type BlockScan struct {
	Daemon *xelisDaemon.RPC
}

var tries int

func (bc *BlockScan) maxRetry(err error) bool {
	if err != nil {
		tries++
		fmt.Printf("Retry %d - %v\n", tries, err)

		if tries >= 3 {
			return true
		}
		time.Sleep(1 * time.Second)
	} else {
		tries = 0
	}

	return false
}

type OnBlockFunc func(block *xelisDaemon.Block) error
type OnTxFunc func(block *xelisDaemon.Block, transaction *xelisDaemon.Transaction) error
type OnNextScanFunc func(height uint64) error

type ScanParams struct {
	ScanHeight uint64
	OnBlock    OnBlockFunc
	OnTx       OnTxFunc
	OnNextScan OnNextScanFunc
}

func (bc *BlockScan) Scan(params *ScanParams) error {
	daemon := bc.Daemon
	ctx := context.Background()
	batch := 20
	scanHeight := params.ScanHeight

	var err error

scanLoop:
	for {
	retry_get_stableheight:
		stableHeight, err := daemon.GetStableHeight(ctx)
		if err != nil {
			if bc.maxRetry(err) {
				break scanLoop
			}
			goto retry_get_stableheight
		}

		for i := scanHeight; i < stableHeight; i += uint64(batch) {
			blockStart := i
			blockEnd := i + uint64(batch) - 1
			if blockEnd > stableHeight {
				blockEnd = stableHeight
			}

			scanHeight = blockEnd

		retry_get_blocks:
			blocks, err := daemon.GetBlocks(ctx, &xelisDaemon.GetRangeParams{
				StartTopoheight: blockStart,
				EndTopoheight:   blockEnd,
			})

			if err != nil {
				if bc.maxRetry(err) {
					break scanLoop
				}
				goto retry_get_blocks
			}

			for _, block := range blocks {
			retry_block:
				err = params.OnBlock(&block)
				if err != nil {
					if bc.maxRetry(err) {
						break scanLoop
					}
					goto retry_block
				}

				txCount := len(block.TxsHashes)
				for a := 0; a < txCount; a += batch {
					txEnd := a + batch
					if txEnd > txCount {
						txEnd = txCount
					}

					txHashes := block.TxsHashes[a:txEnd]

				retry_get_transactions:
					txs, err := daemon.GetTransactions(ctx, &xelisDaemon.GetTransactionsParams{
						TxHashes: txHashes,
					})
					if err != nil {
						if bc.maxRetry(err) {
							break scanLoop
						}
						goto retry_get_transactions
					}

					for _, tx := range txs {
					retry_tx:
						err = params.OnTx(&block, &tx)
						if err != nil {
							if bc.maxRetry(err) {
								break scanLoop
							}
							goto retry_tx
						}
					}
				}
			}

		retry_next_scan:
			err = params.OnNextScan(scanHeight)
			if err != nil {
				if bc.maxRetry(err) {
					break scanLoop
				}
				goto retry_next_scan
			}

			time.Sleep(1 * time.Second)
		}

		if scanHeight == stableHeight {
			fmt.Println("Waiting for next block...")
			time.Sleep(30 * time.Second)
		}
	}

	return err
}
