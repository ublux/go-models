package models

type Recording struct {
	ErrorMessage               string     `bson:"errorMessage" json:"errorMessage"`
	Id                         string     `bson:"id" json:"id"`
	RecordingDurationInSeconds int32      `bson:"recordingDurationInSeconds" json:"recordingDurationInSeconds"`
	RecordingMp3               StoredFile `bson:"recordingMp3" json:"recordingMp3"`
}

// Implementing interface UbluxSubDocument
func (x Recording) GetId() string {
	return x.Id
}

// BUILDER from bson map:
func BuildRecording(m map[string]interface{}, x *Recording) {
	if val, ok := m["errorMessage"]; ok && val != nil {
		x.ErrorMessage = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["recordingDurationInSeconds"]; ok && val != nil {
		x.RecordingDurationInSeconds = val.(int32)
	}
	if val, ok := m["recordingMp3"]; ok && val != nil {
		BuildStoredFile(val.(map[string]interface{}), &x.RecordingMp3)
	}
}
