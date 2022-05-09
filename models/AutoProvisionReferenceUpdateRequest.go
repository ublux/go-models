package models

type AutoProvisionReferenceUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x AutoProvisionReferenceUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildAutoProvisionReferenceUpdateRequest(m map[string]interface{}, x *AutoProvisionReferenceUpdateRequest) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
