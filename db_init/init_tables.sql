-- Active: 1726124406018@@127.0.0.1@5432@note_taker
CREATE TABLE users (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    login TEXT NOT NULL UNIQUE, 
    password TEXT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE notes (
    id INT GENERATED ALWAYS AS IDENTITY,
    user_id UUID NOT NULL,
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (id),
    CONSTRAINT fk_user_uuid FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE ON UPDATE CASCADE
);