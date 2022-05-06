package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FaxOutgoingGroupFilterRequest struct {
	ContainsError_EQ                      bool               `bson:"containsError_EQ" json:"containsError_EQ"`
	DateCreated_EQ                        primitive.DateTime `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                       primitive.DateTime `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                       primitive.DateTime `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                        primitive.DateTime `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                       primitive.DateTime `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                       primitive.DateTime `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	FaxEmail_From_CON                     string             `bson:"faxEmail_From_CON" json:"faxEmail_From_CON"`
	FaxEmail_From_EQ                      string             `bson:"faxEmail_From_EQ" json:"faxEmail_From_EQ"`
	FaxEmail_From_REG                     string             `bson:"faxEmail_From_REG" json:"faxEmail_From_REG"`
	FaxEmail_Id_CON                       string             `bson:"faxEmail_Id_CON" json:"faxEmail_Id_CON"`
	FaxEmail_Id_EQ                        string             `bson:"faxEmail_Id_EQ" json:"faxEmail_Id_EQ"`
	FaxEmail_Id_REG                       string             `bson:"faxEmail_Id_REG" json:"faxEmail_Id_REG"`
	FaxEmail_MessageID_CON                string             `bson:"faxEmail_MessageID_CON" json:"faxEmail_MessageID_CON"`
	FaxEmail_MessageID_EQ                 string             `bson:"faxEmail_MessageID_EQ" json:"faxEmail_MessageID_EQ"`
	FaxEmail_MessageID_REG                string             `bson:"faxEmail_MessageID_REG" json:"faxEmail_MessageID_REG"`
	FaxEmail_Subject_CON                  string             `bson:"faxEmail_Subject_CON" json:"faxEmail_Subject_CON"`
	FaxEmail_Subject_EQ                   string             `bson:"faxEmail_Subject_EQ" json:"faxEmail_Subject_EQ"`
	FaxEmail_Subject_REG                  string             `bson:"faxEmail_Subject_REG" json:"faxEmail_Subject_REG"`
	FaxEmail_ThreadId_CON                 string             `bson:"faxEmail_ThreadId_CON" json:"faxEmail_ThreadId_CON"`
	FaxEmail_ThreadId_EQ                  string             `bson:"faxEmail_ThreadId_EQ" json:"faxEmail_ThreadId_EQ"`
	FaxEmail_ThreadId_REG                 string             `bson:"faxEmail_ThreadId_REG" json:"faxEmail_ThreadId_REG"`
	FaxesOutgoing_ErrorMessage_CON        string             `bson:"faxesOutgoing_ErrorMessage_CON" json:"faxesOutgoing_ErrorMessage_CON"`
	FaxesOutgoing_ErrorMessage_EQ         string             `bson:"faxesOutgoing_ErrorMessage_EQ" json:"faxesOutgoing_ErrorMessage_EQ"`
	FaxesOutgoing_ErrorMessage_REG        string             `bson:"faxesOutgoing_ErrorMessage_REG" json:"faxesOutgoing_ErrorMessage_REG"`
	FaxesOutgoing_FaxStatus_CON           string             `bson:"faxesOutgoing_FaxStatus_CON" json:"faxesOutgoing_FaxStatus_CON"`
	FaxesOutgoing_FaxStatus_EQ            string             `bson:"faxesOutgoing_FaxStatus_EQ" json:"faxesOutgoing_FaxStatus_EQ"`
	FaxesOutgoing_FaxStatus_REG           string             `bson:"faxesOutgoing_FaxStatus_REG" json:"faxesOutgoing_FaxStatus_REG"`
	FaxesOutgoing_Id_CON                  string             `bson:"faxesOutgoing_Id_CON" json:"faxesOutgoing_Id_CON"`
	FaxesOutgoing_Id_EQ                   string             `bson:"faxesOutgoing_Id_EQ" json:"faxesOutgoing_Id_EQ"`
	FaxesOutgoing_Id_REG                  string             `bson:"faxesOutgoing_Id_REG" json:"faxesOutgoing_Id_REG"`
	FaxesOutgoing_IsPortrait_EQ           bool               `bson:"faxesOutgoing_IsPortrait_EQ" json:"faxesOutgoing_IsPortrait_EQ"`
	FaxesOutgoing_Name_CON                string             `bson:"faxesOutgoing_Name_CON" json:"faxesOutgoing_Name_CON"`
	FaxesOutgoing_Name_EQ                 string             `bson:"faxesOutgoing_Name_EQ" json:"faxesOutgoing_Name_EQ"`
	FaxesOutgoing_Name_REG                string             `bson:"faxesOutgoing_Name_REG" json:"faxesOutgoing_Name_REG"`
	FaxesOutgoing_NumberOfAttempts_EQ     int32              `bson:"faxesOutgoing_NumberOfAttempts_EQ" json:"faxesOutgoing_NumberOfAttempts_EQ"`
	FaxesOutgoing_NumberOfAttempts_GTE    int32              `bson:"faxesOutgoing_NumberOfAttempts_GTE" json:"faxesOutgoing_NumberOfAttempts_GTE"`
	FaxesOutgoing_NumberOfAttempts_LTE    int32              `bson:"faxesOutgoing_NumberOfAttempts_LTE" json:"faxesOutgoing_NumberOfAttempts_LTE"`
	FaxesOutgoing_NumberOfPages_EQ        int32              `bson:"faxesOutgoing_NumberOfPages_EQ" json:"faxesOutgoing_NumberOfPages_EQ"`
	FaxesOutgoing_NumberOfPages_GTE       int32              `bson:"faxesOutgoing_NumberOfPages_GTE" json:"faxesOutgoing_NumberOfPages_GTE"`
	FaxesOutgoing_NumberOfPages_LTE       int32              `bson:"faxesOutgoing_NumberOfPages_LTE" json:"faxesOutgoing_NumberOfPages_LTE"`
	FaxesOutgoing_NumberOfPagesSent_EQ    int32              `bson:"faxesOutgoing_NumberOfPagesSent_EQ" json:"faxesOutgoing_NumberOfPagesSent_EQ"`
	FaxesOutgoing_NumberOfPagesSent_GTE   int32              `bson:"faxesOutgoing_NumberOfPagesSent_GTE" json:"faxesOutgoing_NumberOfPagesSent_GTE"`
	FaxesOutgoing_NumberOfPagesSent_LTE   int32              `bson:"faxesOutgoing_NumberOfPagesSent_LTE" json:"faxesOutgoing_NumberOfPagesSent_LTE"`
	FaxesOutgoing_Pdf_FileSizeInBytes_EQ  int32              `bson:"faxesOutgoing_Pdf_FileSizeInBytes_EQ" json:"faxesOutgoing_Pdf_FileSizeInBytes_EQ"`
	FaxesOutgoing_Pdf_FileSizeInBytes_GTE int32              `bson:"faxesOutgoing_Pdf_FileSizeInBytes_GTE" json:"faxesOutgoing_Pdf_FileSizeInBytes_GTE"`
	FaxesOutgoing_Pdf_FileSizeInBytes_LTE int32              `bson:"faxesOutgoing_Pdf_FileSizeInBytes_LTE" json:"faxesOutgoing_Pdf_FileSizeInBytes_LTE"`
	FaxesOutgoing_Pdf_Id_CON              string             `bson:"faxesOutgoing_Pdf_Id_CON" json:"faxesOutgoing_Pdf_Id_CON"`
	FaxesOutgoing_Pdf_Id_EQ               string             `bson:"faxesOutgoing_Pdf_Id_EQ" json:"faxesOutgoing_Pdf_Id_EQ"`
	FaxesOutgoing_Pdf_Id_REG              string             `bson:"faxesOutgoing_Pdf_Id_REG" json:"faxesOutgoing_Pdf_Id_REG"`
	FaxesOutgoing_Pdf_Md5Hash_CON         string             `bson:"faxesOutgoing_Pdf_Md5Hash_CON" json:"faxesOutgoing_Pdf_Md5Hash_CON"`
	FaxesOutgoing_Pdf_Md5Hash_EQ          string             `bson:"faxesOutgoing_Pdf_Md5Hash_EQ" json:"faxesOutgoing_Pdf_Md5Hash_EQ"`
	FaxesOutgoing_Pdf_Md5Hash_REG         string             `bson:"faxesOutgoing_Pdf_Md5Hash_REG" json:"faxesOutgoing_Pdf_Md5Hash_REG"`
	FaxesOutgoing_Pdf_Url_CON             string             `bson:"faxesOutgoing_Pdf_Url_CON" json:"faxesOutgoing_Pdf_Url_CON"`
	FaxesOutgoing_Pdf_Url_EQ              string             `bson:"faxesOutgoing_Pdf_Url_EQ" json:"faxesOutgoing_Pdf_Url_EQ"`
	FaxesOutgoing_Pdf_Url_REG             string             `bson:"faxesOutgoing_Pdf_Url_REG" json:"faxesOutgoing_Pdf_Url_REG"`
	FaxesOutgoing_To_CON                  string             `bson:"faxesOutgoing_To_CON" json:"faxesOutgoing_To_CON"`
	FaxesOutgoing_To_EQ                   string             `bson:"faxesOutgoing_To_EQ" json:"faxesOutgoing_To_EQ"`
	FaxesOutgoing_To_REG                  string             `bson:"faxesOutgoing_To_REG" json:"faxesOutgoing_To_REG"`
	From_CON                              string             `bson:"from_CON" json:"from_CON"`
	From_EQ                               string             `bson:"from_EQ" json:"from_EQ"`
	From_REG                              string             `bson:"from_REG" json:"from_REG"`
	Id_CON                                string             `bson:"id_CON" json:"id_CON"`
	Id_EQ                                 string             `bson:"id_EQ" json:"id_EQ"`
	Id_REG                                string             `bson:"id_REG" json:"id_REG"`
	IdVoipNumberFax_CON                   string             `bson:"idVoipNumberFax_CON" json:"idVoipNumberFax_CON"`
	IdVoipNumberFax_EQ                    string             `bson:"idVoipNumberFax_EQ" json:"idVoipNumberFax_EQ"`
	IdVoipNumberFax_REG                   string             `bson:"idVoipNumberFax_REG" json:"idVoipNumberFax_REG"`
	SendConfirmationToEmails_CON          string             `bson:"sendConfirmationToEmails_CON" json:"sendConfirmationToEmails_CON"`
	SendConfirmationToEmails_EQ           string             `bson:"sendConfirmationToEmails_EQ" json:"sendConfirmationToEmails_EQ"`
	SendConfirmationToEmails_REG          string             `bson:"sendConfirmationToEmails_REG" json:"sendConfirmationToEmails_REG"`
}

// BUILDER from bson map:
func BuildFaxOutgoingGroupFilterRequest(m map[string]interface{}, x *FaxOutgoingGroupFilterRequest) {
	if val, ok := m["containsError_EQ"]; ok && val != nil {
		x.ContainsError_EQ = val.(bool)
	}
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
	if val, ok := m["faxEmail_From_CON"]; ok && val != nil {
		x.FaxEmail_From_CON = val.(string)
	}
	if val, ok := m["faxEmail_From_EQ"]; ok && val != nil {
		x.FaxEmail_From_EQ = val.(string)
	}
	if val, ok := m["faxEmail_From_REG"]; ok && val != nil {
		x.FaxEmail_From_REG = val.(string)
	}
	if val, ok := m["faxEmail_Id_CON"]; ok && val != nil {
		x.FaxEmail_Id_CON = val.(string)
	}
	if val, ok := m["faxEmail_Id_EQ"]; ok && val != nil {
		x.FaxEmail_Id_EQ = val.(string)
	}
	if val, ok := m["faxEmail_Id_REG"]; ok && val != nil {
		x.FaxEmail_Id_REG = val.(string)
	}
	if val, ok := m["faxEmail_MessageID_CON"]; ok && val != nil {
		x.FaxEmail_MessageID_CON = val.(string)
	}
	if val, ok := m["faxEmail_MessageID_EQ"]; ok && val != nil {
		x.FaxEmail_MessageID_EQ = val.(string)
	}
	if val, ok := m["faxEmail_MessageID_REG"]; ok && val != nil {
		x.FaxEmail_MessageID_REG = val.(string)
	}
	if val, ok := m["faxEmail_Subject_CON"]; ok && val != nil {
		x.FaxEmail_Subject_CON = val.(string)
	}
	if val, ok := m["faxEmail_Subject_EQ"]; ok && val != nil {
		x.FaxEmail_Subject_EQ = val.(string)
	}
	if val, ok := m["faxEmail_Subject_REG"]; ok && val != nil {
		x.FaxEmail_Subject_REG = val.(string)
	}
	if val, ok := m["faxEmail_ThreadId_CON"]; ok && val != nil {
		x.FaxEmail_ThreadId_CON = val.(string)
	}
	if val, ok := m["faxEmail_ThreadId_EQ"]; ok && val != nil {
		x.FaxEmail_ThreadId_EQ = val.(string)
	}
	if val, ok := m["faxEmail_ThreadId_REG"]; ok && val != nil {
		x.FaxEmail_ThreadId_REG = val.(string)
	}
	if val, ok := m["faxesOutgoing_ErrorMessage_CON"]; ok && val != nil {
		x.FaxesOutgoing_ErrorMessage_CON = val.(string)
	}
	if val, ok := m["faxesOutgoing_ErrorMessage_EQ"]; ok && val != nil {
		x.FaxesOutgoing_ErrorMessage_EQ = val.(string)
	}
	if val, ok := m["faxesOutgoing_ErrorMessage_REG"]; ok && val != nil {
		x.FaxesOutgoing_ErrorMessage_REG = val.(string)
	}
	if val, ok := m["faxesOutgoing_FaxStatus_CON"]; ok && val != nil {
		x.FaxesOutgoing_FaxStatus_CON = val.(string)
	}
	if val, ok := m["faxesOutgoing_FaxStatus_EQ"]; ok && val != nil {
		x.FaxesOutgoing_FaxStatus_EQ = val.(string)
	}
	if val, ok := m["faxesOutgoing_FaxStatus_REG"]; ok && val != nil {
		x.FaxesOutgoing_FaxStatus_REG = val.(string)
	}
	if val, ok := m["faxesOutgoing_Id_CON"]; ok && val != nil {
		x.FaxesOutgoing_Id_CON = val.(string)
	}
	if val, ok := m["faxesOutgoing_Id_EQ"]; ok && val != nil {
		x.FaxesOutgoing_Id_EQ = val.(string)
	}
	if val, ok := m["faxesOutgoing_Id_REG"]; ok && val != nil {
		x.FaxesOutgoing_Id_REG = val.(string)
	}
	if val, ok := m["faxesOutgoing_IsPortrait_EQ"]; ok && val != nil {
		x.FaxesOutgoing_IsPortrait_EQ = val.(bool)
	}
	if val, ok := m["faxesOutgoing_Name_CON"]; ok && val != nil {
		x.FaxesOutgoing_Name_CON = val.(string)
	}
	if val, ok := m["faxesOutgoing_Name_EQ"]; ok && val != nil {
		x.FaxesOutgoing_Name_EQ = val.(string)
	}
	if val, ok := m["faxesOutgoing_Name_REG"]; ok && val != nil {
		x.FaxesOutgoing_Name_REG = val.(string)
	}
	if val, ok := m["faxesOutgoing_NumberOfAttempts_EQ"]; ok && val != nil {
		x.FaxesOutgoing_NumberOfAttempts_EQ = val.(int32)
	}
	if val, ok := m["faxesOutgoing_NumberOfAttempts_GTE"]; ok && val != nil {
		x.FaxesOutgoing_NumberOfAttempts_GTE = val.(int32)
	}
	if val, ok := m["faxesOutgoing_NumberOfAttempts_LTE"]; ok && val != nil {
		x.FaxesOutgoing_NumberOfAttempts_LTE = val.(int32)
	}
	if val, ok := m["faxesOutgoing_NumberOfPages_EQ"]; ok && val != nil {
		x.FaxesOutgoing_NumberOfPages_EQ = val.(int32)
	}
	if val, ok := m["faxesOutgoing_NumberOfPages_GTE"]; ok && val != nil {
		x.FaxesOutgoing_NumberOfPages_GTE = val.(int32)
	}
	if val, ok := m["faxesOutgoing_NumberOfPages_LTE"]; ok && val != nil {
		x.FaxesOutgoing_NumberOfPages_LTE = val.(int32)
	}
	if val, ok := m["faxesOutgoing_NumberOfPagesSent_EQ"]; ok && val != nil {
		x.FaxesOutgoing_NumberOfPagesSent_EQ = val.(int32)
	}
	if val, ok := m["faxesOutgoing_NumberOfPagesSent_GTE"]; ok && val != nil {
		x.FaxesOutgoing_NumberOfPagesSent_GTE = val.(int32)
	}
	if val, ok := m["faxesOutgoing_NumberOfPagesSent_LTE"]; ok && val != nil {
		x.FaxesOutgoing_NumberOfPagesSent_LTE = val.(int32)
	}
	if val, ok := m["faxesOutgoing_Pdf_FileSizeInBytes_EQ"]; ok && val != nil {
		x.FaxesOutgoing_Pdf_FileSizeInBytes_EQ = val.(int32)
	}
	if val, ok := m["faxesOutgoing_Pdf_FileSizeInBytes_GTE"]; ok && val != nil {
		x.FaxesOutgoing_Pdf_FileSizeInBytes_GTE = val.(int32)
	}
	if val, ok := m["faxesOutgoing_Pdf_FileSizeInBytes_LTE"]; ok && val != nil {
		x.FaxesOutgoing_Pdf_FileSizeInBytes_LTE = val.(int32)
	}
	if val, ok := m["faxesOutgoing_Pdf_Id_CON"]; ok && val != nil {
		x.FaxesOutgoing_Pdf_Id_CON = val.(string)
	}
	if val, ok := m["faxesOutgoing_Pdf_Id_EQ"]; ok && val != nil {
		x.FaxesOutgoing_Pdf_Id_EQ = val.(string)
	}
	if val, ok := m["faxesOutgoing_Pdf_Id_REG"]; ok && val != nil {
		x.FaxesOutgoing_Pdf_Id_REG = val.(string)
	}
	if val, ok := m["faxesOutgoing_Pdf_Md5Hash_CON"]; ok && val != nil {
		x.FaxesOutgoing_Pdf_Md5Hash_CON = val.(string)
	}
	if val, ok := m["faxesOutgoing_Pdf_Md5Hash_EQ"]; ok && val != nil {
		x.FaxesOutgoing_Pdf_Md5Hash_EQ = val.(string)
	}
	if val, ok := m["faxesOutgoing_Pdf_Md5Hash_REG"]; ok && val != nil {
		x.FaxesOutgoing_Pdf_Md5Hash_REG = val.(string)
	}
	if val, ok := m["faxesOutgoing_Pdf_Url_CON"]; ok && val != nil {
		x.FaxesOutgoing_Pdf_Url_CON = val.(string)
	}
	if val, ok := m["faxesOutgoing_Pdf_Url_EQ"]; ok && val != nil {
		x.FaxesOutgoing_Pdf_Url_EQ = val.(string)
	}
	if val, ok := m["faxesOutgoing_Pdf_Url_REG"]; ok && val != nil {
		x.FaxesOutgoing_Pdf_Url_REG = val.(string)
	}
	if val, ok := m["faxesOutgoing_To_CON"]; ok && val != nil {
		x.FaxesOutgoing_To_CON = val.(string)
	}
	if val, ok := m["faxesOutgoing_To_EQ"]; ok && val != nil {
		x.FaxesOutgoing_To_EQ = val.(string)
	}
	if val, ok := m["faxesOutgoing_To_REG"]; ok && val != nil {
		x.FaxesOutgoing_To_REG = val.(string)
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
	if val, ok := m["sendConfirmationToEmails_CON"]; ok && val != nil {
		x.SendConfirmationToEmails_CON = val.(string)
	}
	if val, ok := m["sendConfirmationToEmails_EQ"]; ok && val != nil {
		x.SendConfirmationToEmails_EQ = val.(string)
	}
	if val, ok := m["sendConfirmationToEmails_REG"]; ok && val != nil {
		x.SendConfirmationToEmails_REG = val.(string)
	}
}
