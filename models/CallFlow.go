package models

import "time"

type CallFlow struct {
	DateCreated  time.Time `bson:"dateCreated" json:"dateCreated"`
	DateDeleted  time.Time `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated  time.Time `bson:"dateUpdated" json:"dateUpdated"`
	Description  string    `bson:"description" json:"description"`
	FriendlyName string    `bson:"friendlyName" json:"friendlyName"`
	Id           string    `bson:"id" json:"id"`
	IdAccount    string    `bson:"idAccount" json:"idAccount"`
	XmlTree      string    `bson:"xmlTree" json:"xmlTree"`
}

// Implementing interface IUbluxDocument
func (x CallFlow) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x CallFlow) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x CallFlow) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x CallFlow) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x CallFlow) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildCallFlow(m map[string]interface{}, x *CallFlow) {
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(time.Time)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(time.Time)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(time.Time)
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
	if val, ok := m["xmlTree"]; ok && val != nil {
		x.XmlTree = val.(string)
	}
}
