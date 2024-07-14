-- This is the SQL script that will be used to initialize the database schema.
-- We will evaluate you based on how well you design your database.
-- 1. How you design the tables.
-- 2. How you choose the data types and keys.
-- 3. How you name the fields.
-- In this assignment we will use PostgreSQL as the database.

-- This is test table. Remove this table and replace with your own tables. 
-- CREATE TABLE test (
-- 	id serial PRIMARY KEY,
-- 	name VARCHAR ( 50 ) UNIQUE NOT NULL,
-- );

CREATE TABLE estates (
    id UUID PRIMARY KEY,
    width INTEGER NOT NULL,
    length INTEGER NOT NULL
);

CREATE TABLE trees (
    id UUID PRIMARY KEY,
    estate_id UUID REFERENCES estates(id),
    x INTEGER NOT NULL,
    y INTEGER NOT NULL,
    height INTEGER NOT NULL,
    UNIQUE (estate_id, x, y)
);
