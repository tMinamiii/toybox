-- Create "departments" table
CREATE TABLE `departments` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "user_attributes" table
CREATE TABLE `user_attributes` (`id` bigint NOT NULL AUTO_INCREMENT, `uuid` char(36) NOT NULL, `google_id_token_sub` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "user_auths" table
CREATE TABLE `user_auths` (`id` bigint NOT NULL AUTO_INCREMENT, `uuid` char(36) NOT NULL, `sid` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `uuid` char(36) NOT NULL, `department_id` bigint NOT NULL, `age` bigint NOT NULL, `name` varchar(255) NOT NULL DEFAULT "unknown", `birthday` timestamp NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`), INDEX `user_department_id` (`department_id`), UNIQUE INDEX `uuid` (`uuid`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
