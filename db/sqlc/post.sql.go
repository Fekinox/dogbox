// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: post.sql

package db

import (
	"context"
)

const createPost = `-- name: CreatePost :one
INSERT INTO
  posts (filename, deletion_key, hash)
VALUES
  (
    $1,
    $2,
    $3
  ) RETURNING id, filename, deletion_key, hash, status, created_at, updated_at
`

type CreatePostParams struct {
	Filename *string `json:"filename"`
	Delkey   *string `json:"delkey"`
	Hash     *string `json:"hash"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (*Post, error) {
	row := q.db.QueryRow(ctx, createPost, arg.Filename, arg.Delkey, arg.Hash)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Filename,
		&i.DeletionKey,
		&i.Hash,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE
  id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deletePost, id)
	return err
}

const getAllPosts = `-- name: GetAllPosts :many
SELECT
  id, filename, deletion_key, hash, status, created_at, updated_at
FROM
  posts
WHERE
    pos_by_id (id) > $1::bigint * $2::bigint
AND pos_by_id (id) <= $1::bigint * (1 + $2::bigint)
`

type GetAllPostsParams struct {
	PageSize int64 `json:"page_size"`
	PageNum  int64 `json:"page_num"`
}

func (q *Queries) GetAllPosts(ctx context.Context, arg GetAllPostsParams) ([]*Post, error) {
	rows, err := q.db.Query(ctx, getAllPosts, arg.PageSize, arg.PageNum)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Filename,
			&i.DeletionKey,
			&i.Hash,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPost = `-- name: GetPost :one
SELECT
  id, filename, deletion_key, hash, status, created_at, updated_at
FROM
  posts
WHERE
  id = $1
LIMIT
  1
`

func (q *Queries) GetPost(ctx context.Context, id int64) (*Post, error) {
	row := q.db.QueryRow(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Filename,
		&i.DeletionKey,
		&i.Hash,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getPostByFilename = `-- name: GetPostByFilename :one
SELECT
  id, filename, deletion_key, hash, status, created_at, updated_at
FROM
  posts
WHERE
  filename = $1
LIMIT
  1
`

func (q *Queries) GetPostByFilename(ctx context.Context, filename *string) (*Post, error) {
	row := q.db.QueryRow(ctx, getPostByFilename, filename)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Filename,
		&i.DeletionKey,
		&i.Hash,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET
  filename = coalesce($1, filename),
  deletion_key = coalesce($2, deletion_key),
  hash = coalesce($3, hash),
  status = coalesce($4, status),
  updated_at = now ()
WHERE
  id = $5 RETURNING id, filename, deletion_key, hash, status, created_at, updated_at
`

type UpdatePostParams struct {
	Filename    *string        `json:"filename"`
	DeletionKey *string        `json:"deletion_key"`
	Hash        *string        `json:"hash"`
	Status      NullPostStatus `json:"status"`
	ID          int64          `json:"id"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (*Post, error) {
	row := q.db.QueryRow(ctx, updatePost,
		arg.Filename,
		arg.DeletionKey,
		arg.Hash,
		arg.Status,
		arg.ID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Filename,
		&i.DeletionKey,
		&i.Hash,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
