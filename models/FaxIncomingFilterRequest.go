package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FaxIncomingFilterRequest struct {
	DateCreated_EQ          primitive.DateTime `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE         primitive.DateTime `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE         primitive.DateTime `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ          primitive.DateTime `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE         primitive.DateTime `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE         primitive.DateTime `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	ErrorMessage_CON        string             `bson:"errorMessage_CON" json:"errorMessage_CON"`
	ErrorMessage_EQ         string             `bson:"errorMessage_EQ" json:"errorMessage_EQ"`
	ErrorMessage_REG        string             `bson:"errorMessage_REG" json:"errorMessage_REG"`
	FaxStatus_CON           string             `bson:"faxStatus_CON" json:"faxStatus_CON"`
	FaxStatus_EQ            string             `bson:"faxStatus_EQ" json:"faxStatus_EQ"`
	FaxStatus_REG           string             `bson:"faxStatus_REG" json:"faxStatus_REG"`
	From_CON                string             `bson:"from_CON" json:"from_CON"`
	From_EQ                 string             `bson:"from_EQ" json:"from_EQ"`
	From_REG                string             `bson:"from_REG" json:"from_REG"`
	Id_CON                  string             `bson:"id_CON" json:"id_CON"`
	Id_EQ                   string             `bson:"id_EQ" json:"id_EQ"`
	Id_REG                  string             `bson:"id_REG" json:"id_REG"`
	IdVoipNumberFax_CON     string             `bson:"idVoipNumberFax_CON" json:"idVoipNumberFax_CON"`
	IdVoipNumberFax_EQ      string             `bson:"idVoipNumberFax_EQ" json:"idVoipNumberFax_EQ"`
	IdVoipNumberFax_REG     string             `bson:"idVoipNumberFax_REG" json:"idVoipNumberFax_REG"`
	NumPages_EQ             int32              `bson:"numPages_EQ" json:"numPages_EQ"`
	NumPages_GTE            int32              `bson:"numPages_GTE" json:"numPages_GTE"`
	NumPages_LTE            int32              `bson:"numPages_LTE" json:"numPages_LTE"`
	Pdf_FileSizeInBytes_EQ  int32              `bson:"pdf_FileSizeInBytes_EQ" json:"pdf_FileSizeInBytes_EQ"`
	Pdf_FileSizeInBytes_GTE int32              `bson:"pdf_FileSizeInBytes_GTE" json:"pdf_FileSizeInBytes_GTE"`
	Pdf_FileSizeInBytes_LTE int32              `bson:"pdf_FileSizeInBytes_LTE" json:"pdf_FileSizeInBytes_LTE"`
	Pdf_Id_CON              string             `bson:"pdf_Id_CON" json:"pdf_Id_CON"`
	Pdf_Id_EQ               string             `bson:"pdf_Id_EQ" json:"pdf_Id_EQ"`
	Pdf_Id_REG              string             `bson:"pdf_Id_REG" json:"pdf_Id_REG"`
	Pdf_Md5Hash_CON         string             `bson:"pdf_Md5Hash_CON" json:"pdf_Md5Hash_CON"`
	Pdf_Md5Hash_EQ          string             `bson:"pdf_Md5Hash_EQ" json:"pdf_Md5Hash_EQ"`
	Pdf_Md5Hash_REG         string             `bson:"pdf_Md5Hash_REG" json:"pdf_Md5Hash_REG"`
	Pdf_Url_CON             string             `bson:"pdf_Url_CON" json:"pdf_Url_CON"`
	Pdf_Url_EQ              string             `bson:"pdf_Url_EQ" json:"pdf_Url_EQ"`
	Pdf_Url_REG             string             `bson:"pdf_Url_REG" json:"pdf_Url_REG"`
	To_CON                  string             `bson:"to_CON" json:"to_CON"`
	To_EQ                   string             `bson:"to_EQ" json:"to_EQ"`
	To_REG                  string             `bson:"to_REG" json:"to_REG"`
}

// BUILDER from bson map:
func BuildFaxIncomingFilterRequest(m map[string]interface{}, x *FaxIncomingFilterRequest) {
	if val, ok := m["dateCreated_EQ"]; ok && val != nil {
		x.DateCreated_EQ = val.(primitive.DateTime)
	}
	if val, ok := m["dateCreated_GTE"]; ok && val != nil {
		x.DateCreated_GTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateCreated_LTE"]; ok && val != nil {
		x.DateCreated_LTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated_EQ"]; ok && val != nil {
		x.DateUpdated_EQ = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated_GTE"]; ok && val != nil {
		x.DateUpdated_GTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated_LTE"]; ok && val != nil {
		x.DateUpdated_LTE = val.(primitive.DateTime)
	}
	if val, ok := m["errorMessage_CON"]; ok && val != nil {
		x.ErrorMessage_CON = val.(string)
	}
	if val, ok := m["errorMessage_EQ"]; ok && val != nil {
		x.ErrorMessage_EQ = val.(string)
	}
	if val, ok := m["errorMessage_REG"]; ok && val != nil {
		x.ErrorMessage_REG = val.(string)
	}
	if val, ok := m["faxStatus_CON"]; ok && val != nil {
		x.FaxStatus_CON = val.(string)
	}
	if val, ok := m["faxStatus_EQ"]; ok && val != nil {
		x.FaxStatus_EQ = val.(string)
	}
	if val, ok := m["faxStatus_REG"]; ok && val != nil {
		x.FaxStatus_REG = val.(string)
	}
	if val, ok := m["from_CON"]; ok && val != nil {
		x.From_CON = val.(string)
	}
	if val, ok := m["from_EQ"]; ok && val != nil {
		x.From_EQ = val.(string)
	}
	if val, ok := m["from_REG"]; ok && val != nil {
		x.From_REG = val.(string)
	}
	if val, ok := m["id_CON"]; ok && val != nil {
		x.Id_CON = val.(string)
	}
	if val, ok := m["id_EQ"]; ok && val != nil {
		x.Id_EQ = val.(string)
	}
	if val, ok := m["id_REG"]; ok && val != nil {
		x.Id_REG = val.(string)
	}
	if val, ok := m["idVoipNumberFax_CON"]; ok && val != nil {
		x.IdVoipNumberFax_CON = val.(string)
	}
	if val, ok := m["idVoipNumberFax_EQ"]; ok && val != nil {
		x.IdVoipNumberFax_EQ = val.(string)
	}
	if val, ok := m["idVoipNumberFax_REG"]; ok && val != nil {
		x.IdVoipNumberFax_REG = val.(string)
	}
	if val, ok := m["numPages_EQ"]; ok && val != nil {
		x.NumPages_EQ = val.(int32)
	}
	if val, ok := m["numPages_GTE"]; ok && val != nil {
		x.NumPages_GTE = val.(int32)
	}
	if val, ok := m["numPages_LTE"]; ok && val != nil {
		x.NumPages_LTE = val.(int32)
	}
	if val, ok := m["pdf_FileSizeInBytes_EQ"]; ok && val != nil {
		x.Pdf_FileSizeInBytes_EQ = val.(int32)
	}
	if val, ok := m["pdf_FileSizeInBytes_GTE"]; ok && val != nil {
		x.Pdf_FileSizeInBytes_GTE = val.(int32)
	}
	if val, ok := m["pdf_FileSizeInBytes_LTE"]; ok && val != nil {
		x.Pdf_FileSizeInBytes_LTE = val.(int32)
	}
	if val, ok := m["pdf_Id_CON"]; ok && val != nil {
		x.Pdf_Id_CON = val.(string)
	}
	if val, ok := m["pdf_Id_EQ"]; ok && val != nil {
		x.Pdf_Id_EQ = val.(string)
	}
	if val, ok := m["pdf_Id_REG"]; ok && val != nil {
		x.Pdf_Id_REG = val.(string)
	}
	if val, ok := m["pdf_Md5Hash_CON"]; ok && val != nil {
		x.Pdf_Md5Hash_CON = val.(string)
	}
	if val, ok := m["pdf_Md5Hash_EQ"]; ok && val != nil {
		x.Pdf_Md5Hash_EQ = val.(string)
	}
	if val, ok := m["pdf_Md5Hash_REG"]; ok && val != nil {
		x.Pdf_Md5Hash_REG = val.(string)
	}
	if val, ok := m["pdf_Url_CON"]; ok && val != nil {
		x.Pdf_Url_CON = val.(string)
	}
	if val, ok := m["pdf_Url_EQ"]; ok && val != nil {
		x.Pdf_Url_EQ = val.(string)
	}
	if val, ok := m["pdf_Url_REG"]; ok && val != nil {
		x.Pdf_Url_REG = val.(string)
	}
	if val, ok := m["to_CON"]; ok && val != nil {
		x.To_CON = val.(string)
	}
	if val, ok := m["to_EQ"]; ok && val != nil {
		x.To_EQ = val.(string)
	}
	if val, ok := m["to_REG"]; ok && val != nil {
		x.To_REG = val.(string)
	}
}
