CREATE TABLE IF NOT EXISTS auth_tokens
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id    INT             NOT NULL,
    token      VARCHAR(255)    NOT NULL,
    expires_at TEXT,
    updated_at TEXT            NOT NULL,
    created_at TEXT            NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE INDEX auth_tokens_user_id_idx ON auth_tokens (user_id);
