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

func easyjsonF642ad3eDecodeGithubComZikwallAppMetricaPkgDomain(in *jlexer.Lexer, out *EventExtended) {
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
		case "event_receive_datetime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.EventReceiveDatetime).UnmarshalJSON(data))
			}
		case "event_receive_timestamp":
			out.EventReceiveTimestamp = int64(in.Int64())
		case "connection_type":
			out.ConnectionType = string(in.String())
		case "operator_name":
			out.OperatorName = string(in.String())
		case "mcc":
			out.Mcc = string(in.String())
		case "mnc":
			out.Mnc = string(in.String())
		case "country_iso_code":
			out.CountryIsoCode = string(in.String())
		case "city":
			out.City = string(in.String())
		case "appmetrica_device_id":
			out.AppmetricaDeviceID = string(in.String())
		case "installation_id":
			out.InstallationID = string(in.String())
		case "session_id":
			out.SessionID = string(in.String())
		case "ip":
			out.IP = string(in.String())
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
func easyjsonF642ad3eEncodeGithubComZikwallAppMetricaPkgDomain(out *jwriter.Writer, in EventExtended) {
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EventExtended) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF642ad3eEncodeGithubComZikwallAppMetricaPkgDomain(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventExtended) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF642ad3eEncodeGithubComZikwallAppMetricaPkgDomain(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventExtended) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF642ad3eDecodeGithubComZikwallAppMetricaPkgDomain(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventExtended) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF642ad3eDecodeGithubComZikwallAppMetricaPkgDomain(l, v)
}
func easyjsonF642ad3eDecodeGithubComZikwallAppMetricaPkgDomain1(in *jlexer.Lexer, out *Event) {
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
func easyjsonF642ad3eEncodeGithubComZikwallAppMetricaPkgDomain1(out *jwriter.Writer, in Event) {
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Event) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF642ad3eEncodeGithubComZikwallAppMetricaPkgDomain1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Event) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF642ad3eEncodeGithubComZikwallAppMetricaPkgDomain1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Event) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF642ad3eDecodeGithubComZikwallAppMetricaPkgDomain1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Event) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF642ad3eDecodeGithubComZikwallAppMetricaPkgDomain1(l, v)
}
