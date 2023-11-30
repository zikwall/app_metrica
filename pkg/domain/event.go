package domain

import (
	"fmt"
	"strings"
	"time"
)

//easyjson:json
type Event struct {
	ApplicationID      int64         `json:"application_id"`
	IosIfa             string        `json:"ios_ifa"`
	IosIfv             string        `json:"ios_ifv"`
	AndroidID          string        `json:"android_id"`
	GoogleAid          string        `json:"google_aid"`
	ProfileID          string        `json:"profile_id"`
	OsName             string        `json:"os_name"`
	OsVersion          string        `json:"os_version"`
	DeviceManufacturer string        `json:"device_manufacturer"`
	DeviceModel        string        `json:"device_model"`
	DeviceType         string        `json:"device_type"`
	DeviceLocale       string        `json:"device_locale"`
	AppVersionName     string        `json:"app_version_name"`
	AppPackageName     string        `json:"app_package_name"`
	EventName          string        `json:"event_name"`
	EventJSON          string        `json:"event_json"`
	EventDatetime      EventDatetime `json:"event_datetime"`
	EventTimestamp     int64         `json:"event_timestamp"`
	ConnectionType     string        `json:"connection_type"`
	OperatorName       string        `json:"operator_name"`
	Mcc                string        `json:"mcc"`
	Mnc                string        `json:"mnc"`
	AppmetricaDeviceID string        `json:"appmetrica_device_id"`
	InstallationID     string        `json:"installation_id"`
	SessionID          string        `json:"session_id"`
}

//easyjson:json
type EventExtended struct {
	ApplicationID         int64     `json:"application_id"`
	IosIfa                string    `json:"ios_ifa"`
	IosIfv                string    `json:"ios_ifv"`
	AndroidID             string    `json:"android_id"`
	GoogleAid             string    `json:"google_aid"`
	ProfileID             string    `json:"profile_id"`
	OsName                string    `json:"os_name"`
	OsVersion             string    `json:"os_version"`
	DeviceManufacturer    string    `json:"device_manufacturer"`
	DeviceModel           string    `json:"device_model"`
	DeviceType            string    `json:"device_type"`
	DeviceLocale          string    `json:"device_locale"`
	AppVersionName        string    `json:"app_version_name"`
	AppPackageName        string    `json:"app_package_name"`
	EventName             string    `json:"event_name"`
	EventJSON             string    `json:"event_json"`
	EventDatetime         time.Time `json:"event_datetime"`
	EventTimestamp        int64     `json:"event_timestamp"`
	EventReceiveDatetime  time.Time `json:"event_receive_datetime"`
	EventReceiveTimestamp int64     `json:"event_receive_timestamp"`
	ConnectionType        string    `json:"connection_type"`
	OperatorName          string    `json:"operator_name"`
	Mcc                   string    `json:"mcc"`
	Mnc                   string    `json:"mnc"`
	CountryIsoCode        string    `json:"country_iso_code"`
	City                  string    `json:"city"`
	AppmetricaDeviceID    string    `json:"appmetrica_device_id"`
	InstallationID        string    `json:"installation_id"`
	SessionID             string    `json:"session_id"`

	// forward to broker next
	IP string `json:"ip"`
}

func ExtendEvent(e *Event) *EventExtended {
	now := time.Now()

	return &EventExtended{
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
		EventDatetime:      e.EventDatetime.Time,
		EventTimestamp:     e.EventTimestamp,
		ConnectionType:     e.ConnectionType,
		OperatorName:       e.OperatorName,
		Mcc:                e.Mcc,
		Mnc:                e.Mnc,
		AppmetricaDeviceID: e.AppmetricaDeviceID,
		InstallationID:     e.InstallationID,
		SessionID:          e.SessionID,

		EventReceiveDatetime:  now,
		EventReceiveTimestamp: now.Unix(),

		CountryIsoCode: "",
		City:           "",
	}
}

//easyjson:skip
type EventDatetime struct {
	time.Time
}

const dtLayout = "2006-01-02 15:04:05"

func (ct *EventDatetime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(dtLayout, s)
	return
}

func (ct *EventDatetime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(dtLayout))), nil
}
