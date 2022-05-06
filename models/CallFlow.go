package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CallFlow struct {
	DateCreated  primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	DateDeleted  primitive.DateTime `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated  primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	Description  string             `bson:"description" json:"description"`
	FriendlyName string             `bson:"friendlyName" json:"friendlyName"`
	Id           string             `bson:"id" json:"id"`
	IdAccount    string             `bson:"idAccount" json:"idAccount"`
	XmlTree      string             `bson:"xmlTree" json:"xmlTree"`
}

// Implementing interface IUbluxDocument
func (x CallFlow) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x CallFlow) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x CallFlow) GetDateUpdated() primitive.DateTime {
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
	if val, ok := m["xmlTree"]; ok && val != nil {
		x.XmlTree = val.(string)
	}
}
