package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Phone struct {
	Id                   string `bson:"_id" json:"id"`
	DateCreated          string `bson:"dateCreated" json:"dateCreated"`
	DateDeleted          string `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated          string `bson:"dateUpdated" json:"dateUpdated"`
	FriendlyName         string `bson:"friendlyName" json:"friendlyName"`
	IdAccount            string `bson:"idAccount" json:"idAccount"`
	IdCloudServicePbx    string `bson:"idCloudServicePbx" json:"idCloudServicePbx"`
	IdIdentityWebApp     string `bson:"idIdentityWebApp" json:"idIdentityWebApp"`
	IdPhoneConfiguration string `bson:"idPhoneConfiguration" json:"idPhoneConfiguration"`
	Lines                []Line `bson:"lines" json:"lines"`
}

// Implementing interface IUbluxDocument
func (x Phone) GetDateCreated() string {
	return x.DateCreated
}
func (x Phone) GetDateDeleted() string {
	return x.DateDeleted
}
func (x Phone) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x Phone) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x Phone) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildPhone(m map[string]interface{}, x *Phone) {
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
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["idCloudServicePbx"]; ok && val != nil {
		x.IdCloudServicePbx = val.(string)
	}
	if val, ok := m["idIdentityWebApp"]; ok && val != nil {
		x.IdIdentityWebApp = val.(string)
	}
	if val, ok := m["idPhoneConfiguration"]; ok && val != nil {
		x.IdPhoneConfiguration = val.(string)
	}
	if val, ok := m["lines"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok {
			for _, val = range array {
				if val != nil {
					tmp := new(Line)
					BuildLine(val.(map[string]interface{}), tmp)
					x.Lines = append(x.Lines, *tmp)
				}
			}
		}
	}
}
