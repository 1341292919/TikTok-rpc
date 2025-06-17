CREATE TABLE tiktok.`video`
(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_id` bigint NOT NULL ,
    `video_url` varchar(255) NOT NULL ,
    `cover_url` varchar(255) NOT NULL ,
    `title` varchar(255) NOT NULL ,
    `description` varchar(255)NOT NULL ,
    `visit_count` bigint NOT NULL DEFAULT 0,
    `like_count` bigint NOT NULL DEFAULT 0,
    `comment_count` bigint NOT NULL DEFAULT 0,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL ON UPDATE current_timestamp DEFAULT current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    CONSTRAINT `id` PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=20000 DEFAULT CHARSET=utf8mb4;
