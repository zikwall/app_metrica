CREATE TABLE IF NOT EXISTS main_distributed
(
    `user_id` UInt32,
    `app` String,
    `host` String,
    `event` String,
    `ip` String,
    `guid` String,
    `created_at` DateTime('Europe/Moscow')
)
ENGINE = Distributed(cluster_1, default, main)