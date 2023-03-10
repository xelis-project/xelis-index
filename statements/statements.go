package statements

import (
	"encoding/json"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	xelisDaemon "github.com/xelis-project/xelis-go-sdk/daemon"
)

const INSERT_UPDATE_BLOCK = `
	INSERT INTO "blocks" ("hash", "topoheight", "timestamp", "block_type", "cumulative_difficulty",
		"supply", "difficulty", "reward", "height", "miner", "nonce", "tips", "total_fees", "size")
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	ON CONFLICT (hash) DO UPDATE SET
		"topoheight" = EXCLUDED."topoheight",
		"timestamp" = EXCLUDED."timestamp",
		"block_type" = EXCLUDED."block_type",
		"cumulative_difficulty" = EXCLUDED."cumulative_difficulty",
		"supply" = EXCLUDED."supply",
		"difficulty" = EXCLUDED."difficulty",
		"reward" = EXCLUDED."reward",
		"height" = EXCLUDED."height",
		"miner" = EXCLUDED."miner",
		"nonce" = EXCLUDED."nonce",
		"tips" = EXCLUDED."tips",
		"total_fees" = EXCLUDED."total_fees",
		"size" = EXCLUDED."size";
`

func ExecInsertUpdateBlock(tx *pg.Tx, block *xelisDaemon.Block) (orm.Result, error) {
	return tx.Exec(INSERT_UPDATE_BLOCK,
		block.Hash, block.Topoheight, block.Timestamp, block.BlockType,
		block.CumulativeDifficulty, block.Supply, block.Difficulty, block.Reward, block.Height,
		block.Miner, block.Nonce, pg.Array(block.Tips), block.TotalFees, block.TotalSizeInBytes)
}

const INSERT_UPDATE_TX = `
	INSERT INTO "transactions" ("hash", "fee", "nonce", "owner", "signature")
	VALUES (?, ?, ?, ?, ?)
	ON CONFLICT ("hash") DO UPDATE SET
		"fee" = EXCLUDED."fee",
		"nonce" = EXCLUDED."nonce",
		"owner" = EXCLUDED."owner",
		"signature" = EXCLUDED."signature";
`

func ExecInsertUpdateTx(tx *pg.Tx, transaction *xelisDaemon.Transaction) (orm.Result, error) {
	return tx.Exec(INSERT_UPDATE_TX,
		transaction.Hash, transaction.Fee, transaction.Nonce,
		transaction.Owner, transaction.Signature)
}

const INSERT_UPDATE_TX_BLOCK = `
	INSERT INTO "transaction_blocks" ("tx_hash", "block_hash")
	VALUES (?, ?)
	ON CONFLICT ("tx_hash") DO UPDATE SET
		"block_hash" = EXCLUDED."block_hash";
`

func ExecInsertUpdateTxBlock(tx *pg.Tx, txHash string, blockHash string) (orm.Result, error) {
	return tx.Exec(INSERT_UPDATE_TX_BLOCK, txHash, blockHash)
}

const INSERT_UPDATE_TX_TRANSFERS = `
	INSERT INTO "transaction_transfers" ("index", "tx_hash", "amount", "asset", "to", "extra_data")
	VALUES (?, ?, ?, ?, ?, ?)
	ON CONFLICT ("index", "tx_hash") DO UPDATE SET
		"amount" = EXCLUDED."amount",
		"asset" = EXCLUDED."asset",
		"to" = EXCLUDED."to",
		"extra_data" = EXCLUDED."extra_data";
`

func ExecInsertUpdateTxTransfer(tx *pg.Tx, index int, txHash string, transfer xelisDaemon.Transfer) (orm.Result, error) {
	var extraData map[string]interface{}
	json.Unmarshal([]byte(transfer.ExtraData), &extraData)

	return tx.Exec(INSERT_UPDATE_TX_TRANSFERS, index, txHash,
		transfer.Amount, transfer.Asset, transfer.To, &extraData)
}
