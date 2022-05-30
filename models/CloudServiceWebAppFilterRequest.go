package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type CloudServiceWebAppFilterRequest struct {
	CloudServiceType_EQ CloudServiceType   `bson:"cloudServiceType_EQ" json:"cloudServiceType_EQ"`
	CountryIsoCode_EQ   CountryIsoCode     `bson:"countryIsoCode_EQ" json:"countryIsoCode_EQ"`
	DateCreated_EQ      primitive.DateTime `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE     primitive.DateTime `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE     primitive.DateTime `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ      primitive.DateTime `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE     primitive.DateTime `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE     primitive.DateTime `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	ExternalIps_CON     string             `bson:"externalIps_CON" json:"externalIps_CON"`
	ExternalIps_EQ      string             `bson:"externalIps_EQ" json:"externalIps_EQ"`
	ExternalIps_REG     string             `bson:"externalIps_REG" json:"externalIps_REG"`
	Id_CON              string             `bson:"id_CON" json:"id_CON"`
	Id_EQ               string             `bson:"id_EQ" json:"id_EQ"`
	Id_REG              string             `bson:"id_REG" json:"id_REG"`
	IsFailover_EQ       bool               `bson:"isFailover_EQ" json:"isFailover_EQ"`
	IsHealthy_EQ        bool               `bson:"isHealthy_EQ" json:"isHealthy_EQ"`
	Localnet_CON        string             `bson:"localnet_CON" json:"localnet_CON"`
	Localnet_EQ         string             `bson:"localnet_EQ" json:"localnet_EQ"`
	Localnet_REG        string             `bson:"localnet_REG" json:"localnet_REG"`
	Nat_EQ              bool               `bson:"nat_EQ" json:"nat_EQ"`
}

// BUILDER from bson map:
func BuildCloudServiceWebAppFilterRequest(m map[string]interface{}, x *CloudServiceWebAppFilterRequest) {
	if val, ok := m["cloudServiceType_EQ"]; ok && val != nil {
		x.CloudServiceType_EQ = CloudServiceType("CloudServiceType_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["countryIsoCode_EQ"]; ok && val != nil {
		x.CountryIsoCode_EQ = CountryIsoCode("CountryIsoCode_EQ_" + val.(string))
	} // is NOT readonly obtained from map
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
	if val, ok := m["externalIps_CON"]; ok && val != nil {
		x.ExternalIps_CON = val.(string)
	}
	if val, ok := m["externalIps_EQ"]; ok && val != nil {
		x.ExternalIps_EQ = val.(string)
	}
	if val, ok := m["externalIps_REG"]; ok && val != nil {
		x.ExternalIps_REG = val.(string)
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
	if val, ok := m["isFailover_EQ"]; ok && val != nil {
		x.IsFailover_EQ = val.(bool)
	}
	if val, ok := m["isHealthy_EQ"]; ok && val != nil {
		x.IsHealthy_EQ = val.(bool)
	}
	if val, ok := m["localnet_CON"]; ok && val != nil {
		x.Localnet_CON = val.(string)
	}
	if val, ok := m["localnet_EQ"]; ok && val != nil {
		x.Localnet_EQ = val.(string)
	}
	if val, ok := m["localnet_REG"]; ok && val != nil {
		x.Localnet_REG = val.(string)
	}
	if val, ok := m["nat_EQ"]; ok && val != nil {
		x.Nat_EQ = val.(bool)
	}
}
