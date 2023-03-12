package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xelis-project/xelis-index/indexer"
	"github.com/xelis-project/xelis-index/statements"

	"github.com/urfave/cli/v2"
	xelisDaemon "github.com/xelis-project/xelis-go-sdk/daemon"
)

var instance *indexer.Instance

func index() error {
	var blockScan = indexer.BlockScan{Daemon: instance.Daemon}

	scanHeight, err := instance.GetScanHeight()
	if err != nil {
		return err
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
		return err
	}

	return nil
}

func execsql(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	tx, err := instance.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(string(data))
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:  "xelis-index",
		Usage: "Xelis Index CLI for indexing the blockchain",
		Before: func(ctx *cli.Context) error {
			newInstance, err := indexer.LoadInstance(".env")
			instance = newInstance
			return err
		},
		Commands: []*cli.Command{
			{
				Name:  "index",
				Usage: "scan / listen to blockchain and index the database",
				Action: func(ctx *cli.Context) error {
					return index()
				},
			},
			{
				Name:  "exec-sql",
				Usage: "exec sql file (set seeds or migrations file)",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "filePath",
						Required: true,
					},
				},
				Action: func(ctx *cli.Context) error {
					filePath := ctx.String("filePath")
					return execsql(filePath)
				},
			},
		},
		After: func(ctx *cli.Context) error {
			return instance.Close()
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
