CREATE TABLE  tiktok.`chat_message`
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