CREATE TABLE `user` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `user_id` bigint(20) NOT NULL ,
                        `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
                        `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
                        `email` varchar(64) COLLATE utf8mb4_general_ci  ,
                        `gender` tinyint(4) NOT NULL default '0',
                        `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        primary key (`id`),
                        unique key `idx_username` (`username`) using btree ,
                        unique key `idx_user_id` (`user_id`) using btree
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
