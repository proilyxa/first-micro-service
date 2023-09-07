CREATE TABLE IF NOT EXISTS users
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name VARCHAR(50)         NOT NULL,
    last_name  VARCHAR(50),
    email      VARCHAR(300) UNIQUE NOT NULL,
    password   VARCHAR(255)        NOT NULL,
    updated_at TEXT                NOT NULL,
    created_at TEXT                NOT NULL
);

-- INSERT into users (id, first_name, last_name, email, password)
-- VALUES (1, 'Ilya', 'Ilya', 'proilyxa@gmail.com', '209011');
