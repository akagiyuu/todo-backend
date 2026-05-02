-- name: CreateTodo :one
INSERT INTO todos(account_id, title, content, priority)
VALUES (@account_id, @title, @content, @priority)
RETURNING id;

-- name: GetTodo :one
SELECT title, content, priority, is_done, created_at
FROM todos
WHERE id = @id AND account_id = @account_id;

-- name: FilterTodo :many
SELECT id, title, content, priority, is_done, created_at
FROM todos
WHERE account_id = @account_id AND
    (
        sqlc.narg('query')::text IS NULL OR
        title LIKE '%' || @query || '%' OR
        content LIKE '%' || @query || '%'
    ) AND
    (sqlc.narg('priority')::priority IS NULL OR priority = @priority) AND
    (sqlc.narg('is_done')::bool IS NULL OR is_done = @is_done);

-- name: UpdateTodo :exec
UPDATE todos
SET
    title = COALESCE(sqlc.narg('title'), title),
    content = COALESCE(sqlc.narg('content'), content),
    priority = COALESCE(sqlc.narg('priority'), priority)
WHERE id = @id AND account_id = @account_id AND is_done = false;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = @id AND account_id = @account_id;
