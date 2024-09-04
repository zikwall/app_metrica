package mediavitrina

import (
	"strconv"
	"time"

	"github.com/zikwall/app_metrica/pkg/domain/event"
)

func QueryParametersToEntity(parameters map[string]string) MediaVitrina {
	now := time.Now()

	return MediaVitrina{
		EventName:                       parameters["event"],
		PlayerID:                        parameters["player_id"],
		VitrinaID:                       parameters["vitrina_id"],
		Player:                          parameters["player"],
		SID:                             parameters["sid"],
		UID:                             parameters["uid"],
		Location:                        parameters["location"],
		Domain:                          parameters["domain"],
		Mode:                            parameters["mode"],
		DRM:                             lookupUInt8(parameters["drm"]),
		DRMSystemName:                   parameters["drm_system_name"],
		Bitrate:                         lookupUInt32(parameters["bitrate"]),
		EventTS:                         parameters["event_ts"],
		ClientTimeZoneOffset:            lookupUInt8(parameters["client_time_zone_offset"]),
		DeviceType:                      parameters["device_type"],
		DeviceVendor:                    parameters["device_vendor"],
		DeviceModel:                     parameters["device_model"],
		UserBrowser:                     parameters["user_browser"],
		UserBrowserVer:                  parameters["user_browser_ver"],
		UserBrowserVerMajor:             parameters["user_browser_ver_major"],
		UserBrowserVerMinor:             parameters["user_browser_ver_minor"],
		UserOSName:                      parameters["user_os_name"],
		UserOSVerMajor:                  parameters["user_os_ver_major"],
		UserOSVerMinor:                  parameters["user_os_ver_minor"],
		StreamTS:                        lookupInt32(parameters["stream_ts"]),
		ApplicationID:                   parameters["application_id"],
		UserRegionISO3166_2:             parameters["user_region_iso3166_2"],
		ContentSec:                      lookupInt32(parameters["content_sec"]),
		PauseSec:                        lookupInt32(parameters["pause_sec"]),
		ErrorTitle:                      parameters["error_title"],
		ErrorAdv:                        parameters["error_adv"],
		BufferingSec:                    lookupInt32(parameters["buffering_sec"]),
		BufferingCount:                  lookupInt32(parameters["buffering_count"]),
		ClientAdSec:                     lookupInt32(parameters["client_ad_sec"]),
		AdPosition:                      parameters["ad_position"],
		InitBeforeStreamOrAdRequestMsec: lookupInt32(parameters["init_before_stream_or_ad_request_msec"]),
		StreamOrAdInitialBufferingMsec:  lookupInt32(parameters["stream_or_ad_initial_buffering_msec"]),
		IsSubtitlesMode:                 lookupUInt8(parameters["is_subtitles_mode"]),
		IsFullscreenMode:                lookupUInt8(parameters["is_fullscreen_mode"]),
		IsMuted:                         lookupUInt8(parameters["is_muted"]),
		Product:                         parameters["product"],
		IsWebPlayer:                     lookupUInt8(parameters["is_web_player"]),
		IsNobanner:                      lookupUInt8(parameters["is_nobanner"]),
		EventDatetime:                   event.EventDatetime{Time: now},
		EventTimestamp:                  now.Unix(),
	}
}

func lookupUInt8(value string) uint8 {
	u8, _ := strconv.ParseInt(value, 10, 8)
	return uint8(u8)
}

func lookupUInt32(value string) uint32 {
	u32, _ := strconv.ParseInt(value, 10, 8)
	return uint32(u32)
}

func lookupInt32(value string) int32 {
	i32, _ := strconv.ParseInt(value, 10, 8)
	return int32(i32)
}
