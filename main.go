package main

import (
	"fmt"
	"log"

	"github.com/xelis-project/xelis-index/indexer"
	"github.com/xelis-project/xelis-index/statements"

	xelisDaemon "github.com/xelis-project/xelis-go-sdk/daemon"
)

func main() {
	instance, err := indexer.LoadInstance(".env")
	if err != nil {
		log.Fatal(err)
	}

	defer instance.Close()

	var blockScan = indexer.BlockScan{Daemon: instance.Daemon}

	scanHeight, err := instance.GetScanHeight()
	if err != nil {
		log.Fatal(err)
	}

	onBlock := func(block *xelisDaemon.Block) error {
		tx, err := instance.DB.Begin()
		if err != nil {
			return err
		}

		_, err = statements.ExecInsertUpdateBlock(tx, block)
		if err != nil {
			return err
		}

		err = tx.Commit()
		if err != nil {
			return err
		}

		fmt.Println("New Block: ", block.Topoheight)
		return nil
	}

	onTx := func(block *xelisDaemon.Block, transaction *xelisDaemon.Transaction) error {
		tx, err := instance.DB.Begin()
		if err != nil {
			return err
		}

		_, err = statements.ExecInsertUpdateTx(tx, transaction)
		if err != nil {
			return err
		}

		_, err = statements.ExecInsertUpdateTxBlock(tx, transaction.Hash, block.Hash)
		if err != nil {
			return err
		}

		for i, transfer := range transaction.Data.Transfer {
			_, err = statements.ExecInsertUpdateTxTransfer(tx, i, transaction.Hash, transfer)
			if err != nil {
				return err
			}
		}

		err = tx.Commit()
		if err != nil {
			return err
		}

		fmt.Println("New Tx: ", transaction.Hash)
		return nil
	}

	err = blockScan.Scan(&indexer.ScanParams{
		ScanHeight: scanHeight,
		OnBlock:    onBlock,
		OnTx:       onTx,
		OnNextScan: func(height uint64) error {
			return instance.SaveScanHeight(height)
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
