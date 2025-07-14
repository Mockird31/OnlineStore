CREATE TABLE IF NOT EXISTS "user" (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    CONSTRAINT email_length_check CHECK (LENGTH(email) >= 5 AND LENGTH(email) <= 30),
    username TEXT NOT NULL UNIQUE,
    CONSTRAINT username_length_check CHECK (LENGTH(username) >= 3 AND LENGTH(username) <= 20),
    thumbnail_url TEXT NOT NULL DEFAULT '/default_avatar.png',
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

---- create above / drop below ----

DROP TABLE IF EXISTS "user";