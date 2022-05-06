package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AirNetworksProvinceFilterRequest struct {
	DateCreated_EQ  primitive.DateTime `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE primitive.DateTime `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE primitive.DateTime `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ  primitive.DateTime `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE primitive.DateTime `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE primitive.DateTime `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	Id_CON          string             `bson:"id_CON" json:"id_CON"`
	Id_EQ           string             `bson:"id_EQ" json:"id_EQ"`
	Id_REG          string             `bson:"id_REG" json:"id_REG"`
	Name_CON        string             `bson:"name_CON" json:"name_CON"`
	Name_EQ         string             `bson:"name_EQ" json:"name_EQ"`
	Name_REG        string             `bson:"name_REG" json:"name_REG"`
	Populations_CON string             `bson:"populations_CON" json:"populations_CON"`
	Populations_EQ  string             `bson:"populations_EQ" json:"populations_EQ"`
	Populations_REG string             `bson:"populations_REG" json:"populations_REG"`
}

// BUILDER from bson map:
func BuildAirNetworksProvinceFilterRequest(m map[string]interface{}, x *AirNetworksProvinceFilterRequest) {
	if val, ok := m["dateCreated_EQ"]; ok && val != nil {
		x.DateCreated_EQ = val.(primitive.DateTime)
	}
	if val, ok := m["dateCreated_GTE"]; ok && val != nil {
		x.DateCreated_GTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateCreated_LTE"]; ok && val != nil {
		x.DateCreated_LTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated_EQ"]; ok && val != nil {
		x.DateUpdated_EQ = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated_GTE"]; ok && val != nil {
		x.DateUpdated_GTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated_LTE"]; ok && val != nil {
		x.DateUpdated_LTE = val.(primitive.DateTime)
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
	if val, ok := m["name_CON"]; ok && val != nil {
		x.Name_CON = val.(string)
	}
	if val, ok := m["name_EQ"]; ok && val != nil {
		x.Name_EQ = val.(string)
	}
	if val, ok := m["name_REG"]; ok && val != nil {
		x.Name_REG = val.(string)
	}
	if val, ok := m["populations_CON"]; ok && val != nil {
		x.Populations_CON = val.(string)
	}
	if val, ok := m["populations_EQ"]; ok && val != nil {
		x.Populations_EQ = val.(string)
	}
	if val, ok := m["populations_REG"]; ok && val != nil {
		x.Populations_REG = val.(string)
	}
}
