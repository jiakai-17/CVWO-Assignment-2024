-- Queries for sqlc to generate Go code.
-- docker run --rm -v "%cd%:/src" -w /src sqlc/sqlc generate
SELECT 1;

-- Returns 1 if the user with the given username exists.
-- name: CheckUserExists :one
SELECT EXISTS
    (SELECT 1 FROM users WHERE LOWER(username) = LOWER($1))
AS is_existing_user;


-- Creates a new user with the given username and password.
-- name: CreateUser :exec
INSERT INTO users (username, password)
VALUES ($1, $2);


-- Returns a username and their password hash.
-- name: GetPasswordHash :one
SELECT username, password
FROM users
WHERE LOWER(username) = LOWER($1);


-- Creates a new thread with the given title, body, and creator. Returns the details of the created thread.
-- name: CreateThread :one
INSERT INTO threads (title, body, creator)
VALUES ($1, $2, $3)
RETURNING id, title, body, creator, created_time, updated_time, num_comments;


-- Returns the details of the thread with the given id, as well as the tags of the thread as an array.
-- name: GetThreadDetails :one
SELECT t.id, t.title, t.body, t.creator, t.created_time, t.updated_time, t.num_comments,
    CASE
    WHEN COUNT(tt.tag_name) > 0 THEN ARRAY_AGG(tt.tag_name ORDER BY tt.tag_name)
        ELSE '{}'::text[]
    END AS tags
FROM threads t
LEFT JOIN thread_tags tt ON t.id = tt.thread_id
WHERE t.id = $1
GROUP BY t.id;


-- Returns the details of all threads.
-- Sort order should be one of 'created_time_asc', 'created_time_desc', 'num_comments_asc', 'num_comments_desc'.
-- name: GetThreads :many
SELECT t.id, t.title, t.body, t.creator, t.created_time, t.updated_time, t.num_comments,
    CASE
    WHEN COUNT(tt.tag_name) > 0 THEN ARRAY_AGG(tt.tag_name ORDER BY tt.tag_name)
        ELSE '{}'::text[]
    END AS tags
FROM threads t
LEFT JOIN thread_tags tt ON t.id = tt.thread_id
GROUP BY t.id
ORDER BY
    CASE WHEN @sortOrder::text = 'created_time_asc' THEN created_time END ASC,
    CASE WHEN @sortOrder::text = 'created_time_desc' THEN created_time END DESC,
    CASE WHEN @sortOrder::text = 'num_comments_asc' THEN num_comments END ASC,
    CASE WHEN @sortOrder::text = 'num_comments_desc' THEN num_comments END DESC
LIMIT $1
OFFSET $2;


-- Checks if a user is the creator of a thread.
-- name: CheckThreadCreator :one
SELECT EXISTS
    (SELECT 1 FROM threads WHERE id = $1 AND creator = $2)
AS is_thread_creator;


-- Updates the thread with the given id.
-- name: UpdateThread :exec
UPDATE threads
SET title = $1, body = $2, updated_time = NOW()
WHERE id = $3
AND creator = $4;


-- Deletes the thread with the given id.
-- name: DeleteThread :exec
DELETE FROM threads
WHERE id = $1
AND creator = $2;


-- Returns the tags of the thread with the given id.
-- name: GetThreadTags :many
SELECT tag_name
FROM thread_tags
WHERE thread_id = $1;


-- Deletes all tags of the thread with the given id.
-- name: DeleteThreadTags :exec
DELETE FROM thread_tags
WHERE thread_id = $1;


-- Deletes tags that are not associated with any threads.
-- name: DeleteUnusedTags :exec
DELETE FROM tags
WHERE name NOT IN (
    SELECT DISTINCT tag_name
    FROM thread_tags
);


-- Adds new tags to the database if they do not already exist.
-- name: AddNewTags :exec
INSERT INTO tags (name)
SELECT new_tags FROM UNNEST(@tagArray::text[]) AS new_tags
WHERE new_tags NOT IN (
    SELECT name
    FROM tags
    WHERE name = new_tags)
ON CONFLICT DO NOTHING;


-- Adds tags to a thread if they do not already exist.
-- name: AddThreadTags :exec
INSERT INTO thread_tags (thread_id, tag_name)
SELECT $1 as thread_id,
       unnest(@tagArray::text[]) as tag_name
ON CONFLICT DO NOTHING;


-- Returns the threads that match the keywords and tags.
-- If the keyword is provided, only threads that match all the keywords will be returned.
-- If the tags are provided, only threads that match all the tags will be returned.
-- name: GetThreadsByCriteria :many
SELECT t.id, t.title, t.body, t.creator, t.created_time, t.updated_time, t.num_comments,
    -- Concatenate all the tags of the thread into an array.
    CASE
       WHEN COUNT(tt.tag_name) > 0 THEN ARRAY_AGG(tt.tag_name ORDER BY tt.tag_name)
       ELSE '{}'::text[]
    END AS tags
FROM threads t
LEFT JOIN thread_tags tt ON t.id = tt.thread_id
WHERE
    -- Handle the case where the keyword is empty (NULL).
    CASE
        WHEN LENGTH(@keywords::text) > 0 THEN TO_TSVECTOR('simple', t.title || ' ' || t.body) @@ TO_TSQUERY
('simple', @keywords::text)
        ELSE TRUE
    END
AND
    -- Handle the case where the tag array is empty.
    CASE
        WHEN ARRAY_LENGTH(@tagArray::text[], 1) > 0 THEN t.id IN (
            SELECT tt.thread_id
            FROM thread_tags tt
            WHERE tt.tag_name = ANY(@tagArray::text[])
            GROUP BY tt.thread_id
            HAVING COUNT(DISTINCT tt.tag_name) = ARRAY_LENGTH(@tagArray::text[], 1)
        )
        ELSE TRUE
    END
GROUP BY t.id
ORDER BY
    CASE WHEN @sortOrder::text = 'created_time_asc' THEN created_time END ASC,
    CASE WHEN @sortOrder::text = 'created_time_desc' THEN created_time END DESC,
    CASE WHEN @sortOrder::text = 'num_comments_asc' THEN num_comments END ASC,
    CASE WHEN @sortOrder::text = 'num_comments_desc' THEN num_comments END DESC
LIMIT $1
OFFSET $2;

-- Counts the total number of threads that match the keywords and tags.
-- name: GetThreadsByCriteriaCount :one
SELECT COUNT(*) AS total_items
FROM threads t
WHERE
    CASE
        WHEN LENGTH(@keywords::text) > 0 THEN TO_TSVECTOR('simple', t.title || ' ' || t.body) @@ TO_TSQUERY('simple', @keywords::text)
        ELSE TRUE
    END
  AND
    CASE
        WHEN ARRAY_LENGTH(@tagArray::text[], 1) > 0 THEN t.id IN (
          SELECT tt.thread_id
          FROM thread_tags tt
          WHERE tt.tag_name = ANY(@tagArray::text[])
          GROUP BY tt.thread_id
          HAVING COUNT(DISTINCT tt.tag_name) = ARRAY_LENGTH(@tagArray::text[], 1)
        )
        ELSE TRUE
    END;


-- Creates a new comment with the given body, creator, and thread_id. Returns the details of the created comment.
-- name: CreateComment :one
INSERT INTO comments (body, creator, thread_id)
VALUES ($1, $2, $3)
RETURNING id, body, creator, thread_id, created_time, updated_time;


-- Get comments for a thread.
-- Sort order should be one of 'created_time_asc', 'created_time_desc'.
-- name: GetComments :many
SELECT id, body, creator, thread_id, created_time, updated_time
FROM comments
WHERE thread_id = $1
ORDER BY
    CASE WHEN @sortOrder::text = 'created_time_asc' THEN created_time END ASC,
    CASE WHEN @sortOrder::text = 'created_time_desc' THEN created_time END DESC
LIMIT $2
OFFSET $3;


-- Counts the total number of comments for a thread.
-- name: GetCommentCount :one
SELECT COUNT(*) AS total_items
FROM comments
WHERE thread_id = $1;

-- Checks if a user is the creator of a comment.
-- name: CheckCommentCreator :one
SELECT EXISTS
    (SELECT 1 FROM comments WHERE id = $1 AND creator = $2)
AS is_comment_creator;


-- Updates the comment with the given id.
-- name: UpdateComment :exec
UPDATE comments
SET body = $1, updated_time = NOW()
WHERE id = $2
AND creator = $3;


-- Deletes the comment with the given id.
-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1
AND creator = $2;
