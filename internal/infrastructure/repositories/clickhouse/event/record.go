package event

import (
	"time"

	"github.com/zikwall/clickhouse-buffer/v4/src/cx"

	"github.com/zikwall/app_metrica/pkg/domain/event"
)

type ScreenOrientation string

const (
	OrientationLandscape = "landscape"
	OrientationPortrait  = "portrait"
)

func (o ScreenOrientation) UInt8() uint8 {
	switch o {
	case OrientationLandscape:
		return 2
	case OrientationPortrait:
		return 1
	}
	return 3
}

type Record struct {
	ApplicationID      int64     `db:"application_id"`
	IosIfa             string    `db:"ios_ifa"`
	IosIfv             string    `db:"ios_ifv"`
	AndroidID          string    `db:"android_id"`
	GoogleAid          string    `db:"google_aid"`
	ProfileID          string    `db:"profile_id"`
	OsName             string    `db:"os_name"`
	OsVersion          string    `db:"os_version"`
	DeviceManufacturer string    `db:"device_manufacturer"`
	DeviceModel        string    `db:"device_model"`
	DeviceType         string    `db:"device_type"`
	DeviceLocale       string    `db:"device_locale"`
	AppVersionName     string    `db:"app_version_name"`
	AppPackageName     string    `db:"app_package_name"`
	EventName          string    `db:"event_name"`
	EventJSON          string    `db:"event_db"`
	EventDatetime      time.Time `db:"event_datetime"`
	EventTimestamp     int64     `db:"event_timestamp"`
	ConnectionType     string    `db:"connection_type"`
	OperatorName       string    `db:"operator_name"`
	Mcc                string    `db:"mcc"`
	Mnc                string    `db:"mnc"`
	AppmetricaDeviceID string    `db:"appmetrica_device_id"`
	InstallationID     string    `db:"installation_id"`
	SessionID          string    `db:"session_id"`

	EventReceiveDatetime  time.Time `db:"event_receive_datetime"`
	EventReceiveTimestamp int64     `db:"event_receive_timestamp"`

	// additional fields
	IP             string `db:"ip"`
	Region         string `db:"region"`
	AS             uint32 `db:"as"`
	ORG            string `db:"org"`
	CountryIsoCode string `db:"country"`
	City           string `db:"city"`

	Timezone float64 `db:"timezone"`

	PhysicalScreenHeight int    `db:"physical_screen_height"`
	PhysicalScreenWidth  int    `db:"physical_screen_width"`
	ScreenHeight         int    `db:"screen_height"`
	ScreenWeight         int    `db:"screen_weight"`
	ScreenAspectRatio    string `db:"screen_aspect_ratio"`
	ScreenOrientation    uint8  `db:"screen_orientation"`

	Browser        string `db:"browser"`
	BrowserVersion string `db:"browser_version"`
	CookieEnabled  bool   `db:"cookie_enabled"`
	JsEnabled      bool   `db:"js_enabled"`
	Title          string `db:"title"`
	URL            string `db:"url"`
	Referer        string `db:"referer"`

	UtmCampaign string `db:"utm_campaign"`
	UtmContent  string `db:"utm_content"`
	UtmSource   string `db:"utm_source"`
	UtmMedium   string `db:"utm_medium"`
	UtmTerm     string `db:"utm_term"`

	UniqID        string `db:"uniq_id"`
	DeviceID      string `db:"device_id"`
	Platform      string `db:"platform"`
	App           string `db:"app"`
	Version       int    `db:"version"`
	UserAgent     string `db:"user_agent"`
	XLHDAgent     string `db:"xlhd_agent"`
	HardwareOrGUI string `db:"hardware_or_gui"`

	ToQueueDatetime    time.Time `db:"to_queue_datetime"`
	ToQueueTimestamp   int64     `db:"to_queue_timestamp"`
	FromQueueDatetime  time.Time `db:"from_queue_datetime"`
	FromQueueTimestamp int64     `db:"from_queue_timestamp"`

	SdkVersion int `db:"sdk_version"`
}

func (r *Record) Row() cx.Vector {
	return cx.Vector{
		r.ApplicationID,
		r.IosIfa,
		r.IosIfv,
		r.AndroidID,
		r.GoogleAid,
		r.ProfileID,
		r.OsName,
		r.OsVersion,
		r.DeviceManufacturer,
		r.DeviceModel,
		r.DeviceType,
		r.DeviceLocale,
		r.AppVersionName,
		r.AppPackageName,
		r.EventName,
		r.EventJSON,
		r.EventDatetime,
		r.EventTimestamp,
		r.ConnectionType,
		r.OperatorName,
		r.Mcc,
		r.Mnc,
		r.AppmetricaDeviceID,
		r.InstallationID,
		r.SessionID,

		r.EventReceiveDatetime,
		r.EventReceiveTimestamp,

		// additional fields
		r.IP,
		r.Region,
		r.AS,
		r.ORG,
		// from origin
		r.CountryIsoCode,
		r.City,

		r.Timezone,

		r.PhysicalScreenHeight,
		r.PhysicalScreenWidth,
		r.ScreenHeight,
		r.ScreenWeight,
		r.ScreenAspectRatio,
		r.ScreenOrientation,

		r.Browser,
		r.BrowserVersion,
		toUint8(r.CookieEnabled),
		toUint8(r.JsEnabled),
		r.Title,
		r.URL,
		r.Referer,

		r.UtmCampaign,
		r.UtmContent,
		r.UtmSource,
		r.UtmMedium,
		r.UtmTerm,

		r.UniqID,
		r.DeviceID,
		r.Platform,
		r.App,
		r.Version,
		r.UserAgent,
		r.XLHDAgent,
		r.HardwareOrGUI,

		r.ToQueueDatetime,
		r.ToQueueTimestamp,
		r.FromQueueDatetime,
		r.FromQueueTimestamp,

		r.SdkVersion,
	}
}

func recordFromEvent(e *event.EventExtended) *Record {
	return &Record{
		ApplicationID:      e.ApplicationID,
		IosIfa:             e.IosIfa,
		IosIfv:             e.IosIfv,
		AndroidID:          e.AndroidID,
		GoogleAid:          e.GoogleAid,
		ProfileID:          e.ProfileID,
		OsName:             e.OsName,
		OsVersion:          e.OsVersion,
		DeviceManufacturer: e.DeviceManufacturer,
		DeviceModel:        e.DeviceModel,
		DeviceType:         e.DeviceType,
		DeviceLocale:       e.DeviceLocale,
		AppVersionName:     e.AppVersionName,
		AppPackageName:     e.AppPackageName,
		EventName:          e.EventName,
		EventJSON:          e.EventJSON,
		EventDatetime:      e.EventDatetime,
		EventTimestamp:     e.EventTimestamp,
		ConnectionType:     e.ConnectionType,
		OperatorName:       e.OperatorName,
		Mcc:                e.Mcc,
		Mnc:                e.Mnc,
		AppmetricaDeviceID: e.AppmetricaDeviceID,
		InstallationID:     e.InstallationID,
		SessionID:          e.SessionID,

		EventReceiveDatetime:  e.EventReceiveDatetime,
		EventReceiveTimestamp: e.EventTimestamp,

		IP:             e.IP,
		Region:         e.Region,
		AS:             e.AS,
		ORG:            e.ORG,
		CountryIsoCode: e.CountryIsoCode,
		City:           e.City,

		Timezone: e.Timezone,

		PhysicalScreenHeight: e.PhysicalScreenHeight,
		PhysicalScreenWidth:  e.PhysicalScreenWidth,
		ScreenHeight:         e.ScreenHeight,
		ScreenWeight:         e.ScreenWeight,
		ScreenAspectRatio:    e.ScreenAspectRatio,
		ScreenOrientation:    ScreenOrientation(e.ScreenOrientation).UInt8(),

		Browser:        e.Browser,
		BrowserVersion: e.BrowserVersion,
		CookieEnabled:  e.CookieEnabled,
		JsEnabled:      e.JsEnabled,
		Title:          e.Title,
		URL:            e.URL,
		Referer:        e.Referer,

		UtmCampaign: e.UtmCampaign,
		UtmMedium:   e.UtmMedium,
		UtmSource:   e.UtmSource,
		UtmTerm:     e.UtmTerm,
		UtmContent:  e.UtmContent,

		UniqID:        e.UniqID,
		DeviceID:      e.DeviceID,
		Platform:      e.Platform,
		App:           e.App,
		Version:       e.Version,
		UserAgent:     e.UserAgent,
		XLHDAgent:     e.XLHDAgent,
		HardwareOrGUI: e.HardwareOrGUI,

		ToQueueDatetime:  e.ToQueueDatetime,
		ToQueueTimestamp: e.ToQueueTimestamp,

		SdkVersion: e.SdkVersion,

		FromQueueDatetime:  e.FromQueueDatetime,
		FromQueueTimestamp: e.FromQueueDatetime.Unix(),
	}
}

func Columns() []string {
	return []string{
		"application_id",
		"ios_ifa",
		"ios_ifv",
		"android_id",
		"google_aid",
		"profile_id",
		"os_name",
		"os_version",
		"device_manufacturer",
		"device_model",
		"device_type",
		"device_locale",
		"app_version_name",
		"app_package_name",
		"event_name",
		"event_json",
		"event_datetime",
		"event_timestamp",
		"connection_type",
		"operator_name",
		"mcc",
		"mnc",
		"appmetrica_device_id",
		"installation_id",
		"session_id",

		"event_receive_datetime",
		"event_receive_timestamp",

		"ip",
		"region",
		"as",
		"org",
		"country_iso_code",
		"city",

		"timezone",

		"physical_screen_height",
		"physical_screen_width",
		"screen_height",
		"screen_weight",
		"screen_aspect_ratio",
		"screen_orientation",

		"browser",
		"browser_version",
		"cookie_enabled",
		"js_enabled",
		"title",
		"url",
		"referer",

		"utm_campaign",
		"utm_content",
		"utm_source",
		"utm_medium",
		"utm_term",

		"uniq_id",
		"device_id",
		"platform",
		"app",
		"version",
		"user_agent",
		"x_lhd_agent",
		"hardware_or_gui",

		"to_queue_datetime",
		"to_queue_timestamp",
		"from_queue_datetime",
		"from_queue_timestamp",

		"sdk_version",
	}
}

func toUint8(value bool) uint8 {
	if value {
		return 1
	}
	return 0
}
