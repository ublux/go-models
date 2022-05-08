package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type CloudServiceWebHost struct {
	CloudServiceType CloudServiceType   `bson:"cloudServiceType" json:"cloudServiceType"`
	CountryIsoCode   CountryIsoCode     `bson:"countryIsoCode" json:"countryIsoCode"`
	DateCreated      primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	DateDeleted      primitive.DateTime `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated      primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	ExternalIps      []string           `bson:"externalIps" json:"externalIps"`
	Id               string             `bson:"_id" json:"id"`
	IdIdentity       string             `bson:"idIdentity" json:"idIdentity"`
	IsFailover       bool               `bson:"isFailover" json:"isFailover"`
	IsHealthy        bool               `bson:"isHealthy" json:"isHealthy"`
	Localnet         string             `bson:"localnet" json:"localnet"`
	Nat              bool               `bson:"nat" json:"nat"`
}

// Implementing interface IUbluxDocument
func (x CloudServiceWebHost) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x CloudServiceWebHost) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x CloudServiceWebHost) GetDateUpdated() primitive.DateTime {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x CloudServiceWebHost) GetId() string {
	return x.Id
}

// Implementing interface CloudService
func (x CloudServiceWebHost) GetIdIdentity() string {
	return x.IdIdentity
}
func (x CloudServiceWebHost) GetCloudServiceType() CloudServiceType {
	return x.CloudServiceType
}
func (x CloudServiceWebHost) GetCountryIsoCode() CountryIsoCode {
	return x.CountryIsoCode
}
func (x CloudServiceWebHost) GetLocalnet() string {
	return x.Localnet
}
func (x CloudServiceWebHost) GetExternalIps() []string {
	return x.ExternalIps
}
func (x CloudServiceWebHost) GetIsFailover() bool {
	return x.IsFailover
}
func (x CloudServiceWebHost) GetNat() bool {
	return x.Nat
}
func (x CloudServiceWebHost) GetIsHealthy() bool {
	return x.IsHealthy
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildCloudServiceWebHost(m map[string]interface{}, x *CloudServiceWebHost) {
	x.CloudServiceType = CloudServiceType_WH // readonly property
	if val, ok := m["countryIsoCode"]; ok && val != nil {
		x.CountryIsoCode = CountryIsoCode("CountryIsoCode_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
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
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
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
	if val, ok := m["nat"]; ok && val != nil {
		x.Nat = val.(bool)
	}
}
