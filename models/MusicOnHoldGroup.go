package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MusicOnHoldGroup struct {
	DateCreated  primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	DateDeleted  primitive.DateTime `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated  primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	Description  string             `bson:"description" json:"description"`
	FriendlyName string             `bson:"friendlyName" json:"friendlyName"`
	Id           string             `bson:"id" json:"id"`
	IdAccount    string             `bson:"idAccount" json:"idAccount"`
	IdsAudios    []string           `bson:"idsAudios" json:"idsAudios"`
}

// Implementing interface IUbluxDocument
func (x MusicOnHoldGroup) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x MusicOnHoldGroup) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x MusicOnHoldGroup) GetDateUpdated() primitive.DateTime {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x MusicOnHoldGroup) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x MusicOnHoldGroup) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildMusicOnHoldGroup(m map[string]interface{}, x *MusicOnHoldGroup) {
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
	}
	if val, ok := m["description"]; ok && val != nil {
		x.Description = val.(string)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["idsAudios"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.IdsAudios = append(x.IdsAudios, val.(string))
				}
			}
		}
	}
}
