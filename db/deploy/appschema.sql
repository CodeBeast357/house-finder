-- Deploy house_finder:appschema to pg

BEGIN;

CREATE TABLE IF NOT EXISTS houses (
    id SERIAL,
    address TEXT UNIQUE NOT NULL,
    arrondissement TEXT NOT NULL,
    price INT NOT NULL,
    link TEXT NOT NULL,
    thumbnail_link TEXT NOT NULL,
    provider_name TEXT NOT NULL,
    coordinates GEOMETRY NOT NULL,
    is_black_listed BOOLEAN DEFAULT false,
    is_favorite BOOLEAN DEFAULT false,
    is_in_sweet_spot BOOLEAN NOT NULL,
    creation_datetime TIMESTAMP DEFAULT now()
);
COMMIT;
