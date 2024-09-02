package domain

import (
	"fmt"
	"strings"
	"time"
)

//easyjson:json
type Events []*Event

//easyjson:json
type Event struct {
	ApplicationID        int64         `json:"application_id"`
	IosIfa               string        `json:"ios_ifa"`
	IosIfv               string        `json:"ios_ifv"`
	AndroidID            string        `json:"android_id"`
	GoogleAid            string        `json:"google_aid"`
	ProfileID            string        `json:"profile_id"`
	OsName               string        `json:"os_name"`
	OsVersion            string        `json:"os_version"`
	DeviceManufacturer   string        `json:"device_manufacturer"`
	DeviceModel          string        `json:"device_model"`
	DeviceType           string        `json:"device_type"`
	DeviceLocale         string        `json:"device_locale"`
	AppVersionName       string        `json:"app_version_name"`
	AppPackageName       string        `json:"app_package_name"`
	EventName            string        `json:"event_name"`
	EventJSON            string        `json:"event_json"`
	EventDatetime        EventDatetime `json:"event_datetime"`
	EventTimestamp       int64         `json:"event_timestamp"`
	ConnectionType       string        `json:"connection_type"`
	OperatorName         string        `json:"operator_name"`
	Mcc                  string        `json:"mcc"`
	Mnc                  string        `json:"mnc"`
	AppmetricaDeviceID   string        `json:"appmetrica_device_id"`
	InstallationID       string        `json:"installation_id"`
	SessionID            string        `json:"session_id"`
	Timezone             float64       `json:"timezone"`
	PhysicalScreenHeight int           `json:"physical_screen_height"`
	PhysicalScreenWidth  int           `json:"physical_screen_width"`
	ScreenHeight         int           `json:"screen_height"`
	ScreenWeight         int           `json:"screen_weight"`
	ScreenAspectRatio    string        `json:"screen_aspect_ratio"`
	ScreenOrientation    string        `json:"screen_orientation"`
	Browser              string        `json:"browser"`
	BrowserVersion       string        `json:"browser_version"`
	CookieEnabled        bool          `json:"cookie_enabled"`
	JsEnabled            bool          `json:"js_enabled"`
	Title                string        `json:"title"`
	URL                  string        `json:"url"`
	Referer              string        `json:"referer"`
	UtmCampaign          string        `json:"utm_campaign"`
	UtmContent           string        `json:"utm_content"`
	UtmSource            string        `json:"utm_source"`
	UtmMedium            string        `json:"utm_medium"`
	UtmTerm              string        `json:"utm_term"`
	DeviceID             string        `json:"device_id"`
	Platform             string        `json:"platform"`
	App                  string        `json:"app"`
	Version              int           `json:"version"`
	SdkVersion           int           `json:"sdk_version"`
	UserAgent            string        `json:"user_agent"`
	XLHDAgent            string        `json:"xlhd_agent"`
	HardwareOrGUI        string        `json:"hardware_or_gui"`
	UniqID               string        `json:"uniq_id"`
}

//easyjson:json
type EventExtended struct {
	ApplicationID      int64     `json:"application_id"`
	IosIfa             string    `json:"ios_ifa"`
	IosIfv             string    `json:"ios_ifv"`
	AndroidID          string    `json:"android_id"`
	GoogleAid          string    `json:"google_aid"`
	ProfileID          string    `json:"profile_id"`
	OsName             string    `json:"os_name"`
	OsVersion          string    `json:"os_version"`
	DeviceManufacturer string    `json:"device_manufacturer"`
	DeviceModel        string    `json:"device_model"`
	DeviceType         string    `json:"device_type"`
	DeviceLocale       string    `json:"device_locale"`
	AppVersionName     string    `json:"app_version_name"`
	AppPackageName     string    `json:"app_package_name"`
	EventName          string    `json:"event_name"`
	EventJSON          string    `json:"event_json"`
	EventDatetime      time.Time `json:"event_datetime"`
	EventTimestamp     int64     `json:"event_timestamp"`
	ConnectionType     string    `json:"connection_type"`
	OperatorName       string    `json:"operator_name"`
	Mcc                string    `json:"mcc"`
	Mnc                string    `json:"mnc"`
	AppmetricaDeviceID string    `json:"appmetrica_device_id"`
	InstallationID     string    `json:"installation_id"`
	SessionID          string    `json:"session_id"`

	// forward to broker next
	IP string `json:"ip"`

	EventReceiveDatetime  time.Time `json:"event_receive_datetime"`
	EventReceiveTimestamp int64     `json:"event_receive_timestamp"`

	CountryIsoCode string `json:"country_iso_code"`
	City           string `json:"city"`

	Timezone float64 `json:"timezone"`

	PhysicalScreenHeight int    `json:"physical_screen_height"`
	PhysicalScreenWidth  int    `json:"physical_screen_width"`
	ScreenHeight         int    `json:"screen_height"`
	ScreenWeight         int    `json:"screen_weight"`
	ScreenAspectRatio    string `json:"screen_aspect_ratio"`
	ScreenOrientation    string `json:"screen_orientation"`

	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	CookieEnabled  bool   `json:"cookie_enabled"`
	JsEnabled      bool   `json:"js_enabled"`
	Title          string `json:"title"`
	URL            string `json:"url"`
	Referer        string `json:"referer"`

	UtmCampaign string `json:"utm_campaign"`
	UtmContent  string `json:"utm_content"`
	UtmSource   string `json:"utm_source"`
	UtmMedium   string `json:"utm_medium"`
	UtmTerm     string `json:"utm_term"`

	UniqID        string `json:"uniq_id"`
	DeviceID      string `json:"device_id"`
	Platform      string `json:"platform"`
	App           string `json:"app"`
	Version       int    `json:"version"`
	UserAgent     string `json:"user_agent"`
	XLHDAgent     string `json:"xlhd_agent"`
	HardwareOrGUI string `json:"hardware_or_gui"`

	ToQueueDatetime  time.Time `json:"to_queue_datetime"`
	ToQueueTimestamp int64     `json:"to_queue_timestamp"`

	SdkVersion int `json:"sdk_version"`

	FromQueueDatetime time.Time `json:"-"`
	Region            string    `json:"-"`
	AS                uint32    `json:"-"`
	ORG               string    `json:"-"`
}

func ExtendEvent(e *Event, now time.Time, receive time.Time) *EventExtended {
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

		IP: "",

		EventReceiveDatetime:  receive,
		EventReceiveTimestamp: receive.Unix(),

		CountryIsoCode: "",
		City:           "",

		Timezone: e.Timezone,

		PhysicalScreenHeight: e.PhysicalScreenHeight,
		PhysicalScreenWidth:  e.PhysicalScreenWidth,
		ScreenHeight:         e.ScreenHeight,
		ScreenWeight:         e.ScreenWeight,
		ScreenAspectRatio:    e.ScreenAspectRatio,
		ScreenOrientation:    e.ScreenOrientation,

		Browser:        e.Browser,
		BrowserVersion: e.BrowserVersion,
		CookieEnabled:  e.CookieEnabled,
		JsEnabled:      e.JsEnabled,
		Title:          e.Title,
		URL:            e.URL,
		Referer:        e.Referer,

		UtmCampaign: e.UtmCampaign,
		UtmContent:  e.UtmContent,
		UtmSource:   e.UtmSource,
		UtmMedium:   e.UtmMedium,
		UtmTerm:     e.UtmTerm,

		UniqID:        e.UniqID,
		DeviceID:      e.DeviceID,
		Platform:      e.Platform,
		App:           e.App,
		Version:       e.Version,
		UserAgent:     e.UserAgent,
		XLHDAgent:     e.XLHDAgent,
		HardwareOrGUI: e.HardwareOrGUI,

		ToQueueDatetime:  now,
		ToQueueTimestamp: now.Unix(),
		SdkVersion:       e.SdkVersion,
	}
}

//easyjson:skip
type EventDatetime struct {
	time.Time
}

const (
	dtLayout       = "2006-01-02 15:04:05"
	dtLayoutWithTz = "2006-01-02 15:04:05 -07:00"

	lenWithTz = 26
)

func (ct *EventDatetime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	if len(s) == lenWithTz {
		ct.Time, err = time.Parse(dtLayoutWithTz, s)
		return
	}
	ct.Time, err = time.Parse(dtLayout, s)
	return
}

func (ct *EventDatetime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(dtLayoutWithTz))), nil
}

func (ct *EventDatetime) Validate() error {
	minDatetime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	maxDatetime := time.Date(2106, 2, 7, 6, 28, 15, 0, time.UTC)

	if ct.IsZero() {
		return fmt.Errorf("invalid event_datetime: zero value")
	}
	if ct.Before(minDatetime) || ct.After(maxDatetime) {
		return fmt.Errorf("invalid event_datetime: out of range [%s, %s]", minDatetime, maxDatetime)
	}
	return nil
}
