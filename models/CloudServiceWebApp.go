package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type CloudServiceWebApp struct {
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
func (x CloudServiceWebApp) GetDateCreated() string {
	return x.DateCreated
}
func (x CloudServiceWebApp) GetDateDeleted() string {
	return x.DateDeleted
}
func (x CloudServiceWebApp) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x CloudServiceWebApp) GetId() string {
	return x.Id
}

// Implementing interface CloudService
func (x CloudServiceWebApp) GetIdIdentity() string {
	return x.IdIdentity
}
func (x CloudServiceWebApp) GetCloudServiceType() CloudServiceType {
	return x.CloudServiceType
}
func (x CloudServiceWebApp) GetCountryIsoCode() CountryIsoCode {
	return x.CountryIsoCode
}
func (x CloudServiceWebApp) GetLocalnet() string {
	return x.Localnet
}
func (x CloudServiceWebApp) GetExternalIps() []string {
	return x.ExternalIps
}
func (x CloudServiceWebApp) GetIsFailover() bool {
	return x.IsFailover
}
func (x CloudServiceWebApp) GetNAT() bool {
	return x.NAT
}
func (x CloudServiceWebApp) GetIsHealthy() bool {
	return x.IsHealthy
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildCloudServiceWebApp(m map[string]interface{}, x *CloudServiceWebApp) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	x.CloudServiceType = CloudServiceType_WA // readonly property
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
