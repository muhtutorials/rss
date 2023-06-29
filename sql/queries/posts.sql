-- name: CreatePost :execresult
INSERT INTO posts (id, feed_id, title, description, url, published_at, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetPostsForUser :many
SELECT posts.* from posts
JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = ?
ORDER BY posts.published_at DESC
LIMIT ?;