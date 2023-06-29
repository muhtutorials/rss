-- name: CreateFeedFollow :execresult
INSERT INTO feed_follows (id, user_id, feed_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);

-- name: GetFeedFollowByID :one
SELECT * FROM feed_follows WHERE id = ?;

-- name: GetUserFeedFollows :many
SELECT * FROM feed_follows WHERE user_id = ?;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE id = ? AND user_id = ?;