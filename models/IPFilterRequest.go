package models

import "time"

type IPFilterRequest struct {
	DateCreated_EQ  time.Time `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE time.Time `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE time.Time `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ  time.Time `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE time.Time `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE time.Time `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	Id_CON          string    `bson:"id_CON" json:"id_CON"`
	Id_EQ           string    `bson:"id_EQ" json:"id_EQ"`
	Id_REG          string    `bson:"id_REG" json:"id_REG"`
	IsBlack_EQ      bool      `bson:"isBlack_EQ" json:"isBlack_EQ"`
}

// BUILDER from bson map:
func BuildIPFilterRequest(m map[string]interface{}, x *IPFilterRequest) {
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
	if val, ok := m["id_CON"]; ok && val != nil {
		x.Id_CON = val.(string)
	}
	if val, ok := m["id_EQ"]; ok && val != nil {
		x.Id_EQ = val.(string)
	}
	if val, ok := m["id_REG"]; ok && val != nil {
		x.Id_REG = val.(string)
	}
	if val, ok := m["isBlack_EQ"]; ok && val != nil {
		x.IsBlack_EQ = val.(bool)
	}
}
