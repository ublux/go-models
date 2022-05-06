package models

import "time"
import "go.mongodb.org/mongo-driver/bson/primitive"

type Phone struct {
	DateCreated          time.Time `bson:"dateCreated" json:"dateCreated"`
	DateDeleted          time.Time `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated          time.Time `bson:"dateUpdated" json:"dateUpdated"`
	FriendlyName         string    `bson:"friendlyName" json:"friendlyName"`
	Id                   string    `bson:"id" json:"id"`
	IdAccount            string    `bson:"idAccount" json:"idAccount"`
	IdCloudServicePbx    string    `bson:"idCloudServicePbx" json:"idCloudServicePbx"`
	IdIdentityWebApp     string    `bson:"idIdentityWebApp" json:"idIdentityWebApp"`
	IdPhoneConfiguration string    `bson:"idPhoneConfiguration" json:"idPhoneConfiguration"`
	Lines                []Line    `bson:"lines" json:"lines"`
}

// Implementing interface IUbluxDocument
func (x Phone) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x Phone) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x Phone) GetDateUpdated() time.Time {
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
