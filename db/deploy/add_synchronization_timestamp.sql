-- Deploy house_finder:add_synchronizatoin_timestamp to pg

BEGIN;

ALTER TABLE houses
ADD COLUMN sync_timestamp TIMESTAMP;

COMMIT;
