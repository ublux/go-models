package models

type SMS struct {
	Id                      string  `bson:"_id" json:"id"`
	Body                    string  `bson:"body" json:"body"`
	Contact                 Contact `bson:"contact" json:"contact"`
	DateCreated             string  `bson:"dateCreated" json:"dateCreated"`
	DateDeleted             string  `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated             string  `bson:"dateUpdated" json:"dateUpdated"`
	FromInternationalFormat string  `bson:"fromInternationalFormat" json:"fromInternationalFormat"`
	IdAccount               string  `bson:"idAccount" json:"idAccount"`
	IdVoipNumber            string  `bson:"idVoipNumber" json:"idVoipNumber"`
	IsIncoming              bool    `bson:"isIncoming" json:"isIncoming"`
	NumSegments             int32   `bson:"numSegments" json:"numSegments"`
	Status                  string  `bson:"status" json:"status"`
	ToInternationalFormat   string  `bson:"toInternationalFormat" json:"toInternationalFormat"`
}

// Implementing interface IUbluxDocument
func (x SMS) GetDateCreated() string {
	return x.DateCreated
}
func (x SMS) GetDateDeleted() string {
	return x.DateDeleted
}
func (x SMS) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x SMS) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x SMS) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildSMS(m map[string]interface{}, x *SMS) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["body"]; ok && val != nil {
		x.Body = val.(string)
	}
	if val, ok := m["contact"]; ok && val != nil {
		BuildContact(val.(map[string]interface{}), &x.Contact)
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
	if val, ok := m["fromInternationalFormat"]; ok && val != nil {
		x.FromInternationalFormat = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["idVoipNumber"]; ok && val != nil {
		x.IdVoipNumber = val.(string)
	}
	if val, ok := m["isIncoming"]; ok && val != nil {
		x.IsIncoming = val.(bool)
	}
	if val, ok := m["numSegments"]; ok && val != nil {
		x.NumSegments = val.(int32)
	}
	if val, ok := m["status"]; ok && val != nil {
		x.Status = val.(string)
	}
	if val, ok := m["toInternationalFormat"]; ok && val != nil {
		x.ToInternationalFormat = val.(string)
	}
}
