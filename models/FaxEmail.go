package models

type FaxEmail struct {
	From      string `bson:"from" json:"from"`
	Id        string `bson:"id" json:"id"`
	MessageID string `bson:"messageID" json:"messageID"`
	Subject   string `bson:"subject" json:"subject"`
	ThreadId  string `bson:"threadId" json:"threadId"`
}

// BUILDER from bson map:
func BuildFaxEmail(m map[string]interface{}, x *FaxEmail) {
	if val, ok := m["from"]; ok && val != nil {
		x.From = val.(string)
	}
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["messageID"]; ok && val != nil {
		x.MessageID = val.(string)
	}
	if val, ok := m["subject"]; ok && val != nil {
		x.Subject = val.(string)
	}
	if val, ok := m["threadId"]; ok && val != nil {
		x.ThreadId = val.(string)
	}
}
