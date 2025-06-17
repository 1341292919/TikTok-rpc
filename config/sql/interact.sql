CREATE TABLE tiktok.`comment`
(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_id` bigint NOT NULL ,
    `content` varchar(255) NOT NULL ,
    `parent_id` bigint NOT NULL,
    `like_count` bigint NOT NULL DEFAULT 0,
    `child_count` bigint NOT NULL DEFAULT 0,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL ON UPDATE current_timestamp DEFAULT current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `type` bigint NOT NULL,-- 0表示视频评论 1表示评论的评论
    CONSTRAINT `id` PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=50000 DEFAULT CHARSET=utf8mb4;


CREATE TABLE  tiktok.`user_likes`
(
    `user_id` bigint NOT NULL,
    `target_id` bigint NOT NULL,
    `liked_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `type` bigint NOT NULL,
    PRIMARY KEY (user_id, target_id),
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    INDEX idx_target_id (target_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
