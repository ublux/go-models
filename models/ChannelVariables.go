package models

import . "github.com/ublux/go-models/enums"

type ChannelVariables struct {
	CallerIdName   string   `bson:"callerIdName" json:"callerIdName"`
	CallerIdNumber string   `bson:"callerIdNumber" json:"callerIdNumber"`
	IdMusicOnHold  string   `bson:"idMusicOnHold" json:"idMusicOnHold"`
	Language       Language `bson:"language" json:"language"`
}

// BUILDER from bson map:
func BuildChannelVariables(m map[string]interface{}, x *ChannelVariables) {
	if val, ok := m["callerIdName"]; ok && val != nil {
		x.CallerIdName = val.(string)
	}
	if val, ok := m["callerIdNumber"]; ok && val != nil {
		x.CallerIdNumber = val.(string)
	}
	if val, ok := m["idMusicOnHold"]; ok && val != nil {
		x.IdMusicOnHold = val.(string)
	}
	if val, ok := m["language"]; ok && val != nil {
		x.Language = Language("Language_" + val.(string))
	} // is NOT readonly obtained from map
}
