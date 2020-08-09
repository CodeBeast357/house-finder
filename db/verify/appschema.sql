-- Verify house_finder:appschema on pg

BEGIN;

SELECT pg_catalog.has_schema_privilege('house_finder', 'houses');

ROLLBACK;
