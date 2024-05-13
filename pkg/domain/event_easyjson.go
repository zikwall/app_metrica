// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package domain

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonF642ad3eDecodeGithubComZikwallAppMetricaPkgDomain(in *jlexer.Lexer, out *Events) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Events, 0, 8)
			} else {
				*out = Events{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *Event
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(Event)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF642ad3eEncodeGithubComZikwallAppMetricaPkgDomain(out *jwriter.Writer, in Events) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				(*v3).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Events) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF642ad3eEncodeGithubComZikwallAppMetricaPkgDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Events) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF642ad3eEncodeGithubComZikwallAppMetricaPkgDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Events) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF642ad3eDecodeGithubComZikwallAppMetricaPkgDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Events) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF642ad3eDecodeGithubComZikwallAppMetricaPkgDomain(l, v)
}
func easyjsonF642ad3eDecodeGithubComZikwallAppMetricaPkgDomain1(in *jlexer.Lexer, out *EventExtended) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "application_id":
			out.ApplicationID = int64(in.Int64())
		case "ios_ifa":
			out.IosIfa = string(in.String())
		case "ios_ifv":
			out.IosIfv = string(in.String())
		case "android_id":
			out.AndroidID = string(in.String())
		case "google_aid":
			out.GoogleAid = string(in.String())
		case "profile_id":
			out.ProfileID = string(in.String())
		case "os_name":
			out.OsName = string(in.String())
		case "os_version":
			out.OsVersion = string(in.String())
		case "device_manufacturer":
			out.DeviceManufacturer = string(in.String())
		case "device_model":
			out.DeviceModel = string(in.String())
		case "device_type":
			out.DeviceType = string(in.String())
		case "device_locale":
			out.DeviceLocale = string(in.String())
		case "app_version_name":
			out.AppVersionName = string(in.String())
		case "app_package_name":
			out.AppPackageName = string(in.String())
		case "event_name":
			out.EventName = string(in.String())
		case "event_json":
			out.EventJSON = string(in.String())
		case "event_datetime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EventDatetime).UnmarshalJSON(data))
			}
		case "event_timestamp":
			out.EventTimestamp = int64(in.Int64())
		case "connection_type":
			out.ConnectionType = string(in.String())
		case "operator_name":
			out.OperatorName = string(in.String())
		case "mcc":
			out.Mcc = string(in.String())
		case "mnc":
			out.Mnc = string(in.String())
		case "appmetrica_device_id":
			out.AppmetricaDeviceID = string(in.String())
		case "installation_id":
			out.InstallationID = string(in.String())
		case "session_id":
			out.SessionID = string(in.String())
		case "ip":
			out.IP = string(in.String())
		case "event_receive_datetime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EventReceiveDatetime).UnmarshalJSON(data))
			}
		case "event_receive_timestamp":
			out.EventReceiveTimestamp = int64(in.Int64())
		case "country_iso_code":
			out.CountryIsoCode = string(in.String())
		case "city":
			out.City = string(in.String())
		case "timezone":
			out.Timezone = float64(in.Float64())
		case "physical_screen_height":
			out.PhysicalScreenHeight = int(in.Int())
		case "physical_screen_width":
			out.PhysicalScreenWidth = int(in.Int())
		case "screen_height":
			out.ScreenHeight = int(in.Int())
		case "screen_weight":
			out.ScreenWeight = int(in.Int())
		case "screen_aspect_ratio":
			out.ScreenAspectRatio = string(in.String())
		case "screen_orientation":
			out.ScreenOrientation = string(in.String())
		case "browser":
			out.Browser = string(in.String())
		case "browser_version":
			out.BrowserVersion = string(in.String())
		case "cookie_enabled":
			out.CookieEnabled = bool(in.Bool())
		case "js_enabled":
			out.JsEnabled = bool(in.Bool())
		case "title":
			out.Title = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "referer":
			out.Referer = string(in.String())
		case "utm_campaign":
			out.UtmCampaign = string(in.String())
		case "utm_content":
			out.UtmContent = string(in.String())
		case "utm_source":
			out.UtmSource = string(in.String())
		case "utm_medium":
			out.UtmMedium = string(in.String())
		case "utm_term":
			out.UtmTerm = string(in.String())
		case "uniq_id":
			out.UniqID = string(in.String())
		case "device_id":
			out.DeviceID = string(in.String())
		case "platform":
			out.Platform = string(in.String())
		case "app":
			out.App = string(in.String())
		case "version":
			out.Version = int(in.Int())
		case "user_agent":
			out.UserAgent = string(in.String())
		case "xlhd_agent":
			out.XLHDAgent = string(in.String())
		case "hardware_or_gui":
			out.HardwareOrGUI = string(in.String())
		case "to_queue_datetime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ToQueueDatetime).UnmarshalJSON(data))
			}
		case "to_queue_timestamp":
			out.ToQueueTimestamp = int64(in.Int64())
		case "sdk_version":
			out.SdkVersion = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF642ad3eEncodeGithubComZikwallAppMetricaPkgDomain1(out *jwriter.Writer, in EventExtended) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"application_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ApplicationID))
	}
	{
		const prefix string = ",\"ios_ifa\":"
		out.RawString(prefix)
		out.String(string(in.IosIfa))
	}
	{
		const prefix string = ",\"ios_ifv\":"
		out.RawString(prefix)
		out.String(string(in.IosIfv))
	}
	{
		const prefix string = ",\"android_id\":"
		out.RawString(prefix)
		out.String(string(in.AndroidID))
	}
	{
		const prefix string = ",\"google_aid\":"
		out.RawString(prefix)
		out.String(string(in.GoogleAid))
	}
	{
		const prefix string = ",\"profile_id\":"
		out.RawString(prefix)
		out.String(string(in.ProfileID))
	}
	{
		const prefix string = ",\"os_name\":"
		out.RawString(prefix)
		out.String(string(in.OsName))
	}
	{
		const prefix string = ",\"os_version\":"
		out.RawString(prefix)
		out.String(string(in.OsVersion))
	}
	{
		const prefix string = ",\"device_manufacturer\":"
		out.RawString(prefix)
		out.String(string(in.DeviceManufacturer))
	}
	{
		const prefix string = ",\"device_model\":"
		out.RawString(prefix)
		out.String(string(in.DeviceModel))
	}
	{
		const prefix string = ",\"device_type\":"
		out.RawString(prefix)
		out.String(string(in.DeviceType))
	}
	{
		const prefix string = ",\"device_locale\":"
		out.RawString(prefix)
		out.String(string(in.DeviceLocale))
	}
	{
		const prefix string = ",\"app_version_name\":"
		out.RawString(prefix)
		out.String(string(in.AppVersionName))
	}
	{
		const prefix string = ",\"app_package_name\":"
		out.RawString(prefix)
		out.String(string(in.AppPackageName))
	}
	{
		const prefix string = ",\"event_name\":"
		out.RawString(prefix)
		out.String(string(in.EventName))
	}
	{
		const prefix string = ",\"event_json\":"
		out.RawString(prefix)
		out.String(string(in.EventJSON))
	}
	{
		const prefix string = ",\"event_datetime\":"
		out.RawString(prefix)
		out.Raw((in.EventDatetime).MarshalJSON())
	}
	{
		const prefix string = ",\"event_timestamp\":"
		out.RawString(prefix)
		out.Int64(int64(in.EventTimestamp))
	}
	{
		const prefix string = ",\"connection_type\":"
		out.RawString(prefix)
		out.String(string(in.ConnectionType))
	}
	{
		const prefix string = ",\"operator_name\":"
		out.RawString(prefix)
		out.String(string(in.OperatorName))
	}
	{
		const prefix string = ",\"mcc\":"
		out.RawString(prefix)
		out.String(string(in.Mcc))
	}
	{
		const prefix string = ",\"mnc\":"
		out.RawString(prefix)
		out.String(string(in.Mnc))
	}
	{
		const prefix string = ",\"appmetrica_device_id\":"
		out.RawString(prefix)
		out.String(string(in.AppmetricaDeviceID))
	}
	{
		const prefix string = ",\"installation_id\":"
		out.RawString(prefix)
		out.String(string(in.InstallationID))
	}
	{
		const prefix string = ",\"session_id\":"
		out.RawString(prefix)
		out.String(string(in.SessionID))
	}
	{
		const prefix string = ",\"ip\":"
		out.RawString(prefix)
		out.String(string(in.IP))
	}
	{
		const prefix string = ",\"event_receive_datetime\":"
		out.RawString(prefix)
		out.Raw((in.EventReceiveDatetime).MarshalJSON())
	}
	{
		const prefix string = ",\"event_receive_timestamp\":"
		out.RawString(prefix)
		out.Int64(int64(in.EventReceiveTimestamp))
	}
	{
		const prefix string = ",\"country_iso_code\":"
		out.RawString(prefix)
		out.String(string(in.CountryIsoCode))
	}
	{
		const prefix string = ",\"city\":"
		out.RawString(prefix)
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"timezone\":"
		out.RawString(prefix)
		out.Float64(float64(in.Timezone))
	}
	{
		const prefix string = ",\"physical_screen_height\":"
		out.RawString(prefix)
		out.Int(int(in.PhysicalScreenHeight))
	}
	{
		const prefix string = ",\"physical_screen_width\":"
		out.RawString(prefix)
		out.Int(int(in.PhysicalScreenWidth))
	}
	{
		const prefix string = ",\"screen_height\":"
		out.RawString(prefix)
		out.Int(int(in.ScreenHeight))
	}
	{
		const prefix string = ",\"screen_weight\":"
		out.RawString(prefix)
		out.Int(int(in.ScreenWeight))
	}
	{
		const prefix string = ",\"screen_aspect_ratio\":"
		out.RawString(prefix)
		out.String(string(in.ScreenAspectRatio))
	}
	{
		const prefix string = ",\"screen_orientation\":"
		out.RawString(prefix)
		out.String(string(in.ScreenOrientation))
	}
	{
		const prefix string = ",\"browser\":"
		out.RawString(prefix)
		out.String(string(in.Browser))
	}
	{
		const prefix string = ",\"browser_version\":"
		out.RawString(prefix)
		out.String(string(in.BrowserVersion))
	}
	{
		const prefix string = ",\"cookie_enabled\":"
		out.RawString(prefix)
		out.Bool(bool(in.CookieEnabled))
	}
	{
		const prefix string = ",\"js_enabled\":"
		out.RawString(prefix)
		out.Bool(bool(in.JsEnabled))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	{
		const prefix string = ",\"referer\":"
		out.RawString(prefix)
		out.String(string(in.Referer))
	}
	{
		const prefix string = ",\"utm_campaign\":"
		out.RawString(prefix)
		out.String(string(in.UtmCampaign))
	}
	{
		const prefix string = ",\"utm_content\":"
		out.RawString(prefix)
		out.String(string(in.UtmContent))
	}
	{
		const prefix string = ",\"utm_source\":"
		out.RawString(prefix)
		out.String(string(in.UtmSource))
	}
	{
		const prefix string = ",\"utm_medium\":"
		out.RawString(prefix)
		out.String(string(in.UtmMedium))
	}
	{
		const prefix string = ",\"utm_term\":"
		out.RawString(prefix)
		out.String(string(in.UtmTerm))
	}
	{
		const prefix string = ",\"uniq_id\":"
		out.RawString(prefix)
		out.String(string(in.UniqID))
	}
	{
		const prefix string = ",\"device_id\":"
		out.RawString(prefix)
		out.String(string(in.DeviceID))
	}
	{
		const prefix string = ",\"platform\":"
		out.RawString(prefix)
		out.String(string(in.Platform))
	}
	{
		const prefix string = ",\"app\":"
		out.RawString(prefix)
		out.String(string(in.App))
	}
	{
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.Int(int(in.Version))
	}
	{
		const prefix string = ",\"user_agent\":"
		out.RawString(prefix)
		out.String(string(in.UserAgent))
	}
	{
		const prefix string = ",\"xlhd_agent\":"
		out.RawString(prefix)
		out.String(string(in.XLHDAgent))
	}
	{
		const prefix string = ",\"hardware_or_gui\":"
		out.RawString(prefix)
		out.String(string(in.HardwareOrGUI))
	}
	{
		const prefix string = ",\"to_queue_datetime\":"
		out.RawString(prefix)
		out.Raw((in.ToQueueDatetime).MarshalJSON())
	}
	{
		const prefix string = ",\"to_queue_timestamp\":"
		out.RawString(prefix)
		out.Int64(int64(in.ToQueueTimestamp))
	}
	{
		const prefix string = ",\"sdk_version\":"
		out.RawString(prefix)
		out.Int(int(in.SdkVersion))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventExtended) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF642ad3eEncodeGithubComZikwallAppMetricaPkgDomain1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventExtended) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF642ad3eEncodeGithubComZikwallAppMetricaPkgDomain1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventExtended) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF642ad3eDecodeGithubComZikwallAppMetricaPkgDomain1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventExtended) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF642ad3eDecodeGithubComZikwallAppMetricaPkgDomain1(l, v)
}
func easyjsonF642ad3eDecodeGithubComZikwallAppMetricaPkgDomain2(in *jlexer.Lexer, out *Event) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "application_id":
			out.ApplicationID = int64(in.Int64())
		case "ios_ifa":
			out.IosIfa = string(in.String())
		case "ios_ifv":
			out.IosIfv = string(in.String())
		case "android_id":
			out.AndroidID = string(in.String())
		case "google_aid":
			out.GoogleAid = string(in.String())
		case "profile_id":
			out.ProfileID = string(in.String())
		case "os_name":
			out.OsName = string(in.String())
		case "os_version":
			out.OsVersion = string(in.String())
		case "device_manufacturer":
			out.DeviceManufacturer = string(in.String())
		case "device_model":
			out.DeviceModel = string(in.String())
		case "device_type":
			out.DeviceType = string(in.String())
		case "device_locale":
			out.DeviceLocale = string(in.String())
		case "app_version_name":
			out.AppVersionName = string(in.String())
		case "app_package_name":
			out.AppPackageName = string(in.String())
		case "event_name":
			out.EventName = string(in.String())
		case "event_json":
			out.EventJSON = string(in.String())
		case "event_datetime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EventDatetime).UnmarshalJSON(data))
			}
		case "event_timestamp":
			out.EventTimestamp = int64(in.Int64())
		case "connection_type":
			out.ConnectionType = string(in.String())
		case "operator_name":
			out.OperatorName = string(in.String())
		case "mcc":
			out.Mcc = string(in.String())
		case "mnc":
			out.Mnc = string(in.String())
		case "appmetrica_device_id":
			out.AppmetricaDeviceID = string(in.String())
		case "installation_id":
			out.InstallationID = string(in.String())
		case "session_id":
			out.SessionID = string(in.String())
		case "timezone":
			out.Timezone = float64(in.Float64())
		case "physical_screen_height":
			out.PhysicalScreenHeight = int(in.Int())
		case "physical_screen_width":
			out.PhysicalScreenWidth = int(in.Int())
		case "screen_height":
			out.ScreenHeight = int(in.Int())
		case "screen_weight":
			out.ScreenWeight = int(in.Int())
		case "screen_aspect_ratio":
			out.ScreenAspectRatio = string(in.String())
		case "screen_orientation":
			out.ScreenOrientation = string(in.String())
		case "browser":
			out.Browser = string(in.String())
		case "browser_version":
			out.BrowserVersion = string(in.String())
		case "cookie_enabled":
			out.CookieEnabled = bool(in.Bool())
		case "js_enabled":
			out.JsEnabled = bool(in.Bool())
		case "title":
			out.Title = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "referer":
			out.Referer = string(in.String())
		case "utm_campaign":
			out.UtmCampaign = string(in.String())
		case "utm_content":
			out.UtmContent = string(in.String())
		case "utm_source":
			out.UtmSource = string(in.String())
		case "utm_medium":
			out.UtmMedium = string(in.String())
		case "utm_term":
			out.UtmTerm = string(in.String())
		case "device_id":
			out.DeviceID = string(in.String())
		case "platform":
			out.Platform = string(in.String())
		case "app":
			out.App = string(in.String())
		case "version":
			out.Version = int(in.Int())
		case "sdk_version":
			out.SdkVersion = int(in.Int())
		case "user_agent":
			out.UserAgent = string(in.String())
		case "xlhd_agent":
			out.XLHDAgent = string(in.String())
		case "hardware_or_gui":
			out.HardwareOrGUI = string(in.String())
		case "uniq_id":
			out.UniqID = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF642ad3eEncodeGithubComZikwallAppMetricaPkgDomain2(out *jwriter.Writer, in Event) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"application_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ApplicationID))
	}
	{
		const prefix string = ",\"ios_ifa\":"
		out.RawString(prefix)
		out.String(string(in.IosIfa))
	}
	{
		const prefix string = ",\"ios_ifv\":"
		out.RawString(prefix)
		out.String(string(in.IosIfv))
	}
	{
		const prefix string = ",\"android_id\":"
		out.RawString(prefix)
		out.String(string(in.AndroidID))
	}
	{
		const prefix string = ",\"google_aid\":"
		out.RawString(prefix)
		out.String(string(in.GoogleAid))
	}
	{
		const prefix string = ",\"profile_id\":"
		out.RawString(prefix)
		out.String(string(in.ProfileID))
	}
	{
		const prefix string = ",\"os_name\":"
		out.RawString(prefix)
		out.String(string(in.OsName))
	}
	{
		const prefix string = ",\"os_version\":"
		out.RawString(prefix)
		out.String(string(in.OsVersion))
	}
	{
		const prefix string = ",\"device_manufacturer\":"
		out.RawString(prefix)
		out.String(string(in.DeviceManufacturer))
	}
	{
		const prefix string = ",\"device_model\":"
		out.RawString(prefix)
		out.String(string(in.DeviceModel))
	}
	{
		const prefix string = ",\"device_type\":"
		out.RawString(prefix)
		out.String(string(in.DeviceType))
	}
	{
		const prefix string = ",\"device_locale\":"
		out.RawString(prefix)
		out.String(string(in.DeviceLocale))
	}
	{
		const prefix string = ",\"app_version_name\":"
		out.RawString(prefix)
		out.String(string(in.AppVersionName))
	}
	{
		const prefix string = ",\"app_package_name\":"
		out.RawString(prefix)
		out.String(string(in.AppPackageName))
	}
	{
		const prefix string = ",\"event_name\":"
		out.RawString(prefix)
		out.String(string(in.EventName))
	}
	{
		const prefix string = ",\"event_json\":"
		out.RawString(prefix)
		out.String(string(in.EventJSON))
	}
	{
		const prefix string = ",\"event_datetime\":"
		out.RawString(prefix)
		out.Raw((in.EventDatetime).MarshalJSON())
	}
	{
		const prefix string = ",\"event_timestamp\":"
		out.RawString(prefix)
		out.Int64(int64(in.EventTimestamp))
	}
	{
		const prefix string = ",\"connection_type\":"
		out.RawString(prefix)
		out.String(string(in.ConnectionType))
	}
	{
		const prefix string = ",\"operator_name\":"
		out.RawString(prefix)
		out.String(string(in.OperatorName))
	}
	{
		const prefix string = ",\"mcc\":"
		out.RawString(prefix)
		out.String(string(in.Mcc))
	}
	{
		const prefix string = ",\"mnc\":"
		out.RawString(prefix)
		out.String(string(in.Mnc))
	}
	{
		const prefix string = ",\"appmetrica_device_id\":"
		out.RawString(prefix)
		out.String(string(in.AppmetricaDeviceID))
	}
	{
		const prefix string = ",\"installation_id\":"
		out.RawString(prefix)
		out.String(string(in.InstallationID))
	}
	{
		const prefix string = ",\"session_id\":"
		out.RawString(prefix)
		out.String(string(in.SessionID))
	}
	{
		const prefix string = ",\"timezone\":"
		out.RawString(prefix)
		out.Float64(float64(in.Timezone))
	}
	{
		const prefix string = ",\"physical_screen_height\":"
		out.RawString(prefix)
		out.Int(int(in.PhysicalScreenHeight))
	}
	{
		const prefix string = ",\"physical_screen_width\":"
		out.RawString(prefix)
		out.Int(int(in.PhysicalScreenWidth))
	}
	{
		const prefix string = ",\"screen_height\":"
		out.RawString(prefix)
		out.Int(int(in.ScreenHeight))
	}
	{
		const prefix string = ",\"screen_weight\":"
		out.RawString(prefix)
		out.Int(int(in.ScreenWeight))
	}
	{
		const prefix string = ",\"screen_aspect_ratio\":"
		out.RawString(prefix)
		out.String(string(in.ScreenAspectRatio))
	}
	{
		const prefix string = ",\"screen_orientation\":"
		out.RawString(prefix)
		out.String(string(in.ScreenOrientation))
	}
	{
		const prefix string = ",\"browser\":"
		out.RawString(prefix)
		out.String(string(in.Browser))
	}
	{
		const prefix string = ",\"browser_version\":"
		out.RawString(prefix)
		out.String(string(in.BrowserVersion))
	}
	{
		const prefix string = ",\"cookie_enabled\":"
		out.RawString(prefix)
		out.Bool(bool(in.CookieEnabled))
	}
	{
		const prefix string = ",\"js_enabled\":"
		out.RawString(prefix)
		out.Bool(bool(in.JsEnabled))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	{
		const prefix string = ",\"referer\":"
		out.RawString(prefix)
		out.String(string(in.Referer))
	}
	{
		const prefix string = ",\"utm_campaign\":"
		out.RawString(prefix)
		out.String(string(in.UtmCampaign))
	}
	{
		const prefix string = ",\"utm_content\":"
		out.RawString(prefix)
		out.String(string(in.UtmContent))
	}
	{
		const prefix string = ",\"utm_source\":"
		out.RawString(prefix)
		out.String(string(in.UtmSource))
	}
	{
		const prefix string = ",\"utm_medium\":"
		out.RawString(prefix)
		out.String(string(in.UtmMedium))
	}
	{
		const prefix string = ",\"utm_term\":"
		out.RawString(prefix)
		out.String(string(in.UtmTerm))
	}
	{
		const prefix string = ",\"device_id\":"
		out.RawString(prefix)
		out.String(string(in.DeviceID))
	}
	{
		const prefix string = ",\"platform\":"
		out.RawString(prefix)
		out.String(string(in.Platform))
	}
	{
		const prefix string = ",\"app\":"
		out.RawString(prefix)
		out.String(string(in.App))
	}
	{
		const prefix string = ",\"version\":"
		out.RawString(prefix)
		out.Int(int(in.Version))
	}
	{
		const prefix string = ",\"sdk_version\":"
		out.RawString(prefix)
		out.Int(int(in.SdkVersion))
	}
	{
		const prefix string = ",\"user_agent\":"
		out.RawString(prefix)
		out.String(string(in.UserAgent))
	}
	{
		const prefix string = ",\"xlhd_agent\":"
		out.RawString(prefix)
		out.String(string(in.XLHDAgent))
	}
	{
		const prefix string = ",\"hardware_or_gui\":"
		out.RawString(prefix)
		out.String(string(in.HardwareOrGUI))
	}
	{
		const prefix string = ",\"uniq_id\":"
		out.RawString(prefix)
		out.String(string(in.UniqID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Event) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF642ad3eEncodeGithubComZikwallAppMetricaPkgDomain2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Event) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF642ad3eEncodeGithubComZikwallAppMetricaPkgDomain2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Event) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF642ad3eDecodeGithubComZikwallAppMetricaPkgDomain2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Event) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF642ad3eDecodeGithubComZikwallAppMetricaPkgDomain2(l, v)
}
