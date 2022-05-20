-- name: SaveLoginInfo :exec
INSERT INTO login_info (username, password) 
VALUES (?, ?);

-- name: GetUsers :many
SELECT * FROM login_info;

-- name: GetUserTasks :many
SELECT
	user_lists.name, user_tasks.text, user_tasks.completed
FROM
	login_info
INNER JOIN
	user_lists
ON
	login_info.id=user_lists.user_id
INNER JOIN
	user_tasks
ON
	user_lists.id=user_tasks.list_id
WHERE
	login_info.id=?;

-- name: GetUserLists :many
SELECT
    user_lists.id
FROM
	login_info
INNER JOIN
	user_lists
ON
	login_info.id=user_lists.user_id
WHERE
	login_info.id=?;

-- name: DeleteUser :exec
DELETE FROM login_info WHERE id=?;

-- name: GetTasks :many
SELECT * FROM user_tasks WHERE list_id=?;

-- name: GetTask :one
SELECT * FROM user_tasks WHERE id=?;

-- name: CreateTask :exec
INSERT INTO user_tasks (list_id, text, completed)
VALUES (?, ?, ?);

-- name: ToggleTask :exec
UPDATE user_tasks SET completed = ? WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM user_tasks WHERE id=?;

-- name: DeleteList :exec
DELETE FROM user_lists WHERE id=?;

-- name: CreateList :exec
INSERT INTO user_lists (user_id, name) 
VALUES (?,?);

-- name: GetLists :many
SELECT * FROM user_lists WHERE user_id=?;

-- name: GetSpecificUser :one
SELECT * FROM login_info WHERE username=?;