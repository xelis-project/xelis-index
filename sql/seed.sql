CREATE TABLE IF NOT EXISTS "blocks" (
  "hash" VARCHAR PRIMARY KEY,
  "topoheight" BIGINT,
  "timestamp" BIGINT,
  "block_type" VARCHAR,
  "cumulative_difficulty" BIGINT,
  "supply" BIGINT,
  "difficulty" BIGINT,
  "reward" BIGINT,
  "height" BIGINT,
  "miner" VARCHAR,
  "nonce" BIGINT,
  "tips" TEXT[],
  "total_fees" BIGINT,
  "size" BIGINT
);

ALTER TABLE "blocks" ENABLE ROW LEVEL SECURITY;
CREATE POLICY "read_blocks" ON "blocks" FOR SELECT USING (true);

CREATE TABLE IF NOT EXISTS "transactions" (
  "hash" VARCHAR PRIMARY KEY,
  "fee" BIGINT,
  "nonce" BIGINT,
  "owner" VARCHAR,
  "signature" VARCHAR
);

ALTER TABLE "transactions" ENABLE ROW LEVEL SECURITY;
CREATE POLICY "read_transactions" ON "transactions" FOR SELECT USING (true);

CREATE TABLE IF NOT EXISTS "transaction_transfers" (
  "index" INT,
  "tx_hash" VARCHAR,
  "amount" BIGINT,
  "asset" VARCHAR,
  "to" VARCHAR,
  "extra_data" JSONB,
  PRIMARY KEY ("index", "tx_hash")
);

ALTER TABLE "transaction_transfers" ENABLE ROW LEVEL SECURITY;
CREATE POLICY "read_transaction_transfers" ON "transaction_transfers" FOR SELECT USING (true);

CREATE TABLE IF NOT EXISTS "transaction_blocks" (
  "tx_hash" VARCHAR PRIMARY KEY,
  "block_hash" VARCHAR
);

ALTER TABLE "transaction_blocks" ENABLE ROW LEVEL SECURITY;
CREATE POLICY "read_transaction_blocks" ON "transaction_blocks" FOR SELECT USING (true);