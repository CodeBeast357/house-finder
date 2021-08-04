-- Revert house_finder:add_synchronizatoin_timestamp from pg

BEGIN;

ALTER TABLE houses
DROP COLUMN sync_timestamp TIMESTAMP;

COMMIT;
