-- Verify house_finder:schools on pg

BEGIN;

SELECT pg_catalog.has_schema_privilege('house_finder', 'schools');

ROLLBACK;
