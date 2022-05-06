package models

import . "github.com/ublux/go-models/enums"
import "time"
import "go.mongodb.org/mongo-driver/bson/primitive"

type TrunkTermination struct {
	CountryIsoCodesThatCanCall []CountryIsoCode `bson:"countryIsoCodesThatCanCall" json:"countryIsoCodesThatCanCall"`
	DateCreated                time.Time        `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                time.Time        `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                time.Time        `bson:"dateUpdated" json:"dateUpdated"`
	FriendlyName               string           `bson:"friendlyName" json:"friendlyName"`
	Id                         string           `bson:"id" json:"id"`
	IdTrunkTerminationExisting string           `bson:"idTrunkTerminationExisting" json:"idTrunkTerminationExisting"`
	IdVoipProvider             string           `bson:"idVoipProvider" json:"idVoipProvider"`
	Password                   string           `bson:"password" json:"password"`
	SID                        string           `bson:"sID" json:"sID"`
	TerminationUri             string           `bson:"terminationUri" json:"terminationUri"`
	Username                   string           `bson:"username" json:"username"`
}

// Implementing interface IUbluxDocument
func (x TrunkTermination) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x TrunkTermination) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x TrunkTermination) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x TrunkTermination) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildTrunkTermination(m map[string]interface{}, x *TrunkTermination) {
	if val, ok := m["countryIsoCodesThatCanCall"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				x.CountryIsoCodesThatCanCall = append(x.CountryIsoCodesThatCanCall, CountryIsoCode("CountryIsoCodesThatCanCall_"+val.(string)))
			} // is NOT readonly obtained from map
		}
	}
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(time.Time)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(time.Time)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(time.Time)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idTrunkTerminationExisting"]; ok && val != nil {
		x.IdTrunkTerminationExisting = val.(string)
	}
	if val, ok := m["idVoipProvider"]; ok && val != nil {
		x.IdVoipProvider = val.(string)
	}
	if val, ok := m["password"]; ok && val != nil {
		x.Password = val.(string)
	}
	if val, ok := m["sID"]; ok && val != nil {
		x.SID = val.(string)
	}
	if val, ok := m["terminationUri"]; ok && val != nil {
		x.TerminationUri = val.(string)
	}
	if val, ok := m["username"]; ok && val != nil {
		x.Username = val.(string)
	}
}
