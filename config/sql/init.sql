CREATE TABLE casaos.`user`
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

CREATE TABLE casaos.`video`
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

CREATE TABLE casaos.`comment`
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

CREATE TABLE casaos.`user_follows`
(
    `follower_id` bigint NOT NULL ,-- 关注者的用户ID
    `followee_id` bigint NOT NULL ,-- 被关注者的用户ID
    `followed_at` timestamp NOT NULL DEFAULT current_timestamp, -- 关注时间
    PRIMARY KEY (follower_id, followee_id), -- 联合主键，防止重复关注
    FOREIGN KEY (follower_id) REFERENCES user(id) ON DELETE CASCADE, -- 外键约束
    FOREIGN KEY (followee_id) REFERENCES user(id) ON DELETE CASCADE -- 外键约束
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

CREATE TABLE casaos.`user_likes`
(
    `user_id` bigint NOT NULL,
    `target_id` bigint NOT NULL,
    `liked_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `type` bigint NOT NULL,
    PRIMARY KEY (user_id, target_id),
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    INDEX idx_target_id (target_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE casaos. `chat_message`
(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_id` bigint NOT NULL ,
    `target_id` bigint NOT NULL ,
    `content` varchar(255)NOT NULL ,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `status` bigint NOT NULL DEFAULT 0,-- 1 已读 0 未读
    `type` bigint NOT NULL DEFAULT  0, -- 0 私聊信息 1群聊信息
    CONSTRAINT `id` PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=80000 DEFAULT CHARSET=utf8mb4;