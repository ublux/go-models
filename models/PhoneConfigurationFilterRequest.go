package models

import "time"

type PhoneConfigurationFilterRequest struct {
	DateCreated_EQ     time.Time `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE    time.Time `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE    time.Time `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ     time.Time `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE    time.Time `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE    time.Time `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	Description_CON    string    `bson:"description_CON" json:"description_CON"`
	Description_EQ     string    `bson:"description_EQ" json:"description_EQ"`
	Description_REG    string    `bson:"description_REG" json:"description_REG"`
	FrienlyName_CON    string    `bson:"frienlyName_CON" json:"frienlyName_CON"`
	FrienlyName_EQ     string    `bson:"frienlyName_EQ" json:"frienlyName_EQ"`
	FrienlyName_REG    string    `bson:"frienlyName_REG" json:"frienlyName_REG"`
	Id_CON             string    `bson:"id_CON" json:"id_CON"`
	Id_EQ              string    `bson:"id_EQ" json:"id_EQ"`
	Id_REG             string    `bson:"id_REG" json:"id_REG"`
	IdLineKeyGroup_CON string    `bson:"idLineKeyGroup_CON" json:"idLineKeyGroup_CON"`
	IdLineKeyGroup_EQ  string    `bson:"idLineKeyGroup_EQ" json:"idLineKeyGroup_EQ"`
	IdLineKeyGroup_REG string    `bson:"idLineKeyGroup_REG" json:"idLineKeyGroup_REG"`
}

// BUILDER from bson map:
func BuildPhoneConfigurationFilterRequest(m map[string]interface{}, x *PhoneConfigurationFilterRequest) {
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
	if val, ok := m["description_CON"]; ok && val != nil {
		x.Description_CON = val.(string)
	}
	if val, ok := m["description_EQ"]; ok && val != nil {
		x.Description_EQ = val.(string)
	}
	if val, ok := m["description_REG"]; ok && val != nil {
		x.Description_REG = val.(string)
	}
	if val, ok := m["frienlyName_CON"]; ok && val != nil {
		x.FrienlyName_CON = val.(string)
	}
	if val, ok := m["frienlyName_EQ"]; ok && val != nil {
		x.FrienlyName_EQ = val.(string)
	}
	if val, ok := m["frienlyName_REG"]; ok && val != nil {
		x.FrienlyName_REG = val.(string)
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
	if val, ok := m["idLineKeyGroup_CON"]; ok && val != nil {
		x.IdLineKeyGroup_CON = val.(string)
	}
	if val, ok := m["idLineKeyGroup_EQ"]; ok && val != nil {
		x.IdLineKeyGroup_EQ = val.(string)
	}
	if val, ok := m["idLineKeyGroup_REG"]; ok && val != nil {
		x.IdLineKeyGroup_REG = val.(string)
	}
}
