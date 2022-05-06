package models

type CallFlowUpdateRequest struct {
	Description  string `bson:"description" json:"description"`
	FriendlyName string `bson:"friendlyName" json:"friendlyName"`
	Id           string `bson:"id" json:"id"`
	XmlTree      string `bson:"xmlTree" json:"xmlTree"`
}

// Implementing interface IUbluxDocumentId
func (x CallFlowUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildCallFlowUpdateRequest(m map[string]interface{}, x *CallFlowUpdateRequest) {
	if val, ok := m["description"]; ok && val != nil {
		x.Description = val.(string)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["xmlTree"]; ok && val != nil {
		x.XmlTree = val.(string)
	}
}
