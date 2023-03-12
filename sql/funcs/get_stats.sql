CREATE OR REPLACE FUNCTION get_stats("interval" text)
RETURNS TABLE (
    "time" timestamp without time zone,
    "block_count" bigint,
    "avg_difficulty" numeric,
    "sum_size" numeric,
    "avg_block_size" numeric,
    "tx_count" bigint,
    "avg_block_time" numeric,
    "sum_block_fees" numeric,
    "avg_block_fees" numeric,
    "sum_block_reward" numeric,
    "avg_block_reward" numeric
) AS $$
BEGIN
  RETURN QUERY
    SELECT date_trunc(interval, to_timestamp(b."timestamp" / 1000) AT TIME ZONE 'UTC') AS "time",
          count(b."hash") AS "block_count",
          round(avg(b."difficulty"), 2) AS "avg_difficulty",
          sum(b."size") AS "sum_size",
          round(avg(b."size"), 2) AS "avg_block_size",
          count(t."hash") AS "tx_count",
          round(avg((nb."timestamp"-b."timestamp") / 1000), 2) AS "avg_block_time",
          sum(b."total_fees") AS "sum_block_fees",
          ceil(avg(b."total_fees")) AS "avg_block_fees",
          sum(b."reward") AS "sum_block_reward",
          ceil(avg(b."reward")) AS "avg_block_reward"
    FROM "blocks" b
    LEFT JOIN "blocks" nb on nb."topoheight" = b."topoheight" + 1 -- nb = nextBlock
    LEFT JOIN "transaction_blocks" tb ON tb."block_hash" = b."hash"
    LEFT JOIN "transactions" t ON t."hash" = tb."tx_hash"
    GROUP BY "time"
    ORDER BY "time";
END;
$$ LANGUAGE plpgsql;