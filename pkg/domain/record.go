package domain

import (
	"time"

	"github.com/zikwall/clickhouse-buffer/v4/src/cx"
)

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

	// from origin
	CountryIsoCode string `db:"country"`
	City           string `db:"city"`

	EventReceiveDatetime  time.Time `db:"event_receive_datetime"`
	EventReceiveTimestamp int64     `db:"event_receive_timestamp"`

	// additional fields
	IP     string `db:"ip"`
	Region string `db:"region"`
	AS     uint32 `db:"as"`
	ORG    string `db:"org"`
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

		// from origin
		r.CountryIsoCode,
		r.City,

		// additional fields
		r.IP,
		r.Region,
		r.AS,
		r.ORG,
	}
}

func RecordFromEvent(e *EventExtended) *Record {
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

		CountryIsoCode: "",
		City:           "",

		IP:     e.IP,
		Region: "",
		AS:     0,
		ORG:    "",
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

		"country_iso_code",
		"city",

		"ip",
		"region",
		"as",
		"org",
	}
}
