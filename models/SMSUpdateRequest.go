package models

type SMSUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x SMSUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildSMSUpdateRequest(m map[string]interface{}, x *SMSUpdateRequest) {
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
