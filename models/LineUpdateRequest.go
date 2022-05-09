package models

type LineUpdateRequest struct {
	CallerIdNumber      string `bson:"callerIdNumber" json:"callerIdNumber"`
	Id                  string `bson:"id" json:"id"`
	RecordExternalCalls bool   `bson:"recordExternalCalls" json:"recordExternalCalls"`
	RecordInternalCalls bool   `bson:"recordInternalCalls" json:"recordInternalCalls"`
}

// Implementing interface IUbluxDocumentId
func (x LineUpdateRequest) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildLineUpdateRequest(m map[string]interface{}, x *LineUpdateRequest) {
	if val, ok := m["callerIdNumber"]; ok && val != nil {
		x.CallerIdNumber = val.(string)
	}
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["recordExternalCalls"]; ok && val != nil {
		x.RecordExternalCalls = val.(bool)
	}
	if val, ok := m["recordInternalCalls"]; ok && val != nil {
		x.RecordInternalCalls = val.(bool)
	}
}
