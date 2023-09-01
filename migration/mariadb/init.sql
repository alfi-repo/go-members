CREATE TABLE `members`  (
    `id` uuid NOT NULL,
    `name` varchar(255) NOT NULL,
    `verified_at` datetime NULL DEFAULT NULL,
    `created_at` timestamp NOT NULL,
    `updated_at` timestamp NOT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;
