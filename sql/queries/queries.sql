-- name: CreateUser :execrows
INSERT INTO `users` (`first_name`, `last_name`, `login`, `password`, `gender`, `birthday`)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: FindUserByID :one
SELECT `first_name`, `last_name`, `gender`, `birthday`, `created_at`
FROM `users` WHERE id = $1 LIMIT 1;

-- name: FindUserWithCheckingPassword :one
SELECT `first_name`, `last_name`, `gender`, `birthday`, `created_at`
FROM `users` WHERE `id` = $1 AND `password` = $2 LIMIT 1;
