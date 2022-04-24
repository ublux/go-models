package models

type Line struct {
	Id                   string               `bson:"_id" json:"id"`
	CallerIdNumber       string               `bson:"callerIdNumber" json:"callerIdNumber"`
	FriendlyName         string               `bson:"friendlyName" json:"friendlyName"`
	LineConnectionStatus LineConnectionStatus `bson:"lineConnectionStatus" json:"lineConnectionStatus"`
	RecordExternalCalls  bool                 `bson:"recordExternalCalls" json:"recordExternalCalls"`
	RecordInternalCalls  bool                 `bson:"recordInternalCalls" json:"recordInternalCalls"`
}

// Implementing interface IUbluxSubDocument
// Implementing interface IUbluxDocumentId
func (x Line) GetId() string {
	return x.Id
}

// Implementing interface UbluxSubDocument

// BUILDER from bson map:
func BuildLine(m map[string]interface{}, x *Line) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["callerIdNumber"]; ok && val != nil {
		x.CallerIdNumber = val.(string)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["lineConnectionStatus"]; ok && val != nil {
		BuildLineConnectionStatus(val.(map[string]interface{}), &x.LineConnectionStatus)
	}
	if val, ok := m["recordExternalCalls"]; ok && val != nil {
		x.RecordExternalCalls = val.(bool)
	}
	if val, ok := m["recordInternalCalls"]; ok && val != nil {
		x.RecordInternalCalls = val.(bool)
	}
}
