package models

import "time"
import "go.mongodb.org/mongo-driver/bson/primitive"

type TrunkTerminationGroup struct {
	DateCreated                           time.Time `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                           time.Time `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                           time.Time `bson:"dateUpdated" json:"dateUpdated"`
	Id                                    string    `bson:"id" json:"id"`
	IdsTrunkTerminationsOrderedByPriority []string  `bson:"idsTrunkTerminationsOrderedByPriority" json:"idsTrunkTerminationsOrderedByPriority"`
}

// Implementing interface IUbluxDocument
func (x TrunkTerminationGroup) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x TrunkTerminationGroup) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x TrunkTerminationGroup) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x TrunkTerminationGroup) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildTrunkTerminationGroup(m map[string]interface{}, x *TrunkTerminationGroup) {
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(time.Time)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(time.Time)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(time.Time)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idsTrunkTerminationsOrderedByPriority"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.IdsTrunkTerminationsOrderedByPriority = append(x.IdsTrunkTerminationsOrderedByPriority, val.(string))
				}
			}
		}
	}
}
