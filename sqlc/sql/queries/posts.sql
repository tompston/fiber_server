-- name: GetPost :one
SELECT post_id, created_at, post_title, post_body, user_id 
FROM posts
WHERE  post_id = $1 LIMIT 1;

-- name: GetPosts :many
SELECT 
     posts.post_id, posts.created_at, posts.post_title, posts.post_body, posts.user_id,
     users.username
FROM posts
INNER JOIN users
ON posts.user_id = users.user_id
ORDER BY posts.created_at DESC
LIMIT $1 OFFSET $2;

-- name: CreatePost :one
INSERT INTO posts   ( post_title, post_body, user_id  ) 
VALUES              ( $1, $2, $3 )
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE post_id = $1;

-- name: UpdatePostTitle :one
UPDATE posts 
SET post_title = $1
WHERE post_id = $2
RETURNING posts.post_id, posts.updated_at, posts.post_title;

-- name: UpdatePostBody :one
UPDATE posts 
SET post_body = $1
WHERE post_id = $2
RETURNING posts.post_id, posts.updated_at, posts.post_title;

-- name: GetPostsFromUser :many
SELECT
     posts.post_id, posts.created_at, posts.post_title, posts.post_body, posts.user_id,
     users.username
FROM posts
INNER JOIN users
ON posts.user_id = users.user_id
WHERE users.username = $1;
