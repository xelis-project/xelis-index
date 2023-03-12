CREATE OR REPLACE FUNCTION get_txs()
RETURNS TABLE (
    "hash" VARCHAR,
    "fee" BIGINT,
    "nonce" BIGINT,
    "owner" VARCHAR,
    "signature" VARCHAR,
    "timestamp" BIGINT,
    "transfer_count" BIGINT
)
AS $$
BEGIN
    RETURN QUERY 
      SELECT t."hash", t."fee", t."nonce", t."owner", t."signature", b."timestamp", count(tt."tx_hash") as "transfer_count" 
      FROM "transactions" t
      JOIN "transaction_transfers" tt ON tt."tx_hash" = t."hash" 
      JOIN "transaction_blocks" tb ON tb."tx_hash" = t."hash"
      JOIN "blocks" b ON b."hash" = tb."block_hash"
      GROUP BY t."hash", t."fee", t."nonce", t."owner", t."signature", b."timestamp";
END;
$$ LANGUAGE plpgsql;