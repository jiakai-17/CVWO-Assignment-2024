-- RESET DATABASE

DROP TABLE IF EXISTS thread_tags;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS threads;
DROP TABLE IF EXISTS users;

-- CREATE TABLES

CREATE TABLE IF NOT EXISTS users (
    username VARCHAR(64) PRIMARY KEY,
    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS threads (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    creator VARCHAR(64) NOT NULL,
    created_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    num_comments INTEGER NOT NULL DEFAULT 0,
    CONSTRAINT fk_creator FOREIGN KEY (creator) REFERENCES users(username) ON DELETE CASCADE,
    CONSTRAINT created_time_not_future CHECK (created_time <= NOW()),
    CONSTRAINT updated_time_not_future CHECK (updated_time <= NOW()),
    CONSTRAINT updated_time_not_before_created_time CHECK (updated_time >= created_time),
    CONSTRAINT num_comments_not_negative CHECK (num_comments >= 0)
);

CREATE TABLE IF NOT EXISTS comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    body TEXT NOT NULL,
    creator VARCHAR(64) NOT NULL,
    thread_id UUID NOT NULL,
    created_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_creator FOREIGN KEY (creator) REFERENCES users(username) ON DELETE CASCADE,
    CONSTRAINT fk_thread FOREIGN KEY (thread_id) REFERENCES threads(id) ON DELETE CASCADE,
    CONSTRAINT created_time_not_future CHECK (created_time <= NOW()),
    CONSTRAINT updated_time_not_future CHECK (updated_time <= NOW()),
    CONSTRAINT updated_time_not_before_created_time CHECK (updated_time >= created_time)
);

CREATE TABLE IF NOT EXISTS tags (
    name VARCHAR(64) PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS thread_tags (
    thread_id UUID NOT NULL,
    tag_name VARCHAR(64) NOT NULL,
    PRIMARY KEY (thread_id, tag_name),
    CONSTRAINT fk_thread FOREIGN KEY (thread_id) REFERENCES threads(id) ON DELETE CASCADE,
    CONSTRAINT fk_tag FOREIGN KEY (tag_name) REFERENCES tags(name) ON DELETE CASCADE
);
