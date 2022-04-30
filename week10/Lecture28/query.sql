CREATE TABLE IF NOT EXISTS top_stories
(
    id integer NOT NULL,
    title VARCHAR(256),
    score integer,
    time_stored VARCHAR(256),
    CONSTRAINT top_stories_pkey PRIMARY KEY (id)
);

-- -- name: GetStory :one
-- SELECT * FROM top_stories
-- WHERE id = $1 LIMIT 1;

-- name: ListStories :many
SELECT * FROM top_stories
ORDER BY score DESC;

-- name: CreateStory :exec
INSERT INTO top_stories(
id,title,score,time_stored
) VALUES(
?, ?, ?, ?
) ON DUPLICATE KEY UPDATE time_stored = ?;

-- -- name: UpdateStory :exec
-- UPDATE top_stories SET title = ?,
--   score = ?,
--   time_stored = ?
-- WHERE id = ?;

-- -- name: DeleteStory :exec
-- DELETE FROM top_stories WHERE id = $1;