-- +goose Up
-- This will run when applying the migration

CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    slug TEXT UNIQUE NOT NULL,
    image TEXT,
    body TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

-- Optional: Add an index on created_at for faster ordering
CREATE INDEX IF NOT EXISTS idx_posts_created_at ON posts(created_at);

-- +goose Down
-- This will run when rolling back the migration

DROP TABLE IF EXISTS posts;
