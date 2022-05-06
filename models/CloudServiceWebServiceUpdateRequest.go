package models

type CloudServiceWebServiceUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x CloudServiceWebServiceUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildCloudServiceWebServiceUpdateRequest(m map[string]interface{}, x *CloudServiceWebServiceUpdateRequest) {
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
