package models

type UnauthorizedTokenUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x UnauthorizedTokenUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildUnauthorizedTokenUpdateRequest(m map[string]interface{}, x *UnauthorizedTokenUpdateRequest) {
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
