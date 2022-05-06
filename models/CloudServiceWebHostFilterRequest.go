package models

import . "github.com/ublux/go-models/enums"
import "time"

type CloudServiceWebHostFilterRequest struct {
	CloudServiceType_EQ CloudServiceType `bson:"cloudServiceType_EQ" json:"cloudServiceType_EQ"`
	CountryIsoCode_EQ   CountryIsoCode   `bson:"countryIsoCode_EQ" json:"countryIsoCode_EQ"`
	DateCreated_EQ      time.Time        `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE     time.Time        `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE     time.Time        `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ      time.Time        `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE     time.Time        `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE     time.Time        `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	ExternalIps_CON     string           `bson:"externalIps_CON" json:"externalIps_CON"`
	ExternalIps_EQ      string           `bson:"externalIps_EQ" json:"externalIps_EQ"`
	ExternalIps_REG     string           `bson:"externalIps_REG" json:"externalIps_REG"`
	Id_CON              string           `bson:"id_CON" json:"id_CON"`
	Id_EQ               string           `bson:"id_EQ" json:"id_EQ"`
	Id_REG              string           `bson:"id_REG" json:"id_REG"`
	IdIdentity_CON      string           `bson:"idIdentity_CON" json:"idIdentity_CON"`
	IdIdentity_EQ       string           `bson:"idIdentity_EQ" json:"idIdentity_EQ"`
	IdIdentity_REG      string           `bson:"idIdentity_REG" json:"idIdentity_REG"`
	IsFailover_EQ       bool             `bson:"isFailover_EQ" json:"isFailover_EQ"`
	IsHealthy_EQ        bool             `bson:"isHealthy_EQ" json:"isHealthy_EQ"`
	Localnet_CON        string           `bson:"localnet_CON" json:"localnet_CON"`
	Localnet_EQ         string           `bson:"localnet_EQ" json:"localnet_EQ"`
	Localnet_REG        string           `bson:"localnet_REG" json:"localnet_REG"`
	NAT_EQ              bool             `bson:"nAT_EQ" json:"nAT_EQ"`
}

// BUILDER from bson map:
func BuildCloudServiceWebHostFilterRequest(m map[string]interface{}, x *CloudServiceWebHostFilterRequest) {
	if val, ok := m["cloudServiceType_EQ"]; ok && val != nil {
		x.CloudServiceType_EQ = CloudServiceType("CloudServiceType_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["countryIsoCode_EQ"]; ok && val != nil {
		x.CountryIsoCode_EQ = CountryIsoCode("CountryIsoCode_EQ_" + val.(string))
	} // is NOT readonly obtained from map
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
	if val, ok := m["idIdentity_CON"]; ok && val != nil {
		x.IdIdentity_CON = val.(string)
	}
	if val, ok := m["idIdentity_EQ"]; ok && val != nil {
		x.IdIdentity_EQ = val.(string)
	}
	if val, ok := m["idIdentity_REG"]; ok && val != nil {
		x.IdIdentity_REG = val.(string)
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
	if val, ok := m["nAT_EQ"]; ok && val != nil {
		x.NAT_EQ = val.(bool)
	}
}
