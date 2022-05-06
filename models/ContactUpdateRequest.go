package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ContactUpdateRequest struct {
	Company        string          `bson:"company" json:"company"`
	ContactEmails  []ContactEmail  `bson:"contactEmails" json:"contactEmails"`
	ContactNumbers []ContactNumber `bson:"contactNumbers" json:"contactNumbers"`
	FirstName      string          `bson:"firstName" json:"firstName"`
	Id             string          `bson:"id" json:"id"`
	JobTittle      string          `bson:"jobTittle" json:"jobTittle"`
	LastName       string          `bson:"lastName" json:"lastName"`
	Notes          string          `bson:"notes" json:"notes"`
	Variables      []Variable      `bson:"variables" json:"variables"`
}

// Implementing interface IUbluxDocumentId
func (x ContactUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildContactUpdateRequest(m map[string]interface{}, x *ContactUpdateRequest) {
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
	if val, ok := m["firstName"]; ok && val != nil {
		x.FirstName = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
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
