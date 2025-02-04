// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: feeds.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const getFeedByUrl = `-- name: GetFeedByUrl :one
SELECT id, url
FROM feeds
WHERE url = $1
`

type GetFeedByUrlRow struct {
	ID  uuid.UUID
	Url string
}

func (q *Queries) GetFeedByUrl(ctx context.Context, url string) (GetFeedByUrlRow, error) {
	row := q.db.QueryRowContext(ctx, getFeedByUrl, url)
	var i GetFeedByUrlRow
	err := row.Scan(&i.ID, &i.Url)
	return i, err
}

const getFeeds = `-- name: GetFeeds :many
SELECT f.name, f.url, u.name AS created_by
FROM feeds f 
JOIN users u on u.id = f.user_id
`

type GetFeedsRow struct {
	Name      string
	Url       string
	CreatedBy string
}

func (q *Queries) GetFeeds(ctx context.Context) ([]GetFeedsRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedsRow
	for rows.Next() {
		var i GetFeedsRow
		if err := rows.Scan(&i.Name, &i.Url, &i.CreatedBy); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertFeed = `-- name: InsertFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING id, created_at, updated_at, name, url, user_id
`

type InsertFeedParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
}

func (q *Queries) InsertFeed(ctx context.Context, arg InsertFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, insertFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
	)
	return i, err
}
