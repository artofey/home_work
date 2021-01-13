CREATE EXTENSION "uuid-ossp";

CREATE DATABASE calendar;

CREATE USER calendar_app WITH ENCRYPTED PASSWORD '12345678';

GRANT ALL PRIVILEGES ON DATABASE calendar TO calendar_app;

-- REVOKE ALL PRIVILEGES ON DATABASE calendar FROM calendar_app;

-- DROP ROLE calendar_app;

-- DROP DATABASE calendar;

-- DROP EXTENSION "uuid-ossp";
