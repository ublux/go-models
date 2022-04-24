package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AirNetworksProvince struct {
	Id          string   `bson:"_id" json:"id"`
	DateCreated string   `bson:"dateCreated" json:"dateCreated"`
	DateDeleted string   `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated string   `bson:"dateUpdated" json:"dateUpdated"`
	Name        string   `bson:"name" json:"name"`
	Populations []string `bson:"populations" json:"populations"`
}

// Implementing interface IUbluxDocument
func (x AirNetworksProvince) GetDateCreated() string {
	return x.DateCreated
}
func (x AirNetworksProvince) GetDateDeleted() string {
	return x.DateDeleted
}
func (x AirNetworksProvince) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x AirNetworksProvince) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildAirNetworksProvince(m map[string]interface{}, x *AirNetworksProvince) {
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
