package models

type UnauthorizedToken struct {
	Id             string `bson:"_id" json:"id"`
	DateCreated    string `bson:"dateCreated" json:"dateCreated"`
	DateDeleted    string `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated    string `bson:"dateUpdated" json:"dateUpdated"`
	ExpirationDate string `bson:"expirationDate" json:"expirationDate"`
	IdAccount      string `bson:"idAccount" json:"idAccount"`
	IdIdentity     string `bson:"idIdentity" json:"idIdentity"`
	Jwt            string `bson:"jwt" json:"jwt"`
}

// Implementing interface IUbluxDocument
func (x UnauthorizedToken) GetDateCreated() string {
	return x.DateCreated
}
func (x UnauthorizedToken) GetDateDeleted() string {
	return x.DateDeleted
}
func (x UnauthorizedToken) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x UnauthorizedToken) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x UnauthorizedToken) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildUnauthorizedToken(m map[string]interface{}, x *UnauthorizedToken) {
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
	if val, ok := m["expirationDate"]; ok && val != nil {
		x.ExpirationDate = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["idIdentity"]; ok && val != nil {
		x.IdIdentity = val.(string)
	}
	if val, ok := m["jwt"]; ok && val != nil {
		x.Jwt = val.(string)
	}
}
