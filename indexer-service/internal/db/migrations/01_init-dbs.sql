-- Stop on error
\set ON_ERROR_STOP on

-- Create FPL Service Database
CREATE DATABASE fpl;

-- Create SofaScore Service Database
CREATE DATABASE sofascore;

-- Optional: Grant all privileges to your main user 'tactify'
-- (Note: Since 'tactify' is the POSTGRES_USER (superuser), it already has access,
-- but this is useful if you create separate service users later).
GRANT ALL PRIVILEGES ON DATABASE fpl TO tactify;
GRANT ALL PRIVILEGES ON DATABASE sofascore TO tactify;