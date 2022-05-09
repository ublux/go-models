package models

type IPUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x IPUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildIPUpdateRequest(m map[string]interface{}, x *IPUpdateRequest) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
