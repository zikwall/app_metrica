CREATE TABLE IF NOT EXISTS appmetrica.mediavitrina ON CLUSTER `ch-cluster1`
(
    `event_name` String,
    `player_id` String,
    `vitrina_id` String,
    `player` String,

    `sid` String,
    `uid` String,
    `location` String,
    `domain` String,
    `mode` String,
    `drm` UInt8,
    `drm_system_name` String,
    `bitrate` UInt32,
    `event_ts` String,
    `client_time_zone_offset` UInt8,
    `device_type` String,
    `device_vendor` String,
    `device_model` String,
    `user_browser` String,
    `user_browser_ver` String,
    `user_browser_ver_major` String,
    `user_browser_ver_minor` String,
    `user_os_name` String,
    `user_os_ver_major` String,
    `user_os_ver_minor` String,
    `stream_ts` Int32,
    `application_id` String,
    `user_region_iso3166_2` String,
    `content_sec` Int32,
    `pause_sec` Int32,
    `error_title` String,
    `error_adv` String,
    `buffering_sec` Int32,
    `buffering_count` Int32,
    `client_ad_sec` Int32,
    `ad_position` String,
    `init_before_stream_or_ad_request_msec` Int32,
    `stream_or_ad_initial_buffering_msec` Int32,
    `is_subtitles_mode` UInt8,
    `is_fullscreen_mode` UInt8,
    `is_muted` UInt8,
    `product` String,
    `is_web_player` UInt8,
    `is_nobanner` UInt8,

    event_datetime DateTime,
    event_timestamp Int64
)
ENGINE = Distributed('ch-cluster1', 'appmetrica', 'mediavitrina_sharded_v2', rand());

CREATE TABLE IF NOT EXISTS appmetrica.mediavitrina_sharded_v2 ON CLUSTER `ch-cluster1`
(
    `event_name` String,
    `player_id` String,
    `vitrina_id` String,
    `player` String,

    `sid` String,
    `uid` String,
    `location` String,
    `domain` String,
    `mode` String,
    `drm` UInt8,
    `drm_system_name` String,
    `bitrate` UInt32,
    `event_ts` String,
    `client_time_zone_offset` UInt8,
    `device_type` String,
    `device_vendor` String,
    `device_model` String,
    `user_browser` String,
    `user_browser_ver` String,
    `user_browser_ver_major` String,
    `user_browser_ver_minor` String,
    `user_os_name` String,
    `user_os_ver_major` String,
    `user_os_ver_minor` String,
    `stream_ts` Int32,
    `application_id` String,
    `user_region_iso3166_2` String,
    `content_sec` Int32,
    `pause_sec` Int32,
    `error_title` String,
    `error_adv` String,
    `buffering_sec` Int32,
    `buffering_count` Int32,
    `client_ad_sec` Int32,
    `ad_position` String,
    `init_before_stream_or_ad_request_msec` Int32,
    `stream_or_ad_initial_buffering_msec` Int32,
    `is_subtitles_mode` UInt8,
    `is_fullscreen_mode` UInt8,
    `is_muted` UInt8,
    `product` String,
    `is_web_player` UInt8,
    `is_nobanner` UInt8,

    event_datetime DateTime,
    event_timestamp Int64
) ENGINE = ReplicatedMergeTree('/clickhouse/tables/{layer}-{shard}/mediavitrina_sharded_v2', '{replica}')
PARTITION BY toYYYYMM(event_datetime)
ORDER BY (event_datetime, uid)
TTL event_datetime + INTERVAL 181 DAY;
