package models

import (
	. "github.com/ublux/go-models/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CloudServicePbx struct {
	Id               string           `bson:"_id" json:"id"`
	CloudServiceType CloudServiceType `bson:"cloudServiceType" json:"cloudServiceType"`
	CountryIsoCode   CountryIsoCode   `bson:"countryIsoCode" json:"countryIsoCode"`
	DateCreated      string           `bson:"dateCreated" json:"dateCreated"`
	DateDeleted      string           `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated      string           `bson:"dateUpdated" json:"dateUpdated"`
	ExternalIps      []string         `bson:"externalIps" json:"externalIps"`
	IdIdentity       string           `bson:"idIdentity" json:"idIdentity"`
	IsFailover       bool             `bson:"isFailover" json:"isFailover"`
	IsHealthy        bool             `bson:"isHealthy" json:"isHealthy"`
	Localnet         string           `bson:"localnet" json:"localnet"`
	NAT              bool             `bson:"nAT" json:"nAT"`
}

// Implementing interface IUbluxDocument
func (x CloudServicePbx) GetDateCreated() string {
	return x.DateCreated
}
func (x CloudServicePbx) GetDateDeleted() string {
	return x.DateDeleted
}
func (x CloudServicePbx) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x CloudServicePbx) GetId() string {
	return x.Id
}

// Implementing interface CloudService
func (x CloudServicePbx) GetIdIdentity() string {
	return x.IdIdentity
}
func (x CloudServicePbx) GetCloudServiceType() CloudServiceType {
	return x.CloudServiceType
}
func (x CloudServicePbx) GetCountryIsoCode() CountryIsoCode {
	return x.CountryIsoCode
}
func (x CloudServicePbx) GetLocalnet() string {
	return x.Localnet
}
func (x CloudServicePbx) GetExternalIps() []string {
	return x.ExternalIps
}
func (x CloudServicePbx) GetIsFailover() bool {
	return x.IsFailover
}
func (x CloudServicePbx) GetNAT() bool {
	return x.NAT
}
func (x CloudServicePbx) GetIsHealthy() bool {
	return x.IsHealthy
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildCloudServicePbx(m map[string]interface{}, x *CloudServicePbx) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	x.CloudServiceType = CloudServiceType_PBX // readonly property
	if val, ok := m["countryIsoCode"]; ok && val != nil {
		x.CountryIsoCode = CountryIsoCode("CountryIsoCode_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(string)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(string)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(string)
	}
	if val, ok := m["externalIps"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.ExternalIps = append(x.ExternalIps, val.(string))
				}
			}
		}
	}
	if val, ok := m["idIdentity"]; ok && val != nil {
		x.IdIdentity = val.(string)
	}
	if val, ok := m["isFailover"]; ok && val != nil {
		x.IsFailover = val.(bool)
	}
	if val, ok := m["isHealthy"]; ok && val != nil {
		x.IsHealthy = val.(bool)
	}
	if val, ok := m["localnet"]; ok && val != nil {
		x.Localnet = val.(string)
	}
	if val, ok := m["nAT"]; ok && val != nil {
		x.NAT = val.(bool)
	}
}
