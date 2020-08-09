-- Revert house_finder:appschema from pg

BEGIN;

DROP TABLE houses;

COMMIT;
