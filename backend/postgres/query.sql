------- USER QUERIES -----

-- name: CheckUserExists :one
SELECT 1
FROM users
WHERE username = $1;

-- name: CreateUser :exec
INSERT INTO users (username, password)
VALUES ($1, $2);

-- name: GetPasswordHash :one
SELECT password
FROM users
WHERE username = $1;

------- THREAD QUERIES -----

-- name: CreateThread :one
INSERT INTO threads (title, body, creator)
VALUES ($1, $2, $3)
RETURNING id, title, body, creator, created_time, updated_time, num_comments;

-- name: GetThreadDetails :one
SELECT id, title, body, creator, created_time, updated_time, num_comments
FROM threads
WHERE id = $1;

-- name: GetThreads :many
SELECT id, title, body, creator, created_time, updated_time, num_comments
FROM threads
ORDER BY
    CASE WHEN $1 = 'created_time_asc' THEN created_time END ASC,
    CASE WHEN $1 = 'created_time_desc' THEN created_time END DESC,
    CASE WHEN $1 = 'num_comments_asc' THEN num_comments END ASC,
    CASE WHEN $1 = 'num_comments_desc' THEN num_comments END DESC
LIMIT $2
OFFSET $3;

-- name: UpdateThread :exec
UPDATE threads
SET title = $1, body = $2, updated_time = NOW()
WHERE id = $3;

-- name: DeleteThread :exec
DELETE FROM threads
WHERE id = $1;

-- name: GetThreadTags :many
SELECT tag_name
FROM thread_tags
WHERE thread_id = $1;

-- name: AddThreadTag :exec
INSERT INTO thread_tags (thread_id, tag_name)
VALUES ($1, $2);

-- name: DeleteThreadTag :exec
DELETE FROM thread_tags
WHERE thread_id = $1 AND tag_name = $2;

-- name: GetThreadsByMultipleTags :many
SELECT t.id, t.title, t.body, t.creator, t.created_time, t.updated_time, t.num_comments
FROM threads t
INNER JOIN thread_tags tt ON t.id = tt.thread_id
WHERE tt.tag_name IN ($1)
GROUP BY t.id
HAVING COUNT(*) = $2
ORDER BY
    CASE WHEN $3 = 'created_time_asc' THEN created_time END ASC,
    CASE WHEN $3 = 'created_time_desc' THEN created_time END DESC,
    CASE WHEN $3 = 'num_comments_asc' THEN num_comments END ASC,
    CASE WHEN $3 = 'num_comments_desc' THEN num_comments END DESC
LIMIT $4
OFFSET $5;

-- name: GetThreadsByMultipleTagsv2 :many
SELECT t.id, t.title, t.body, t.creator, t.created_time, t.updated_time, t.num_comments
FROM threads t
INNER JOIN thread_tags tt ON t.id = tt.thread_id
WHERE tt.tag_name = ANY($1)
GROUP BY t.id
HAVING COUNT(DISTINCT tt.tag_name) = array_length($1, 1)  -- Count the distinct tags to match all provided tags
ORDER BY
    CASE WHEN $2 = 'created_time_asc' THEN t.created_time END ASC,
    CASE WHEN $2 = 'created_time_desc' THEN t.created_time END DESC,
    CASE WHEN $2 = 'num_comments_asc' THEN t.num_comments END ASC,
    CASE WHEN $2 = 'num_comments_desc' THEN t.num_comments END DESC
LIMIT $3
OFFSET $4;

-- name: GetThreadsByMultipleKeyword :many
-- Join keywords with '&'
SELECT t.id, t.title, t.body, t.creator, t.created_time, t.updated_time, t.num_comments
FROM threads t
WHERE to_tsvector('simple', title || ' ' || body) @@ to_tsquery('simple', $1)
ORDER BY
    CASE WHEN $2 = 'created_time_asc' THEN created_time END ASC,
    CASE WHEN $2 = 'created_time_desc' THEN created_time END DESC,
    CASE WHEN $2 = 'num_comments_asc' THEN num_comments END ASC,
    CASE WHEN $2 = 'num_comments_desc' THEN num_comments END DESC
LIMIT $3
OFFSET $4;


------- COMMENT QUERIES -----

-- name: CreateComment :one
INSERT INTO comments (body, creator, thread_id)
VALUES ($1, $2, $3)
RETURNING id, body, creator, thread_id, created_time, updated_time;

-- name: GetComments :many
SELECT id, body, creator, thread_id, created_time, updated_time
FROM comments
WHERE thread_id = $1
ORDER BY
    CASE WHEN $2 = 'created_time_asc' THEN created_time END ASC,
    CASE WHEN $2 = 'created_time_desc' THEN created_time END DESC
LIMIT $3
OFFSET $4;

-- name: UpdateComment :exec
UPDATE comments
SET body = $1, updated_time = NOW()
WHERE id = $2;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;
