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
JOIN users u on u.id = iff.user_id
JOIN feeds f on f.id = iff.feed_id;