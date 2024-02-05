BEGIN;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id              UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    phone_number    VARCHAR(13)  UNIQUE NOT NULL,
    full_name       VARCHAR(64)  NOT NULL,
    password        VARCHAR(255) NOT NULL,
    created_at      TIMESTAMPTZ  DEFAULT NOW(),
    updated_at      TIMESTAMPTZ  DEFAULT NOW()
);

COMMIT;

CREATE TABLE user_tokens (
    id                  SERIAL      PRIMARY KEY,
    user_id             UUID        UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    success_login_count INT         DEFAULT 0,
    last_login_at       TIMESTAMPTZ DEFAULT NOW()
);
