package models

type FaxIncomingUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x FaxIncomingUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildFaxIncomingUpdateRequest(m map[string]interface{}, x *FaxIncomingUpdateRequest) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
