-- Verify house_finder:add_synchronizatoin_timestamp on pg

BEGIN;

SELECT 1/COUNT(*)
FROM information_schema.columns
WHERE table_name    = 'houses'
  AND column_name   = 'sync_timestamp';

ROLLBACK;
