CREATE TABLE `users` (
    `id` BIGINT,
    `screen_name` VARCHAR(255),
    `observed_time` DATETIME(3),
    PRIMARY KEY (`id`, `screen_name`)
);

CREATE TABLE `micro_blogs` (
    `id` BIGINT,
    `text` TEXT,
    `user_id` BIGINT,
    `created_at` DATETIME(3),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `statistics` (
    `id` BIGINT AUTO_INCREMENT,
    `micro_blog_id` BIGINT,
    `observed_time` DATETIME(3),
    `num_thumb_up` INT,
    `num_comment` INT,
    `num_repost` INT,
    PRIMARY KEY (`id`),
    KEY (`micro_blog_id`, `observed_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;