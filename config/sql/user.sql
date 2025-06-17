CREATE TABLE  tiktok.`user`
(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `username` varchar(255) NOT NULL ,
    `password` varchar(255) NOT NULL ,
    `avatar_url` varchar(255) NOT NULL DEFAULT "no",
    `opt_secret` varchar(255) NOT NULL DEFAULT '',
    `mfa_status` bigint NOT NULL DEFAULT 0,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL ON UPDATE current_timestamp DEFAULT current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    CONSTRAINT `id` PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4;
