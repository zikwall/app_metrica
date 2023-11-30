CREATE TABLE appmetrica.events ON CLUSTER main_cluster
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

    `event_receive_datetime` DateTime,
    `event_receive_timestamp` Int64,
    `country_iso_code` String,
    `city` String,

    `ip` String,
    `region` String,
    `as` UInt32,
    `org` String
)
ENGINE = Distributed('main_cluster', 'appmetrica', 'events_sharded', rand());

CREATE TABLE appmetrica.events_sharded ON CLUSTER main_cluster
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

   `event_receive_datetime` DateTime,
   `event_receive_timestamp` Int64,
   `country_iso_code` String,
   `city` String,

   `ip` String,
   `region` String,
   `as` UInt32,
   `org` String
) ENGINE = MergeTree()
PARTITION BY toYYYYMM(event_receive_datetime)
ORDER BY (event_receive_datetime, appmetrica_device_id, ip)
TTL event_receive_datetime + INTERVAL 16 DAY;