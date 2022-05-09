package models

type ExtensionCallFlowUpdateRequest struct {
	CallFlowLabel                 string `bson:"callFlowLabel" json:"callFlowLabel"`
	Id                            string `bson:"id" json:"id"`
	IdCallFlow                    string `bson:"idCallFlow" json:"idCallFlow"`
	IdMusicOnHoldGroup            string `bson:"idMusicOnHoldGroup" json:"idMusicOnHoldGroup"`
	InjectExtensionNameToCallerId bool   `bson:"injectExtensionNameToCallerId" json:"injectExtensionNameToCallerId"`
	Number                        string `bson:"number" json:"number"`
	TimeZone                      string `bson:"timeZone" json:"timeZone"`
}

// Implementing interface IUbluxDocumentId
func (x ExtensionCallFlowUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildExtensionCallFlowUpdateRequest(m map[string]interface{}, x *ExtensionCallFlowUpdateRequest) {
	if val, ok := m["callFlowLabel"]; ok && val != nil {
		x.CallFlowLabel = val.(string)
	}
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idCallFlow"]; ok && val != nil {
		x.IdCallFlow = val.(string)
	}
	if val, ok := m["idMusicOnHoldGroup"]; ok && val != nil {
		x.IdMusicOnHoldGroup = val.(string)
	}
	if val, ok := m["injectExtensionNameToCallerId"]; ok && val != nil {
		x.InjectExtensionNameToCallerId = val.(bool)
	}
	if val, ok := m["number"]; ok && val != nil {
		x.Number = val.(string)
	}
	if val, ok := m["timeZone"]; ok && val != nil {
		x.TimeZone = val.(string)
	}
}
