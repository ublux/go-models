package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contact struct {
	Company        string             `bson:"company" json:"company"`
	ContactEmails  []ContactEmail     `bson:"contactEmails" json:"contactEmails"`
	ContactNumbers []ContactNumber    `bson:"contactNumbers" json:"contactNumbers"`
	DateCreated    primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	DateDeleted    primitive.DateTime `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated    primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	FirstName      string             `bson:"firstName" json:"firstName"`
	Hash           string             `bson:"hash" json:"hash"`
	Id             string             `bson:"_id" json:"id"`
	IdAccount      string             `bson:"idAccount" json:"idAccount"`
	JobTittle      string             `bson:"jobTittle" json:"jobTittle"`
	LastName       string             `bson:"lastName" json:"lastName"`
	Notes          string             `bson:"notes" json:"notes"`
	Variables      []Variable         `bson:"variables" json:"variables"`
}

// Implementing interface IUbluxDocument
func (x Contact) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x Contact) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x Contact) GetDateUpdated() primitive.DateTime {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x Contact) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x Contact) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildContact(m map[string]interface{}, x *Contact) {
	if val, ok := m["company"]; ok && val != nil {
		x.Company = val.(string)
	}
	if val, ok := m["contactEmails"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok {
			for _, val = range array {
				if val != nil {
					tmp := new(ContactEmail)
					BuildContactEmail(val.(map[string]interface{}), tmp)
					x.ContactEmails = append(x.ContactEmails, *tmp)
				}
			}
		}
	}
	if val, ok := m["contactNumbers"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok {
			for _, val = range array {
				if val != nil {
					tmp := new(ContactNumber)
					BuildContactNumber(val.(map[string]interface{}), tmp)
					x.ContactNumbers = append(x.ContactNumbers, *tmp)
				}
			}
		}
	}
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
	}
	if val, ok := m["firstName"]; ok && val != nil {
		x.FirstName = val.(string)
	}
	if val, ok := m["hash"]; ok && val != nil {
		x.Hash = val.(string)
	}
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["jobTittle"]; ok && val != nil {
		x.JobTittle = val.(string)
	}
	if val, ok := m["lastName"]; ok && val != nil {
		x.LastName = val.(string)
	}
	if val, ok := m["notes"]; ok && val != nil {
		x.Notes = val.(string)
	}
	if val, ok := m["variables"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok {
			for _, val = range array {
				if val != nil {
					tmp := new(Variable)
					BuildVariable(val.(map[string]interface{}), tmp)
					x.Variables = append(x.Variables, *tmp)
				}
			}
		}
	}
}
