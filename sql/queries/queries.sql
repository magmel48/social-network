-- name: CreateUser :execresult
INSERT INTO `users` (`first_name`, `last_name`, `password`, `gender`, `birthday`)
VALUES (?, ?, ?, ?, ?);

-- name: FindUserByID :one
SELECT `id`, `password`, `first_name`, `last_name`, `gender`, `birthday`, `created_at`
FROM `users` WHERE `id` = ? LIMIT 1;

-- name: UpsertCity :execresult
INSERT INTO `cities` (`name`) VALUES (?) ON DUPLICATE KEY UPDATE `name` = `name`;

-- name: UpsertHobby :execresult
INSERT INTO `hobbies` (`name`) VALUES (?) ON DUPLICATE KEY UPDATE `name` = `name`;

-- name: InsertUserHobby :exec
INSERT INTO `users_hobbies` (`user_id`, `hobby_id`) VALUES (?, ?);

-- name: InsertUserCity :exec
INSERT INTO `users_cities` (`user_id`, `city_id`) VALUES (?, ?);
