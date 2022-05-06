package models

import "time"
import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type ExtensionConference struct {
	AnnounceParticipants          bool          `bson:"announceParticipants" json:"announceParticipants"`
	DateCreated                   time.Time     `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                   time.Time     `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                   time.Time     `bson:"dateUpdated" json:"dateUpdated"`
	ExtensionType                 ExtensionType `bson:"extensionType" json:"extensionType"`
	Id                            string        `bson:"id" json:"id"`
	IdAccount                     string        `bson:"idAccount" json:"idAccount"`
	IdMusicOnHoldGroup            string        `bson:"idMusicOnHoldGroup" json:"idMusicOnHoldGroup"`
	IdsAudiosWhenOneParticipant   []string      `bson:"idsAudiosWhenOneParticipant" json:"idsAudiosWhenOneParticipant"`
	InjectExtensionNameToCallerId bool          `bson:"injectExtensionNameToCallerId" json:"injectExtensionNameToCallerId"`
	Number                        string        `bson:"number" json:"number"`
	Pin                           string        `bson:"pin" json:"pin"`
}

// Implementing interface IUbluxDocument
func (x ExtensionConference) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x ExtensionConference) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x ExtensionConference) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x ExtensionConference) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x ExtensionConference) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface Extension
func (x ExtensionConference) GetIdMusicOnHoldGroup() string {
	return x.IdMusicOnHoldGroup
}
func (x ExtensionConference) GetExtensionType() ExtensionType {
	return x.ExtensionType
}
func (x ExtensionConference) GetNumber() string {
	return x.Number
}
func (x ExtensionConference) GetInjectExtensionNameToCallerId() bool {
	return x.InjectExtensionNameToCallerId
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildExtensionConference(m map[string]interface{}, x *ExtensionConference) {
	if val, ok := m["announceParticipants"]; ok && val != nil {
		x.AnnounceParticipants = val.(bool)
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
	x.ExtensionType = ExtensionType_Conference // readonly property
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup"]; ok && val != nil {
		x.IdMusicOnHoldGroup = val.(string)
	}
	if val, ok := m["idsAudiosWhenOneParticipant"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.IdsAudiosWhenOneParticipant = append(x.IdsAudiosWhenOneParticipant, val.(string))
				}
			}
		}
	}
	if val, ok := m["injectExtensionNameToCallerId"]; ok && val != nil {
		x.InjectExtensionNameToCallerId = val.(bool)
	}
	if val, ok := m["number"]; ok && val != nil {
		x.Number = val.(string)
	}
	if val, ok := m["pin"]; ok && val != nil {
		x.Pin = val.(string)
	}
}
