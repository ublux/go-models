package models

type CloudServiceWebHostUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x CloudServiceWebHostUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildCloudServiceWebHostUpdateRequest(m map[string]interface{}, x *CloudServiceWebHostUpdateRequest) {
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
