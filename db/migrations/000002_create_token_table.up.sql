CREATE TABLE user_tokens (
    id                  SERIAL      PRIMARY KEY,
    user_id             UUID        UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    success_login_count INT         DEFAULT 0,
    last_login_at       TIMESTAMPTZ DEFAULT NOW()
);
