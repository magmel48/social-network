CREATE TABLE `users` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `first_name` varchar(255) NOT NULL,
    `last_name` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `gender` enum('M', 'F', 'T'),
    `birthday` datetime NOT NULL,
    `biography` text NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` datetime,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
