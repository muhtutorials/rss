-- name: CreateFeed :execresult
INSERT INTO feeds (id, user_id, name, url, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetFeedByID :one
SELECT * FROM feeds WHERE id = ?;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at LIMIT ?;

-- name: MarkFeedFetched :execresult
UPDATE feeds
SET last_fetched_at = CURRENT_TIMESTAMP(), updated_at = CURRENT_TIMESTAMP()
WHERE id = ?;