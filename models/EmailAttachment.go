package models

type EmailAttachment struct {
	FilePath       string `bson:"filePath" json:"filePath"`
	Id             string `bson:"id" json:"id"`
	IdFaxEmailInfo string `bson:"idFaxEmailInfo" json:"idFaxEmailInfo"`
	MimeType       string `bson:"mimeType" json:"mimeType"`
	Name           string `bson:"name" json:"name"`
}

// BUILDER from bson map:
func BuildEmailAttachment(m map[string]interface{}, x *EmailAttachment) {
	if val, ok := m["filePath"]; ok && val != nil {
		x.FilePath = val.(string)
	}
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idFaxEmailInfo"]; ok && val != nil {
		x.IdFaxEmailInfo = val.(string)
	}
	if val, ok := m["mimeType"]; ok && val != nil {
		x.MimeType = val.(string)
	}
	if val, ok := m["name"]; ok && val != nil {
		x.Name = val.(string)
	}
}
