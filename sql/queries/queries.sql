-- name: CreateUser :execresult
INSERT INTO `users` (`first_name`, `last_name`, `password`, `gender`, `birthday`)
VALUES (?, ?, ?, ?, ?);

-- name: FindUserByID :one
SELECT `first_name`, `last_name`, `gender`, `birthday`, `created_at`
FROM `users` WHERE `id` = ? LIMIT 1;

-- name: FindUserWithCheckingPassword :one
SELECT `first_name`, `last_name`, `gender`, `birthday`, `created_at`
FROM `users` WHERE `id` = ? AND `password` = ? LIMIT 1;
