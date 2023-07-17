CREATE TABLE `users_hobbies` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `user_id` int unsigned NOT NULL,
    `hobby_id` int unsigned NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `users_hobbies_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    CONSTRAINT `users_hobbies_hobby_id_fk` FOREIGN KEY (`hobby_id`) REFERENCES `hobbies` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
