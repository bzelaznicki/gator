-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    )
RETURNING *
)
SELECT iff.*, 
f.name AS feed_name, 
u.name AS user_name
FROM inserted_feed_follow iff
JOIN users u ON u.id = iff.user_id
JOIN feeds f ON f.id = iff.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT
ff.id as feed_follow_id,
f.name as feed_name,
f.url as feed_url,
u.name as user_name
FROM feed_follows ff
JOIN feeds f ON f.id = ff.feed_id
JOIN users u ON u.id = ff.user_id
WHERE u.id = $1;

-- name: GetSingleFeedForUser :one
SELECT ff.* FROM feed_follows ff
JOIN feeds f on f.id = ff.feed_id
WHERE ff.user_id = $1 AND f.url = $2;

-- name: DeleteFeed :one
DELETE FROM feed_follows
WHERE id = $1 RETURNING *;