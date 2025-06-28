-- Create user if not exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_catalog.pg_user WHERE usename = 'ugoapi') THEN
        CREATE USER ugoapi WITH PASSWORD 'pgoapi';
    END IF;
END
$$;

-- Create database if not exists
SELECT 'CREATE DATABASE goapi OWNER ugoapi'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'goapi')\gexec

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE goapi TO ugoapi;
