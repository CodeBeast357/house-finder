-- Revert house_finder:schools from pg

BEGIN;

DROP TABLE IF EXISTS "schools" CASCADE;
DELETE FROM geometry_columns WHERE f_table_name = 'schools' AND f_table_schema = 'house_finder';

COMMIT;
