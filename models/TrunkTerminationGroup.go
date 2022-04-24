package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TrunkTerminationGroup struct {
	Id                                    string   `bson:"_id" json:"id"`
	DateCreated                           string   `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                           string   `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                           string   `bson:"dateUpdated" json:"dateUpdated"`
	IdsTrunkTerminationsOrderedByPriority []string `bson:"idsTrunkTerminationsOrderedByPriority" json:"idsTrunkTerminationsOrderedByPriority"`
}

// Implementing interface IUbluxDocument
func (x TrunkTerminationGroup) GetDateCreated() string {
	return x.DateCreated
}
func (x TrunkTerminationGroup) GetDateDeleted() string {
	return x.DateDeleted
}
func (x TrunkTerminationGroup) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x TrunkTerminationGroup) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildTrunkTerminationGroup(m map[string]interface{}, x *TrunkTerminationGroup) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(string)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(string)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(string)
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
