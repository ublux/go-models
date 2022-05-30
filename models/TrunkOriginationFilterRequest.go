package models

import "go.mongodb.org/mongo-driver/bson/primitive"
import . "github.com/ublux/go-models/enums"

type TrunkOriginationFilterRequest struct {
	DateCreated_EQ          primitive.DateTime   `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE         primitive.DateTime   `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE         primitive.DateTime   `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ          primitive.DateTime   `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE         primitive.DateTime   `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE         primitive.DateTime   `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	Id_CON                  string               `bson:"id_CON" json:"id_CON"`
	Id_EQ                   string               `bson:"id_EQ" json:"id_EQ"`
	Id_REG                  string               `bson:"id_REG" json:"id_REG"`
	TrunkOriginationType_EQ TrunkOriginationType `bson:"trunkOriginationType_EQ" json:"trunkOriginationType_EQ"`
}

// BUILDER from bson map:
func BuildTrunkOriginationFilterRequest(m map[string]interface{}, x *TrunkOriginationFilterRequest) {
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
	if val, ok := m["trunkOriginationType_EQ"]; ok && val != nil {
		x.TrunkOriginationType_EQ = TrunkOriginationType("TrunkOriginationType_EQ_" + val.(string))
	} // is NOT readonly obtained from map
}
