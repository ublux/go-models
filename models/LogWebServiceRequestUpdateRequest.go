package models

type LogWebServiceRequestUpdateRequest struct {
	Id string `bson:"id" json:"id"`
}

// Implementing interface IUbluxDocumentId
func (x LogWebServiceRequestUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildLogWebServiceRequestUpdateRequest(m map[string]interface{}, x *LogWebServiceRequestUpdateRequest) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
}
