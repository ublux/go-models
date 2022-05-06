package models

import "time"
import . "github.com/ublux/go-models/enums"

type ExtensionCallFlow struct {
	CallFlowLabel                 string        `bson:"callFlowLabel" json:"callFlowLabel"`
	DateCreated                   time.Time     `bson:"dateCreated" json:"dateCreated"`
	DateDeleted                   time.Time     `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated                   time.Time     `bson:"dateUpdated" json:"dateUpdated"`
	ExtensionType                 ExtensionType `bson:"extensionType" json:"extensionType"`
	Id                            string        `bson:"id" json:"id"`
	IdAccount                     string        `bson:"idAccount" json:"idAccount"`
	IdCallFlow                    string        `bson:"idCallFlow" json:"idCallFlow"`
	IdMusicOnHoldGroup            string        `bson:"idMusicOnHoldGroup" json:"idMusicOnHoldGroup"`
	InjectExtensionNameToCallerId bool          `bson:"injectExtensionNameToCallerId" json:"injectExtensionNameToCallerId"`
	Number                        string        `bson:"number" json:"number"`
	TimeZone                      string        `bson:"timeZone" json:"timeZone"`
}

// Implementing interface IUbluxDocument
func (x ExtensionCallFlow) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x ExtensionCallFlow) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x ExtensionCallFlow) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x ExtensionCallFlow) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x ExtensionCallFlow) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface Extension
func (x ExtensionCallFlow) GetIdMusicOnHoldGroup() string {
	return x.IdMusicOnHoldGroup
}
func (x ExtensionCallFlow) GetExtensionType() ExtensionType {
	return x.ExtensionType
}
func (x ExtensionCallFlow) GetNumber() string {
	return x.Number
}
func (x ExtensionCallFlow) GetInjectExtensionNameToCallerId() bool {
	return x.InjectExtensionNameToCallerId
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildExtensionCallFlow(m map[string]interface{}, x *ExtensionCallFlow) {
	if val, ok := m["callFlowLabel"]; ok && val != nil {
		x.CallFlowLabel = val.(string)
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
	x.ExtensionType = ExtensionType_CallFlow // readonly property
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["idCallFlow"]; ok && val != nil {
		x.IdCallFlow = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup"]; ok && val != nil {
		x.IdMusicOnHoldGroup = val.(string)
	}
	if val, ok := m["injectExtensionNameToCallerId"]; ok && val != nil {
		x.InjectExtensionNameToCallerId = val.(bool)
	}
	if val, ok := m["number"]; ok && val != nil {
		x.Number = val.(string)
	}
	if val, ok := m["timeZone"]; ok && val != nil {
		x.TimeZone = val.(string)
	}
}
