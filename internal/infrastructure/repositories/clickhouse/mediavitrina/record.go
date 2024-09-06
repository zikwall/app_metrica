package mediavitrina

import (
	"fmt"
	"time"

	"github.com/zikwall/clickhouse-buffer/v4/src/cx"

	"github.com/zikwall/app_metrica/pkg/domain/mediavitrina"
)

type Record struct {
	EventName                       string    `db:"event_name"`
	PlayerID                        string    `db:"player_id"`
	VitrinaID                       string    `db:"vitrina_id"`
	Player                          string    `db:"player"`
	SID                             string    `db:"sid"`
	UID                             string    `db:"uid"`
	Location                        string    `db:"location"`
	Domain                          string    `db:"domain"`
	Mode                            string    `db:"mode"`
	DRM                             uint8     `db:"drm"`
	DRMSystemName                   string    `db:"drm_system_name"`
	Bitrate                         uint32    `db:"bitrate"`
	EventTS                         string    `db:"event_ts"`
	ClientTimeZoneOffset            uint8     `db:"client_time_zone_offset"`
	DeviceType                      string    `db:"device_type"`
	DeviceVendor                    string    `db:"device_vendor"`
	DeviceModel                     string    `db:"device_model"`
	UserBrowser                     string    `db:"user_browser"`
	UserBrowserVer                  string    `db:"user_browser_ver"`
	UserBrowserVerMajor             string    `db:"user_browser_ver_major"`
	UserBrowserVerMinor             string    `db:"user_browser_ver_minor"`
	UserOSName                      string    `db:"user_os_name"`
	UserOSVerMajor                  string    `db:"user_os_ver_major"`
	UserOSVerMinor                  string    `db:"user_os_ver_minor"`
	StreamTS                        int32     `db:"stream_ts"`
	ApplicationID                   string    `db:"application_id"`
	UserRegionISO3166_2             string    `db:"user_region_iso3166_2"`
	ContentSec                      int32     `db:"content_sec"`
	PauseSec                        int32     `db:"pause_sec"`
	ErrorTitle                      string    `db:"error_title"`
	ErrorAdv                        string    `db:"error_adv"`
	BufferingSec                    int32     `db:"buffering_sec"`
	BufferingCount                  int32     `db:"buffering_count"`
	ClientAdSec                     int32     `db:"client_ad_sec"`
	AdPosition                      string    `db:"ad_position"`
	InitBeforeStreamOrAdRequestMsec int32     `db:"init_before_stream_or_ad_request_msec"`
	StreamOrAdInitialBufferingMsec  int32     `db:"stream_or_ad_initial_buffering_msec"`
	IsSubtitlesMode                 uint8     `db:"is_subtitles_mode"`
	IsFullscreenMode                uint8     `db:"is_fullscreen_mode"`
	IsMuted                         uint8     `db:"is_muted"`
	Product                         string    `db:"product"`
	IsWebPlayer                     uint8     `db:"is_web_player"`
	IsNobanner                      uint8     `db:"is_nobanner"`
	EventDatetime                   time.Time `db:"event_datetime"`
	EventTimestamp                  int64     `db:"event_timestamp"`
}

func (r *Record) Row() cx.Vector {
	return cx.Vector{
		r.EventName,
		r.PlayerID,
		r.VitrinaID,
		r.Player,
		r.SID,
		r.UID,
		r.Location,
		r.Domain,
		r.Mode,
		r.DRM,
		r.DRMSystemName,
		r.Bitrate,
		r.EventTS,
		r.ClientTimeZoneOffset,
		r.DeviceType,
		r.DeviceVendor,
		r.DeviceModel,
		r.UserBrowser,
		r.UserBrowserVer,
		r.UserBrowserVerMajor,
		r.UserBrowserVerMinor,
		r.UserOSName,
		r.UserOSVerMajor,
		r.UserOSVerMinor,
		r.StreamTS,
		r.ApplicationID,
		r.UserRegionISO3166_2,
		r.ContentSec,
		r.PauseSec,
		r.ErrorTitle,
		r.ErrorAdv,
		r.BufferingSec,
		r.BufferingCount,
		r.ClientAdSec,
		r.AdPosition,
		r.InitBeforeStreamOrAdRequestMsec,
		r.StreamOrAdInitialBufferingMsec,
		r.IsSubtitlesMode,
		r.IsFullscreenMode,
		r.IsMuted,
		r.Product,
		r.IsWebPlayer,
		r.IsNobanner,
		r.EventDatetime,
		r.EventTimestamp,
	}
}

func recordFromEvent(e *mediavitrina.MediaVitrina) *Record {
	eventDatetime := time.Now()

	fmt.Println(eventDatetime)

	return &Record{
		EventName:                       e.EventName,
		PlayerID:                        e.PlayerID,
		VitrinaID:                       e.VitrinaID,
		Player:                          e.Player,
		SID:                             e.SID,
		UID:                             e.UID,
		Location:                        e.Location,
		Domain:                          e.Domain,
		Mode:                            e.Mode,
		DRM:                             e.DRM,
		DRMSystemName:                   e.DRMSystemName,
		Bitrate:                         e.Bitrate,
		EventTS:                         e.EventTS,
		ClientTimeZoneOffset:            e.ClientTimeZoneOffset,
		DeviceType:                      e.DeviceType,
		DeviceVendor:                    e.DeviceVendor,
		DeviceModel:                     e.DeviceModel,
		UserBrowser:                     e.UserBrowser,
		UserBrowserVer:                  e.UserBrowserVer,
		UserBrowserVerMajor:             e.UserBrowserVerMajor,
		UserBrowserVerMinor:             e.UserBrowserVerMinor,
		UserOSName:                      e.UserOSName,
		UserOSVerMajor:                  e.UserOSVerMajor,
		UserOSVerMinor:                  e.UserOSVerMinor,
		StreamTS:                        e.StreamTS,
		ApplicationID:                   e.ApplicationID,
		UserRegionISO3166_2:             e.UserRegionISO3166_2,
		ContentSec:                      e.ContentSec,
		PauseSec:                        e.PauseSec,
		ErrorTitle:                      e.ErrorTitle,
		ErrorAdv:                        e.ErrorAdv,
		BufferingSec:                    e.BufferingSec,
		BufferingCount:                  e.BufferingCount,
		ClientAdSec:                     e.ClientAdSec,
		AdPosition:                      e.AdPosition,
		InitBeforeStreamOrAdRequestMsec: e.InitBeforeStreamOrAdRequestMsec,
		StreamOrAdInitialBufferingMsec:  e.StreamOrAdInitialBufferingMsec,
		IsSubtitlesMode:                 e.IsSubtitlesMode,
		IsFullscreenMode:                e.IsFullscreenMode,
		IsMuted:                         e.IsMuted,
		Product:                         e.Product,
		IsWebPlayer:                     e.IsWebPlayer,
		IsNobanner:                      e.IsNobanner,
		EventDatetime:                   eventDatetime,
		EventTimestamp:                  eventDatetime.Unix(),
	}
}

func Columns() []string {
	return []string{
		"event_name",
		"player_id",
		"vitrina_id",
		"player",
		"sid",
		"uid",
		"location",
		"domain",
		"mode",
		"drm",
		"drm_system_name",
		"bitrate",
		"event_ts",
		"client_time_zone_offset",
		"device_type",
		"device_vendor",
		"device_model",
		"user_browser",
		"user_browser_ver",
		"user_browser_ver_major",
		"user_browser_ver_minor",
		"user_os_name",
		"user_os_ver_major",
		"user_os_ver_minor",
		"stream_ts",
		"application_id",
		"user_region_iso3166_2",
		"content_sec",
		"pause_sec",
		"error_title",
		"error_adv",
		"buffering_sec",
		"buffering_count",
		"client_ad_sec",
		"ad_position",
		"init_before_stream_or_ad_request_msec",
		"stream_or_ad_initial_buffering_msec",
		"is_subtitles_mode",
		"is_fullscreen_mode",
		"is_muted",
		"product",
		"is_web_player",
		"is_nobanner",
		"event_datetime",
		"event_timestamp",
	}
}
