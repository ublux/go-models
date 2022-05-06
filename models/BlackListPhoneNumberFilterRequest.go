package models

import "time"

type BlackListPhoneNumberFilterRequest struct {
	DateCreated_EQ                   time.Time `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                  time.Time `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                  time.Time `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                   time.Time `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                  time.Time `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                  time.Time `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	FriendlyName_CON                 string    `bson:"friendlyName_CON" json:"friendlyName_CON"`
	FriendlyName_EQ                  string    `bson:"friendlyName_EQ" json:"friendlyName_EQ"`
	FriendlyName_REG                 string    `bson:"friendlyName_REG" json:"friendlyName_REG"`
	Id_CON                           string    `bson:"id_CON" json:"id_CON"`
	Id_EQ                            string    `bson:"id_EQ" json:"id_EQ"`
	Id_REG                           string    `bson:"id_REG" json:"id_REG"`
	IdAudioToPlayIfCallIsBlocked_CON string    `bson:"idAudioToPlayIfCallIsBlocked_CON" json:"idAudioToPlayIfCallIsBlocked_CON"`
	IdAudioToPlayIfCallIsBlocked_EQ  string    `bson:"idAudioToPlayIfCallIsBlocked_EQ" json:"idAudioToPlayIfCallIsBlocked_EQ"`
	IdAudioToPlayIfCallIsBlocked_REG string    `bson:"idAudioToPlayIfCallIsBlocked_REG" json:"idAudioToPlayIfCallIsBlocked_REG"`
	PhoneNumber_CON                  string    `bson:"phoneNumber_CON" json:"phoneNumber_CON"`
	PhoneNumber_EQ                   string    `bson:"phoneNumber_EQ" json:"phoneNumber_EQ"`
	PhoneNumber_REG                  string    `bson:"phoneNumber_REG" json:"phoneNumber_REG"`
}

// BUILDER from bson map:
func BuildBlackListPhoneNumberFilterRequest(m map[string]interface{}, x *BlackListPhoneNumberFilterRequest) {
	if val, ok := m["dateCreated_EQ"]; ok && val != nil {
		x.DateCreated_EQ = val.(time.Time)
	}
	if val, ok := m["dateCreated_GTE"]; ok && val != nil {
		x.DateCreated_GTE = val.(time.Time)
	}
	if val, ok := m["dateCreated_LTE"]; ok && val != nil {
		x.DateCreated_LTE = val.(time.Time)
	}
	if val, ok := m["dateUpdated_EQ"]; ok && val != nil {
		x.DateUpdated_EQ = val.(time.Time)
	}
	if val, ok := m["dateUpdated_GTE"]; ok && val != nil {
		x.DateUpdated_GTE = val.(time.Time)
	}
	if val, ok := m["dateUpdated_LTE"]; ok && val != nil {
		x.DateUpdated_LTE = val.(time.Time)
	}
	if val, ok := m["friendlyName_CON"]; ok && val != nil {
		x.FriendlyName_CON = val.(string)
	}
	if val, ok := m["friendlyName_EQ"]; ok && val != nil {
		x.FriendlyName_EQ = val.(string)
	}
	if val, ok := m["friendlyName_REG"]; ok && val != nil {
		x.FriendlyName_REG = val.(string)
	}
	if val, ok := m["id_CON"]; ok && val != nil {
		x.Id_CON = val.(string)
	}
	if val, ok := m["id_EQ"]; ok && val != nil {
		x.Id_EQ = val.(string)
	}
	if val, ok := m["id_REG"]; ok && val != nil {
		x.Id_REG = val.(string)
	}
	if val, ok := m["idAudioToPlayIfCallIsBlocked_CON"]; ok && val != nil {
		x.IdAudioToPlayIfCallIsBlocked_CON = val.(string)
	}
	if val, ok := m["idAudioToPlayIfCallIsBlocked_EQ"]; ok && val != nil {
		x.IdAudioToPlayIfCallIsBlocked_EQ = val.(string)
	}
	if val, ok := m["idAudioToPlayIfCallIsBlocked_REG"]; ok && val != nil {
		x.IdAudioToPlayIfCallIsBlocked_REG = val.(string)
	}
	if val, ok := m["phoneNumber_CON"]; ok && val != nil {
		x.PhoneNumber_CON = val.(string)
	}
	if val, ok := m["phoneNumber_EQ"]; ok && val != nil {
		x.PhoneNumber_EQ = val.(string)
	}
	if val, ok := m["phoneNumber_REG"]; ok && val != nil {
		x.PhoneNumber_REG = val.(string)
	}
}
