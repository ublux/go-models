package models

type IP struct {
	Id          string `bson:"_id" json:"id"`
	DateCreated string `bson:"dateCreated" json:"dateCreated"`
	DateDeleted string `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated string `bson:"dateUpdated" json:"dateUpdated"`
	IdIdentity  string `bson:"idIdentity" json:"idIdentity"`
	IpOrigin    string `bson:"ipOrigin" json:"ipOrigin"`
	IsBlack     bool   `bson:"isBlack" json:"isBlack"`
}

// Implementing interface IUbluxDocument
func (x IP) GetDateCreated() string {
	return x.DateCreated
}
func (x IP) GetDateDeleted() string {
	return x.DateDeleted
}
func (x IP) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x IP) GetId() string {
	return x.Id
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildIP(m map[string]interface{}, x *IP) {
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
	if val, ok := m["idIdentity"]; ok && val != nil {
		x.IdIdentity = val.(string)
	}
	if val, ok := m["ipOrigin"]; ok && val != nil {
		x.IpOrigin = val.(string)
	}
	if val, ok := m["isBlack"]; ok && val != nil {
		x.IsBlack = val.(bool)
	}
}
