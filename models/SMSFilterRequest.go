package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type SMSFilterRequest struct {
	Body_CON                                             string             `bson:"body_CON" json:"body_CON"`
	Body_EQ                                              string             `bson:"body_EQ" json:"body_EQ"`
	Body_REG                                             string             `bson:"body_REG" json:"body_REG"`
	Contact_Company_CON                                  string             `bson:"contact_Company_CON" json:"contact_Company_CON"`
	Contact_Company_EQ                                   string             `bson:"contact_Company_EQ" json:"contact_Company_EQ"`
	Contact_Company_REG                                  string             `bson:"contact_Company_REG" json:"contact_Company_REG"`
	Contact_ContactEmails_Email_CON                      string             `bson:"contact_ContactEmails_Email_CON" json:"contact_ContactEmails_Email_CON"`
	Contact_ContactEmails_Email_EQ                       string             `bson:"contact_ContactEmails_Email_EQ" json:"contact_ContactEmails_Email_EQ"`
	Contact_ContactEmails_Email_REG                      string             `bson:"contact_ContactEmails_Email_REG" json:"contact_ContactEmails_Email_REG"`
	Contact_ContactEmails_Label_EQ                       LabelEmailType     `bson:"contact_ContactEmails_Label_EQ" json:"contact_ContactEmails_Label_EQ"`
	Contact_ContactNumbers_Label_EQ                      LabelNumber        `bson:"contact_ContactNumbers_Label_EQ" json:"contact_ContactNumbers_Label_EQ"`
	Contact_ContactNumbers_Number_CON                    string             `bson:"contact_ContactNumbers_Number_CON" json:"contact_ContactNumbers_Number_CON"`
	Contact_ContactNumbers_Number_EQ                     string             `bson:"contact_ContactNumbers_Number_EQ" json:"contact_ContactNumbers_Number_EQ"`
	Contact_ContactNumbers_Number_REG                    string             `bson:"contact_ContactNumbers_Number_REG" json:"contact_ContactNumbers_Number_REG"`
	Contact_ContactNumbers_NumberInternationalFormat_CON string             `bson:"contact_ContactNumbers_NumberInternationalFormat_CON" json:"contact_ContactNumbers_NumberInternationalFormat_CON"`
	Contact_ContactNumbers_NumberInternationalFormat_EQ  string             `bson:"contact_ContactNumbers_NumberInternationalFormat_EQ" json:"contact_ContactNumbers_NumberInternationalFormat_EQ"`
	Contact_ContactNumbers_NumberInternationalFormat_REG string             `bson:"contact_ContactNumbers_NumberInternationalFormat_REG" json:"contact_ContactNumbers_NumberInternationalFormat_REG"`
	Contact_DateCreated_EQ                               primitive.DateTime `bson:"contact_DateCreated_EQ" json:"contact_DateCreated_EQ"`
	Contact_DateCreated_GTE                              primitive.DateTime `bson:"contact_DateCreated_GTE" json:"contact_DateCreated_GTE"`
	Contact_DateCreated_LTE                              primitive.DateTime `bson:"contact_DateCreated_LTE" json:"contact_DateCreated_LTE"`
	Contact_DateUpdated_EQ                               primitive.DateTime `bson:"contact_DateUpdated_EQ" json:"contact_DateUpdated_EQ"`
	Contact_DateUpdated_GTE                              primitive.DateTime `bson:"contact_DateUpdated_GTE" json:"contact_DateUpdated_GTE"`
	Contact_DateUpdated_LTE                              primitive.DateTime `bson:"contact_DateUpdated_LTE" json:"contact_DateUpdated_LTE"`
	Contact_FirstName_CON                                string             `bson:"contact_FirstName_CON" json:"contact_FirstName_CON"`
	Contact_FirstName_EQ                                 string             `bson:"contact_FirstName_EQ" json:"contact_FirstName_EQ"`
	Contact_FirstName_REG                                string             `bson:"contact_FirstName_REG" json:"contact_FirstName_REG"`
	Contact_Hash_CON                                     string             `bson:"contact_Hash_CON" json:"contact_Hash_CON"`
	Contact_Hash_EQ                                      string             `bson:"contact_Hash_EQ" json:"contact_Hash_EQ"`
	Contact_Hash_REG                                     string             `bson:"contact_Hash_REG" json:"contact_Hash_REG"`
	Contact_Id_CON                                       string             `bson:"contact_Id_CON" json:"contact_Id_CON"`
	Contact_Id_EQ                                        string             `bson:"contact_Id_EQ" json:"contact_Id_EQ"`
	Contact_Id_REG                                       string             `bson:"contact_Id_REG" json:"contact_Id_REG"`
	Contact_JobTittle_CON                                string             `bson:"contact_JobTittle_CON" json:"contact_JobTittle_CON"`
	Contact_JobTittle_EQ                                 string             `bson:"contact_JobTittle_EQ" json:"contact_JobTittle_EQ"`
	Contact_JobTittle_REG                                string             `bson:"contact_JobTittle_REG" json:"contact_JobTittle_REG"`
	Contact_LastName_CON                                 string             `bson:"contact_LastName_CON" json:"contact_LastName_CON"`
	Contact_LastName_EQ                                  string             `bson:"contact_LastName_EQ" json:"contact_LastName_EQ"`
	Contact_LastName_REG                                 string             `bson:"contact_LastName_REG" json:"contact_LastName_REG"`
	Contact_Notes_CON                                    string             `bson:"contact_Notes_CON" json:"contact_Notes_CON"`
	Contact_Notes_EQ                                     string             `bson:"contact_Notes_EQ" json:"contact_Notes_EQ"`
	Contact_Notes_REG                                    string             `bson:"contact_Notes_REG" json:"contact_Notes_REG"`
	Contact_Variables_JsonValue_CON                      string             `bson:"contact_Variables_JsonValue_CON" json:"contact_Variables_JsonValue_CON"`
	Contact_Variables_JsonValue_EQ                       string             `bson:"contact_Variables_JsonValue_EQ" json:"contact_Variables_JsonValue_EQ"`
	Contact_Variables_JsonValue_REG                      string             `bson:"contact_Variables_JsonValue_REG" json:"contact_Variables_JsonValue_REG"`
	Contact_Variables_Name_CON                           string             `bson:"contact_Variables_Name_CON" json:"contact_Variables_Name_CON"`
	Contact_Variables_Name_EQ                            string             `bson:"contact_Variables_Name_EQ" json:"contact_Variables_Name_EQ"`
	Contact_Variables_Name_REG                           string             `bson:"contact_Variables_Name_REG" json:"contact_Variables_Name_REG"`
	DateCreated_EQ                                       primitive.DateTime `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                                      primitive.DateTime `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                                      primitive.DateTime `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                                       primitive.DateTime `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                                      primitive.DateTime `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                                      primitive.DateTime `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	From_CON                                             string             `bson:"from_CON" json:"from_CON"`
	From_EQ                                              string             `bson:"from_EQ" json:"from_EQ"`
	From_REG                                             string             `bson:"from_REG" json:"from_REG"`
	Id_CON                                               string             `bson:"id_CON" json:"id_CON"`
	Id_EQ                                                string             `bson:"id_EQ" json:"id_EQ"`
	Id_REG                                               string             `bson:"id_REG" json:"id_REG"`
	IdVoipNumber_CON                                     string             `bson:"idVoipNumber_CON" json:"idVoipNumber_CON"`
	IdVoipNumber_EQ                                      string             `bson:"idVoipNumber_EQ" json:"idVoipNumber_EQ"`
	IdVoipNumber_REG                                     string             `bson:"idVoipNumber_REG" json:"idVoipNumber_REG"`
	IsIncoming_EQ                                        bool               `bson:"isIncoming_EQ" json:"isIncoming_EQ"`
	NumSegments_EQ                                       int32              `bson:"numSegments_EQ" json:"numSegments_EQ"`
	NumSegments_GTE                                      int32              `bson:"numSegments_GTE" json:"numSegments_GTE"`
	NumSegments_LTE                                      int32              `bson:"numSegments_LTE" json:"numSegments_LTE"`
	Status_CON                                           string             `bson:"status_CON" json:"status_CON"`
	Status_EQ                                            string             `bson:"status_EQ" json:"status_EQ"`
	Status_REG                                           string             `bson:"status_REG" json:"status_REG"`
	To_CON                                               string             `bson:"to_CON" json:"to_CON"`
	To_EQ                                                string             `bson:"to_EQ" json:"to_EQ"`
	To_REG                                               string             `bson:"to_REG" json:"to_REG"`
}

// BUILDER from bson map:
func BuildSMSFilterRequest(m map[string]interface{}, x *SMSFilterRequest) {
	if val, ok := m["body_CON"]; ok && val != nil {
		x.Body_CON = val.(string)
	}
	if val, ok := m["body_EQ"]; ok && val != nil {
		x.Body_EQ = val.(string)
	}
	if val, ok := m["body_REG"]; ok && val != nil {
		x.Body_REG = val.(string)
	}
	if val, ok := m["contact_Company_CON"]; ok && val != nil {
		x.Contact_Company_CON = val.(string)
	}
	if val, ok := m["contact_Company_EQ"]; ok && val != nil {
		x.Contact_Company_EQ = val.(string)
	}
	if val, ok := m["contact_Company_REG"]; ok && val != nil {
		x.Contact_Company_REG = val.(string)
	}
	if val, ok := m["contact_ContactEmails_Email_CON"]; ok && val != nil {
		x.Contact_ContactEmails_Email_CON = val.(string)
	}
	if val, ok := m["contact_ContactEmails_Email_EQ"]; ok && val != nil {
		x.Contact_ContactEmails_Email_EQ = val.(string)
	}
	if val, ok := m["contact_ContactEmails_Email_REG"]; ok && val != nil {
		x.Contact_ContactEmails_Email_REG = val.(string)
	}
	if val, ok := m["contact_ContactEmails_Label_EQ"]; ok && val != nil {
		x.Contact_ContactEmails_Label_EQ = LabelEmailType("Contact_ContactEmails_Label_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["contact_ContactNumbers_Label_EQ"]; ok && val != nil {
		x.Contact_ContactNumbers_Label_EQ = LabelNumber("Contact_ContactNumbers_Label_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["contact_ContactNumbers_Number_CON"]; ok && val != nil {
		x.Contact_ContactNumbers_Number_CON = val.(string)
	}
	if val, ok := m["contact_ContactNumbers_Number_EQ"]; ok && val != nil {
		x.Contact_ContactNumbers_Number_EQ = val.(string)
	}
	if val, ok := m["contact_ContactNumbers_Number_REG"]; ok && val != nil {
		x.Contact_ContactNumbers_Number_REG = val.(string)
	}
	if val, ok := m["contact_ContactNumbers_NumberInternationalFormat_CON"]; ok && val != nil {
		x.Contact_ContactNumbers_NumberInternationalFormat_CON = val.(string)
	}
	if val, ok := m["contact_ContactNumbers_NumberInternationalFormat_EQ"]; ok && val != nil {
		x.Contact_ContactNumbers_NumberInternationalFormat_EQ = val.(string)
	}
	if val, ok := m["contact_ContactNumbers_NumberInternationalFormat_REG"]; ok && val != nil {
		x.Contact_ContactNumbers_NumberInternationalFormat_REG = val.(string)
	}
	if val, ok := m["contact_DateCreated_EQ"]; ok && val != nil {
		x.Contact_DateCreated_EQ = val.(primitive.DateTime)
	}
	if val, ok := m["contact_DateCreated_GTE"]; ok && val != nil {
		x.Contact_DateCreated_GTE = val.(primitive.DateTime)
	}
	if val, ok := m["contact_DateCreated_LTE"]; ok && val != nil {
		x.Contact_DateCreated_LTE = val.(primitive.DateTime)
	}
	if val, ok := m["contact_DateUpdated_EQ"]; ok && val != nil {
		x.Contact_DateUpdated_EQ = val.(primitive.DateTime)
	}
	if val, ok := m["contact_DateUpdated_GTE"]; ok && val != nil {
		x.Contact_DateUpdated_GTE = val.(primitive.DateTime)
	}
	if val, ok := m["contact_DateUpdated_LTE"]; ok && val != nil {
		x.Contact_DateUpdated_LTE = val.(primitive.DateTime)
	}
	if val, ok := m["contact_FirstName_CON"]; ok && val != nil {
		x.Contact_FirstName_CON = val.(string)
	}
	if val, ok := m["contact_FirstName_EQ"]; ok && val != nil {
		x.Contact_FirstName_EQ = val.(string)
	}
	if val, ok := m["contact_FirstName_REG"]; ok && val != nil {
		x.Contact_FirstName_REG = val.(string)
	}
	if val, ok := m["contact_Hash_CON"]; ok && val != nil {
		x.Contact_Hash_CON = val.(string)
	}
	if val, ok := m["contact_Hash_EQ"]; ok && val != nil {
		x.Contact_Hash_EQ = val.(string)
	}
	if val, ok := m["contact_Hash_REG"]; ok && val != nil {
		x.Contact_Hash_REG = val.(string)
	}
	if val, ok := m["contact_Id_CON"]; ok && val != nil {
		x.Contact_Id_CON = val.(string)
	}
	if val, ok := m["contact_Id_EQ"]; ok && val != nil {
		x.Contact_Id_EQ = val.(string)
	}
	if val, ok := m["contact_Id_REG"]; ok && val != nil {
		x.Contact_Id_REG = val.(string)
	}
	if val, ok := m["contact_JobTittle_CON"]; ok && val != nil {
		x.Contact_JobTittle_CON = val.(string)
	}
	if val, ok := m["contact_JobTittle_EQ"]; ok && val != nil {
		x.Contact_JobTittle_EQ = val.(string)
	}
	if val, ok := m["contact_JobTittle_REG"]; ok && val != nil {
		x.Contact_JobTittle_REG = val.(string)
	}
	if val, ok := m["contact_LastName_CON"]; ok && val != nil {
		x.Contact_LastName_CON = val.(string)
	}
	if val, ok := m["contact_LastName_EQ"]; ok && val != nil {
		x.Contact_LastName_EQ = val.(string)
	}
	if val, ok := m["contact_LastName_REG"]; ok && val != nil {
		x.Contact_LastName_REG = val.(string)
	}
	if val, ok := m["contact_Notes_CON"]; ok && val != nil {
		x.Contact_Notes_CON = val.(string)
	}
	if val, ok := m["contact_Notes_EQ"]; ok && val != nil {
		x.Contact_Notes_EQ = val.(string)
	}
	if val, ok := m["contact_Notes_REG"]; ok && val != nil {
		x.Contact_Notes_REG = val.(string)
	}
	if val, ok := m["contact_Variables_JsonValue_CON"]; ok && val != nil {
		x.Contact_Variables_JsonValue_CON = val.(string)
	}
	if val, ok := m["contact_Variables_JsonValue_EQ"]; ok && val != nil {
		x.Contact_Variables_JsonValue_EQ = val.(string)
	}
	if val, ok := m["contact_Variables_JsonValue_REG"]; ok && val != nil {
		x.Contact_Variables_JsonValue_REG = val.(string)
	}
	if val, ok := m["contact_Variables_Name_CON"]; ok && val != nil {
		x.Contact_Variables_Name_CON = val.(string)
	}
	if val, ok := m["contact_Variables_Name_EQ"]; ok && val != nil {
		x.Contact_Variables_Name_EQ = val.(string)
	}
	if val, ok := m["contact_Variables_Name_REG"]; ok && val != nil {
		x.Contact_Variables_Name_REG = val.(string)
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
	if val, ok := m["idVoipNumber_CON"]; ok && val != nil {
		x.IdVoipNumber_CON = val.(string)
	}
	if val, ok := m["idVoipNumber_EQ"]; ok && val != nil {
		x.IdVoipNumber_EQ = val.(string)
	}
	if val, ok := m["idVoipNumber_REG"]; ok && val != nil {
		x.IdVoipNumber_REG = val.(string)
	}
	if val, ok := m["isIncoming_EQ"]; ok && val != nil {
		x.IsIncoming_EQ = val.(bool)
	}
	if val, ok := m["numSegments_EQ"]; ok && val != nil {
		x.NumSegments_EQ = val.(int32)
	}
	if val, ok := m["numSegments_GTE"]; ok && val != nil {
		x.NumSegments_GTE = val.(int32)
	}
	if val, ok := m["numSegments_LTE"]; ok && val != nil {
		x.NumSegments_LTE = val.(int32)
	}
	if val, ok := m["status_CON"]; ok && val != nil {
		x.Status_CON = val.(string)
	}
	if val, ok := m["status_EQ"]; ok && val != nil {
		x.Status_EQ = val.(string)
	}
	if val, ok := m["status_REG"]; ok && val != nil {
		x.Status_REG = val.(string)
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
