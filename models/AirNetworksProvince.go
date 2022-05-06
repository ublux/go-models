package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AirNetworksProvince struct {
	DateCreated primitive.DateTime `bson:"dateCreated" json:"dateCreated"`
	DateDeleted primitive.DateTime `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated primitive.DateTime `bson:"dateUpdated" json:"dateUpdated"`
	Id          string             `bson:"id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Populations []string           `bson:"populations" json:"populations"`
}

// Implementing interface IUbluxDocument
func (x AirNetworksProvince) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x AirNetworksProvince) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x AirNetworksProvince) GetDateUpdated() primitive.DateTime {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x AirNetworksProvince) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildAirNetworksProvince(m map[string]interface{}, x *AirNetworksProvince) {
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["name"]; ok && val != nil {
		x.Name = val.(string)
	}
	if val, ok := m["populations"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.Populations = append(x.Populations, val.(string))
				}
			}
		}
	}
}
