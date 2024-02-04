BEGIN;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id              UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    phone_number    VARCHAR(13)  NOT NULL,
    full_name       VARCHAR(64)  NOT NULL,
    password        VARCHAR(255) NOT NULL,
    created_at      TIMESTAMPTZ  DEFAULT NOW(),
    updated_at      TIMESTAMPTZ  DEFAULT NOW()
);

COMMIT;
