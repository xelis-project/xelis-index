CREATE OR REPLACE FUNCTION get_blocks()
RETURNS TABLE (
    "hash" VARCHAR,
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
    "size" BIGINT,
    "total_fees" BIGINT,
    "tx_count" BIGINT
)
AS $$
BEGIN
    RETURN QUERY 
        SELECT b."hash", b."topoheight", b."timestamp", b."block_type", b."cumulative_difficulty", b."supply", b."difficulty",
        b."reward", b."height", b."miner", b."nonce", b."tips", b."size", b."total_fees", count(tb."block_hash") as "tx_count" 
        FROM "blocks" b
        LEFT JOIN "transaction_blocks" tb ON tb."block_hash" = b."hash"
        GROUP BY b."hash", b."topoheight", b."timestamp", b."block_type", b."cumulative_difficulty", b."supply", b."difficulty",
        b."reward", b."height", b."miner", b."nonce", b."tips", b."size", b."total_fees"
        ORDER BY "tx_count" DESC;
END;
$$ LANGUAGE plpgsql;
