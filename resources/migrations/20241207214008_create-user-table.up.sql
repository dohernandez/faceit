CREATE TABLE IF NOT EXISTS users
(
    id            UUID PRIMARY KEY,
    first_name    VARCHAR(255)        NOT NULL,
    last_name     VARCHAR(255)        NOT NULL,
    nickname      VARCHAR(120),
    password_hash VARCHAR(128)        NOT NULL, -- length of 128 characters is generally sufficient for most secure password hashing algorithms.
    email         VARCHAR(255) UNIQUE NOT NULL,
    country       CHAR(2),
    created_at    TIMESTAMP           NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP           NOT NULL DEFAULT NOW()
);

-- idx_users_country is an index use to get users by country.
CREATE INDEX IF NOT EXISTS idx_users_country ON users (country);