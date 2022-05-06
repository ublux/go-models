package models

import . "github.com/ublux/go-models/enums"
import "time"

type ContactFilterRequest struct {
	Company_CON                                  string         `bson:"company_CON" json:"company_CON"`
	Company_EQ                                   string         `bson:"company_EQ" json:"company_EQ"`
	Company_REG                                  string         `bson:"company_REG" json:"company_REG"`
	ContactEmails_Email_CON                      string         `bson:"contactEmails_Email_CON" json:"contactEmails_Email_CON"`
	ContactEmails_Email_EQ                       string         `bson:"contactEmails_Email_EQ" json:"contactEmails_Email_EQ"`
	ContactEmails_Email_REG                      string         `bson:"contactEmails_Email_REG" json:"contactEmails_Email_REG"`
	ContactEmails_Label_EQ                       LabelEmailType `bson:"contactEmails_Label_EQ" json:"contactEmails_Label_EQ"`
	ContactNumbers_Label_EQ                      LabelNumber    `bson:"contactNumbers_Label_EQ" json:"contactNumbers_Label_EQ"`
	ContactNumbers_Number_CON                    string         `bson:"contactNumbers_Number_CON" json:"contactNumbers_Number_CON"`
	ContactNumbers_Number_EQ                     string         `bson:"contactNumbers_Number_EQ" json:"contactNumbers_Number_EQ"`
	ContactNumbers_Number_REG                    string         `bson:"contactNumbers_Number_REG" json:"contactNumbers_Number_REG"`
	ContactNumbers_NumberInternationalFormat_CON string         `bson:"contactNumbers_NumberInternationalFormat_CON" json:"contactNumbers_NumberInternationalFormat_CON"`
	ContactNumbers_NumberInternationalFormat_EQ  string         `bson:"contactNumbers_NumberInternationalFormat_EQ" json:"contactNumbers_NumberInternationalFormat_EQ"`
	ContactNumbers_NumberInternationalFormat_REG string         `bson:"contactNumbers_NumberInternationalFormat_REG" json:"contactNumbers_NumberInternationalFormat_REG"`
	DateCreated_EQ                               time.Time      `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                              time.Time      `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                              time.Time      `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                               time.Time      `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                              time.Time      `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                              time.Time      `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	FirstName_CON                                string         `bson:"firstName_CON" json:"firstName_CON"`
	FirstName_EQ                                 string         `bson:"firstName_EQ" json:"firstName_EQ"`
	FirstName_REG                                string         `bson:"firstName_REG" json:"firstName_REG"`
	Hash_CON                                     string         `bson:"hash_CON" json:"hash_CON"`
	Hash_EQ                                      string         `bson:"hash_EQ" json:"hash_EQ"`
	Hash_REG                                     string         `bson:"hash_REG" json:"hash_REG"`
	Id_CON                                       string         `bson:"id_CON" json:"id_CON"`
	Id_EQ                                        string         `bson:"id_EQ" json:"id_EQ"`
	Id_REG                                       string         `bson:"id_REG" json:"id_REG"`
	JobTittle_CON                                string         `bson:"jobTittle_CON" json:"jobTittle_CON"`
	JobTittle_EQ                                 string         `bson:"jobTittle_EQ" json:"jobTittle_EQ"`
	JobTittle_REG                                string         `bson:"jobTittle_REG" json:"jobTittle_REG"`
	LastName_CON                                 string         `bson:"lastName_CON" json:"lastName_CON"`
	LastName_EQ                                  string         `bson:"lastName_EQ" json:"lastName_EQ"`
	LastName_REG                                 string         `bson:"lastName_REG" json:"lastName_REG"`
	Notes_CON                                    string         `bson:"notes_CON" json:"notes_CON"`
	Notes_EQ                                     string         `bson:"notes_EQ" json:"notes_EQ"`
	Notes_REG                                    string         `bson:"notes_REG" json:"notes_REG"`
	Variables_JsonValue_CON                      string         `bson:"variables_JsonValue_CON" json:"variables_JsonValue_CON"`
	Variables_JsonValue_EQ                       string         `bson:"variables_JsonValue_EQ" json:"variables_JsonValue_EQ"`
	Variables_JsonValue_REG                      string         `bson:"variables_JsonValue_REG" json:"variables_JsonValue_REG"`
	Variables_Name_CON                           string         `bson:"variables_Name_CON" json:"variables_Name_CON"`
	Variables_Name_EQ                            string         `bson:"variables_Name_EQ" json:"variables_Name_EQ"`
	Variables_Name_REG                           string         `bson:"variables_Name_REG" json:"variables_Name_REG"`
}

// BUILDER from bson map:
func BuildContactFilterRequest(m map[string]interface{}, x *ContactFilterRequest) {
	if val, ok := m["company_CON"]; ok && val != nil {
		x.Company_CON = val.(string)
	}
	if val, ok := m["company_EQ"]; ok && val != nil {
		x.Company_EQ = val.(string)
	}
	if val, ok := m["company_REG"]; ok && val != nil {
		x.Company_REG = val.(string)
	}
	if val, ok := m["contactEmails_Email_CON"]; ok && val != nil {
		x.ContactEmails_Email_CON = val.(string)
	}
	if val, ok := m["contactEmails_Email_EQ"]; ok && val != nil {
		x.ContactEmails_Email_EQ = val.(string)
	}
	if val, ok := m["contactEmails_Email_REG"]; ok && val != nil {
		x.ContactEmails_Email_REG = val.(string)
	}
	if val, ok := m["contactEmails_Label_EQ"]; ok && val != nil {
		x.ContactEmails_Label_EQ = LabelEmailType("ContactEmails_Label_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["contactNumbers_Label_EQ"]; ok && val != nil {
		x.ContactNumbers_Label_EQ = LabelNumber("ContactNumbers_Label_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["contactNumbers_Number_CON"]; ok && val != nil {
		x.ContactNumbers_Number_CON = val.(string)
	}
	if val, ok := m["contactNumbers_Number_EQ"]; ok && val != nil {
		x.ContactNumbers_Number_EQ = val.(string)
	}
	if val, ok := m["contactNumbers_Number_REG"]; ok && val != nil {
		x.ContactNumbers_Number_REG = val.(string)
	}
	if val, ok := m["contactNumbers_NumberInternationalFormat_CON"]; ok && val != nil {
		x.ContactNumbers_NumberInternationalFormat_CON = val.(string)
	}
	if val, ok := m["contactNumbers_NumberInternationalFormat_EQ"]; ok && val != nil {
		x.ContactNumbers_NumberInternationalFormat_EQ = val.(string)
	}
	if val, ok := m["contactNumbers_NumberInternationalFormat_REG"]; ok && val != nil {
		x.ContactNumbers_NumberInternationalFormat_REG = val.(string)
	}
	if val, ok := m["dateCreated_EQ"]; ok && val != nil {
		x.DateCreated_EQ = val.(time.Time)
	}
	if val, ok := m["dateCreated_GTE"]; ok && val != nil {
		x.DateCreated_GTE = val.(time.Time)
	}
	if val, ok := m["dateCreated_LTE"]; ok && val != nil {
		x.DateCreated_LTE = val.(time.Time)
	}
	if val, ok := m["dateUpdated_EQ"]; ok && val != nil {
		x.DateUpdated_EQ = val.(time.Time)
	}
	if val, ok := m["dateUpdated_GTE"]; ok && val != nil {
		x.DateUpdated_GTE = val.(time.Time)
	}
	if val, ok := m["dateUpdated_LTE"]; ok && val != nil {
		x.DateUpdated_LTE = val.(time.Time)
	}
	if val, ok := m["firstName_CON"]; ok && val != nil {
		x.FirstName_CON = val.(string)
	}
	if val, ok := m["firstName_EQ"]; ok && val != nil {
		x.FirstName_EQ = val.(string)
	}
	if val, ok := m["firstName_REG"]; ok && val != nil {
		x.FirstName_REG = val.(string)
	}
	if val, ok := m["hash_CON"]; ok && val != nil {
		x.Hash_CON = val.(string)
	}
	if val, ok := m["hash_EQ"]; ok && val != nil {
		x.Hash_EQ = val.(string)
	}
	if val, ok := m["hash_REG"]; ok && val != nil {
		x.Hash_REG = val.(string)
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
	if val, ok := m["jobTittle_CON"]; ok && val != nil {
		x.JobTittle_CON = val.(string)
	}
	if val, ok := m["jobTittle_EQ"]; ok && val != nil {
		x.JobTittle_EQ = val.(string)
	}
	if val, ok := m["jobTittle_REG"]; ok && val != nil {
		x.JobTittle_REG = val.(string)
	}
	if val, ok := m["lastName_CON"]; ok && val != nil {
		x.LastName_CON = val.(string)
	}
	if val, ok := m["lastName_EQ"]; ok && val != nil {
		x.LastName_EQ = val.(string)
	}
	if val, ok := m["lastName_REG"]; ok && val != nil {
		x.LastName_REG = val.(string)
	}
	if val, ok := m["notes_CON"]; ok && val != nil {
		x.Notes_CON = val.(string)
	}
	if val, ok := m["notes_EQ"]; ok && val != nil {
		x.Notes_EQ = val.(string)
	}
	if val, ok := m["notes_REG"]; ok && val != nil {
		x.Notes_REG = val.(string)
	}
	if val, ok := m["variables_JsonValue_CON"]; ok && val != nil {
		x.Variables_JsonValue_CON = val.(string)
	}
	if val, ok := m["variables_JsonValue_EQ"]; ok && val != nil {
		x.Variables_JsonValue_EQ = val.(string)
	}
	if val, ok := m["variables_JsonValue_REG"]; ok && val != nil {
		x.Variables_JsonValue_REG = val.(string)
	}
	if val, ok := m["variables_Name_CON"]; ok && val != nil {
		x.Variables_Name_CON = val.(string)
	}
	if val, ok := m["variables_Name_EQ"]; ok && val != nil {
		x.Variables_Name_EQ = val.(string)
	}
	if val, ok := m["variables_Name_REG"]; ok && val != nil {
		x.Variables_Name_REG = val.(string)
	}
}
