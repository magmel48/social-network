-- name: CreateUser :execresult
INSERT INTO `users` (`first_name`, `last_name`, `password`, `gender`, `birthday`, `biography`)
VALUES (?, ?, ?, ?, ?, ?);

-- name: FindUserByID :one
SELECT `id`, `password`, `first_name`, `last_name`, `gender`, `birthday`, `biography`, `created_at`
FROM `users` WHERE `id` = ? LIMIT 1;

-- name: UpsertCity :execresult
INSERT INTO `cities` (`name`) VALUES (?) ON DUPLICATE KEY UPDATE `name` = `name`;

-- name: FindCityByName :one
SELECT `id`, `name` FROM `cities` WHERE `name` = ?;

-- name: InsertUserCity :exec
INSERT INTO `users_cities` (`user_id`, `city_id`) VALUES (?, ?);
