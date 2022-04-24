package models

type Audio struct {
	Id                string     `bson:"_id" json:"id"`
	AudioMp3          StoredFile `bson:"audioMp3" json:"audioMp3"`
	AudioSln          StoredFile `bson:"audioSln" json:"audioSln"`
	AudioWav          StoredFile `bson:"audioWav" json:"audioWav"`
	DateCreated       string     `bson:"dateCreated" json:"dateCreated"`
	DateDeleted       string     `bson:"dateDeleted" json:"dateDeleted"`
	DateUpdated       string     `bson:"dateUpdated" json:"dateUpdated"`
	Description       string     `bson:"description" json:"description"`
	DurationInSeconds int32      `bson:"durationInSeconds" json:"durationInSeconds"`
	FriendlyName      string     `bson:"friendlyName" json:"friendlyName"`
	IdAccount         string     `bson:"idAccount" json:"idAccount"`
}

// Implementing interface IUbluxDocument
func (x Audio) GetDateCreated() string {
	return x.DateCreated
}
func (x Audio) GetDateDeleted() string {
	return x.DateDeleted
}
func (x Audio) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x Audio) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x Audio) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildAudio(m map[string]interface{}, x *Audio) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["audioMp3"]; ok && val != nil {
		BuildStoredFile(val.(map[string]interface{}), &x.AudioMp3)
	}
	if val, ok := m["audioSln"]; ok && val != nil {
		BuildStoredFile(val.(map[string]interface{}), &x.AudioSln)
	}
	if val, ok := m["audioWav"]; ok && val != nil {
		BuildStoredFile(val.(map[string]interface{}), &x.AudioWav)
	}
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(string)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(string)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(string)
	}
	if val, ok := m["description"]; ok && val != nil {
		x.Description = val.(string)
	}
	if val, ok := m["durationInSeconds"]; ok && val != nil {
		x.DurationInSeconds = val.(int32)
	}
	if val, ok := m["friendlyName"]; ok && val != nil {
		x.FriendlyName = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
}
