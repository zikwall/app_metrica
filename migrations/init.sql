CREATE DATABASE IF NOT EXISTS appmetrica ON CLUSTER cluster_1;

CREATE TABLE IF NOT EXISTS appmetrica.events ON CLUSTER cluster_1
(
    `application_id` Int64,
    `ios_ifa` String,
    `ios_ifv` String,
    `android_id` String,
    `google_aid` String,
    `profile_id` String,
    `os_name` String,
    `os_version` String,
    `device_manufacturer` String,
    `device_model` String,
    `device_type` String,
    `device_locale` String,
    `app_version_name` String,
    `app_package_name` String,
    `event_name` String,
    `event_json` String,
    `event_datetime` DateTime,
    `event_timestamp` Int64,
    `connection_type` String,
    `operator_name` String,
    `mcc` String,
    `mnc` String,
    `appmetrica_device_id` String,
    `installation_id` String,
    `session_id` String,

    -- set automatically
    `event_receive_datetime` DateTime,
    `event_receive_timestamp` Int64,

    -- enrich with geodata
    `ip` String,
    `region` String,
    `as` UInt32,
    `org` String,
    `country_iso_code` String,
    `city` String,

    -- additional from app metrica
    `timezone` Int32,

    -- device
    `physical_screen_height` UInt32,
    `physical_screen_width` UInt32,
    `screen_height` UInt32,
    `screen_weight` UInt32,
    `screen_aspect_ratio` String,
    `screen_orientation` UInt8,

    -- web specific fields
    `browser` String,
    `browser_version` String,
    `cookie_enabled` UInt8,
    `js_enabled` UInt8,
    `title` String,
    `url` String,
    `referer` String,

    -- ad campaign
    `utm_campaign` String,
    `utm_content` String,
    `utm_source` String,
    `utm_medium` String,
    `utm_term` String,

    -- internal fields
    `uniq_id` String,
    `device_id` String,
    `platform` String,
    `app` String,
    `version` UInt32,
    `user_agent` String,
    `x_lhd_agent` String,
    `hardware_or_gui` String,

    -- system level
    `to_queue_datetime` DateTime,
    `to_queue_timestamp` Int64,
    `from_queue_datetime` DateTime,
    `from_queue_timestamp` Int64,
    `sdk_version` UInt32
)
ENGINE = Distributed('cluster_1', 'appmetrica', 'events_sharded', rand());

CREATE TABLE IF NOT EXISTS appmetrica.events_sharded ON CLUSTER cluster_1
(
    `application_id` Int64,
    `ios_ifa` String,
    `ios_ifv` String,
    `android_id` String,
    `google_aid` String,
    `profile_id` String,
    `os_name` String,
    `os_version` String,
    `device_manufacturer` String,
    `device_model` String,
    `device_type` String,
    `device_locale` String,
    `app_version_name` String,
    `app_package_name` String,
    `event_name` String,
    `event_json` String,
    `event_datetime` DateTime,
    `event_timestamp` Int64,
    `connection_type` String,
    `operator_name` String,
    `mcc` String,
    `mnc` String,
    `appmetrica_device_id` String,
    `installation_id` String,
    `session_id` String,

    -- set automatically
    `event_receive_datetime` DateTime,
    `event_receive_timestamp` Int64,

    -- enrich with geodata
    `ip` String,
    `region` String,
    `as` UInt32,
    `org` String,
    `country_iso_code` String,
    `city` String,

    -- additional from app metrica
    `timezone` Int32,

    -- device
    `physical_screen_height` UInt32,
    `physical_screen_width` UInt32,
    `screen_height` UInt32,
    `screen_weight` UInt32,
    `screen_aspect_ratio` String,
    `screen_orientation` UInt8,

    -- web specific fields
    `browser` String,
    `browser_version` String,
    `cookie_enabled` UInt8,
    `js_enabled` UInt8,
    `title` String,
    `url` String,
    `referer` String,

    -- ad campaign
    `utm_campaign` String,
    `utm_content` String,
    `utm_source` String,
    `utm_medium` String,
    `utm_term` String,

    -- internal fields
    `uniq_id` String,
    `device_id` String,
    `platform` String,
    `app` String,
    `version` UInt32,
    `user_agent` String,
    `x_lhd_agent` String,
    `hardware_or_gui` String,

    -- system level
    `to_queue_datetime` DateTime,
    `to_queue_timestamp` Int64,
    `from_queue_datetime` DateTime,
    `from_queue_timestamp` Int64,
    `sdk_version` UInt32
) ENGINE = ReplicatedMergeTree('/clickhouse/tables/{layer}-{shard}/events_sharded', '{replica}')
PARTITION BY toYYYYMM(event_receive_datetime)
ORDER BY (event_receive_datetime, appmetrica_device_id, ip)
TTL event_receive_datetime + INTERVAL 16 DAY;

ALTER TABLE appmetrica.events ON CLUSTER cluster_1 ADD COLUMN event_insert_datetime DateTime;
ALTER TABLE appmetrica.events_sharded ON CLUSTER cluster_1 ADD COLUMN event_insert_datetime DateTime default now();
