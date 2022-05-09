package models

type FaxOutgoing struct {
	EmailAttachment    EmailAttachment `bson:"emailAttachment" json:"emailAttachment"`
	ErrorMessage       string          `bson:"errorMessage" json:"errorMessage"`
	FaxStatus          string          `bson:"faxStatus" json:"faxStatus"`
	Id                 string          `bson:"id" json:"id"`
	IdFaxOutgoingGroup string          `bson:"idFaxOutgoingGroup" json:"idFaxOutgoingGroup"`
	IsComplete         bool            `bson:"isComplete" json:"isComplete"`
	IsPortrait         bool            `bson:"isPortrait" json:"isPortrait"`
	Name               string          `bson:"name" json:"name"`
	NumberOfAttempts   int32           `bson:"numberOfAttempts" json:"numberOfAttempts"`
	NumberOfPages      int32           `bson:"numberOfPages" json:"numberOfPages"`
	NumberOfPagesSent  int32           `bson:"numberOfPagesSent" json:"numberOfPagesSent"`
	Pdf                StoredFile      `bson:"pdf" json:"pdf"`
	To                 string          `bson:"to" json:"to"`
}

// Implementing interface UbluxSubDocument
func (x FaxOutgoing) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildFaxOutgoing(m map[string]interface{}, x *FaxOutgoing) {
	if val, ok := m["emailAttachment"]; ok && val != nil {
		BuildEmailAttachment(val.(map[string]interface{}), &x.EmailAttachment)
	}
	if val, ok := m["errorMessage"]; ok && val != nil {
		x.ErrorMessage = val.(string)
	}
	if val, ok := m["faxStatus"]; ok && val != nil {
		x.FaxStatus = val.(string)
	}
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idFaxOutgoingGroup"]; ok && val != nil {
		x.IdFaxOutgoingGroup = val.(string)
	}
	if val, ok := m["isComplete"]; ok && val != nil {
		x.IsComplete = val.(bool)
	}
	if val, ok := m["isPortrait"]; ok && val != nil {
		x.IsPortrait = val.(bool)
	}
	if val, ok := m["name"]; ok && val != nil {
		x.Name = val.(string)
	}
	if val, ok := m["numberOfAttempts"]; ok && val != nil {
		x.NumberOfAttempts = val.(int32)
	}
	if val, ok := m["numberOfPages"]; ok && val != nil {
		x.NumberOfPages = val.(int32)
	}
	if val, ok := m["numberOfPagesSent"]; ok && val != nil {
		x.NumberOfPagesSent = val.(int32)
	}
	if val, ok := m["pdf"]; ok && val != nil {
		BuildStoredFile(val.(map[string]interface{}), &x.Pdf)
	}
	if val, ok := m["to"]; ok && val != nil {
		x.To = val.(string)
	}
}
