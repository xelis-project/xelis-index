create table if not exists "blocks" (
  "hash" varchar primary key,
  "topoheight" bigint,
  "timestamp" bigint,
  "block_type" varchar,
  "cumulative_difficulty" bigint,
  "supply" bigint,
  "difficulty" bigint,
  "reward" bigint,
  "height" bigint,
  "miner" varchar,
  "nonce" bigint,
  "tips" text[]
);

alter table "blocks" enable row level security;
create policy "read_blocks" on "blocks" for select using (true);

create table if not exists "transactions" (
  "hash" varchar primary key,
  "fee" bigint,
  "nonce" bigint,
  "owner" varchar,
  "signature" varchar
);

alter table "transactions" enable row level security;
create policy "read_transactions" on "transactions" for select using (true);

create table if not exists "transaction_transfers" (
  "index" int,
  "tx_hash" varchar,
  "amount" bigint,
  "asset" varchar,
  "to" varchar,
  "extra_data" jsonb,
  primary key ("index", "tx_hash")
);

alter table "transaction_transfers" enable row level security;
create policy "read_transaction_transfers" on "transaction_transfers" for select using (true);

create table if not exists "transaction_blocks" (
  "tx_hash" varchar primary key,
  "block_hash" varchar
);

alter table "transaction_blocks" enable row level security;
create policy "read_transaction_blocks" on "transaction_blocks" for select using (true);
