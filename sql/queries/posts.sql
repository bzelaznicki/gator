-- name: CreatePost :exec
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
) ON CONFLICT(url) DO NOTHING;

-- name: GetPostsForUser :many
SELECT p.*, f.name as feed_name
FROM posts p
JOIN feeds f on p.feed_id = f.id
JOIN feed_follows ff on ff.feed_id = f.id
WHERE ff.user_id = $1
ORDER BY p.published_at desc
LIMIT $2;