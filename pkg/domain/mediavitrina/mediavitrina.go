package mediavitrina

import (
	"github.com/mailru/easyjson"

	"github.com/zikwall/app_metrica/pkg/domain/event"
)

//easyjson:json
type MediaVitrina struct {
	EventName                       string              `json:"event_name"`
	PlayerID                        string              `json:"player_id"`
	VitrinaID                       string              `json:"vitrina_id"`
	Player                          string              `json:"player"`
	SID                             string              `json:"sid"`
	UID                             string              `json:"uid"`
	Location                        string              `json:"location"`
	Domain                          string              `json:"domain"`
	Mode                            string              `json:"mode"`
	DRM                             uint8               `json:"drm"`
	DRMSystemName                   string              `json:"drm_system_name"`
	Bitrate                         uint32              `json:"bitrate"`
	EventTS                         string              `json:"event_ts"`
	ClientTimeZoneOffset            uint8               `json:"client_time_zone_offset"`
	DeviceType                      string              `json:"device_type"`
	DeviceVendor                    string              `json:"device_vendor"`
	DeviceModel                     string              `json:"device_model"`
	UserBrowser                     string              `json:"user_browser"`
	UserBrowserVer                  string              `json:"user_browser_ver"`
	UserBrowserVerMajor             string              `json:"user_browser_ver_major"`
	UserBrowserVerMinor             string              `json:"user_browser_ver_minor"`
	UserOSName                      string              `json:"user_os_name"`
	UserOSVerMajor                  string              `json:"user_os_ver_major"`
	UserOSVerMinor                  string              `json:"user_os_ver_minor"`
	StreamTS                        int32               `json:"stream_ts"`
	ApplicationID                   string              `json:"application_id"`
	UserRegionISO3166_2             string              `json:"user_region_iso_3166_2"`
	ContentSec                      int32               `json:"content_sec"`
	PauseSec                        int32               `json:"pause_sec"`
	ErrorTitle                      string              `json:"error_title"`
	ErrorAdv                        string              `json:"error_adv"`
	BufferingSec                    int32               `json:"buffering_sec"`
	BufferingCount                  int32               `json:"buffering_count"`
	ClientAdSec                     int32               `json:"client_ad_sec"`
	AdPosition                      string              `json:"ad_position"`
	InitBeforeStreamOrAdRequestMsec int32               `json:"init_before_stream_or_ad_request_msec"`
	StreamOrAdInitialBufferingMsec  int32               `json:"stream_or_ad_initial_buffering_msec"`
	IsSubtitlesMode                 uint8               `json:"is_subtitles_mode"`
	IsFullscreenMode                uint8               `json:"is_fullscreen_mode"`
	IsMuted                         uint8               `json:"is_muted"`
	Product                         string              `json:"product"`
	IsWebPlayer                     uint8               `json:"is_web_player"`
	IsNobanner                      uint8               `json:"is_nobanner"`
	EventDatetime                   event.EventDatetime `json:"event_datetime"`
	EventTimestamp                  int64               `json:"event_timestamp"`
}

func (m *MediaVitrina) Bytes() ([]byte, error) {
	b, err := easyjson.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}
