package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type CallOutgoingToPSTNFilterRequest struct {
	CallType_EQ                                                     CallType           `bson:"callType_EQ" json:"callType_EQ"`
	ChannelVariables_CallerIdName_CON                               string             `bson:"channelVariables_CallerIdName_CON" json:"channelVariables_CallerIdName_CON"`
	ChannelVariables_CallerIdName_EQ                                string             `bson:"channelVariables_CallerIdName_EQ" json:"channelVariables_CallerIdName_EQ"`
	ChannelVariables_CallerIdName_REG                               string             `bson:"channelVariables_CallerIdName_REG" json:"channelVariables_CallerIdName_REG"`
	ChannelVariables_CallerIdNumber_CON                             string             `bson:"channelVariables_CallerIdNumber_CON" json:"channelVariables_CallerIdNumber_CON"`
	ChannelVariables_CallerIdNumber_EQ                              string             `bson:"channelVariables_CallerIdNumber_EQ" json:"channelVariables_CallerIdNumber_EQ"`
	ChannelVariables_CallerIdNumber_REG                             string             `bson:"channelVariables_CallerIdNumber_REG" json:"channelVariables_CallerIdNumber_REG"`
	ChannelVariables_IdMusicOnHold_CON                              string             `bson:"channelVariables_IdMusicOnHold_CON" json:"channelVariables_IdMusicOnHold_CON"`
	ChannelVariables_IdMusicOnHold_EQ                               string             `bson:"channelVariables_IdMusicOnHold_EQ" json:"channelVariables_IdMusicOnHold_EQ"`
	ChannelVariables_IdMusicOnHold_REG                              string             `bson:"channelVariables_IdMusicOnHold_REG" json:"channelVariables_IdMusicOnHold_REG"`
	ChannelVariables_Language_EQ                                    Language           `bson:"channelVariables_Language_EQ" json:"channelVariables_Language_EQ"`
	ChildCalls_ChildCallType_EQ                                     ChildCallType      `bson:"childCalls_ChildCallType_EQ" json:"childCalls_ChildCallType_EQ"`
	Contact_Company_CON                                             string             `bson:"contact_Company_CON" json:"contact_Company_CON"`
	Contact_Company_EQ                                              string             `bson:"contact_Company_EQ" json:"contact_Company_EQ"`
	Contact_Company_REG                                             string             `bson:"contact_Company_REG" json:"contact_Company_REG"`
	Contact_ContactEmails_Email_CON                                 string             `bson:"contact_ContactEmails_Email_CON" json:"contact_ContactEmails_Email_CON"`
	Contact_ContactEmails_Email_EQ                                  string             `bson:"contact_ContactEmails_Email_EQ" json:"contact_ContactEmails_Email_EQ"`
	Contact_ContactEmails_Email_REG                                 string             `bson:"contact_ContactEmails_Email_REG" json:"contact_ContactEmails_Email_REG"`
	Contact_ContactEmails_Label_EQ                                  LabelEmailType     `bson:"contact_ContactEmails_Label_EQ" json:"contact_ContactEmails_Label_EQ"`
	Contact_ContactEmails_SearchIndex_CON                           string             `bson:"contact_ContactEmails_SearchIndex_CON" json:"contact_ContactEmails_SearchIndex_CON"`
	Contact_ContactEmails_SearchIndex_EQ                            string             `bson:"contact_ContactEmails_SearchIndex_EQ" json:"contact_ContactEmails_SearchIndex_EQ"`
	Contact_ContactEmails_SearchIndex_REG                           string             `bson:"contact_ContactEmails_SearchIndex_REG" json:"contact_ContactEmails_SearchIndex_REG"`
	Contact_ContactNumbers_Label_EQ                                 LabelNumber        `bson:"contact_ContactNumbers_Label_EQ" json:"contact_ContactNumbers_Label_EQ"`
	Contact_ContactNumbers_Number_CON                               string             `bson:"contact_ContactNumbers_Number_CON" json:"contact_ContactNumbers_Number_CON"`
	Contact_ContactNumbers_Number_EQ                                string             `bson:"contact_ContactNumbers_Number_EQ" json:"contact_ContactNumbers_Number_EQ"`
	Contact_ContactNumbers_Number_REG                               string             `bson:"contact_ContactNumbers_Number_REG" json:"contact_ContactNumbers_Number_REG"`
	Contact_ContactNumbers_NumberInternationalFormat_CON            string             `bson:"contact_ContactNumbers_NumberInternationalFormat_CON" json:"contact_ContactNumbers_NumberInternationalFormat_CON"`
	Contact_ContactNumbers_NumberInternationalFormat_EQ             string             `bson:"contact_ContactNumbers_NumberInternationalFormat_EQ" json:"contact_ContactNumbers_NumberInternationalFormat_EQ"`
	Contact_ContactNumbers_NumberInternationalFormat_REG            string             `bson:"contact_ContactNumbers_NumberInternationalFormat_REG" json:"contact_ContactNumbers_NumberInternationalFormat_REG"`
	Contact_ContactNumbers_SearchIndex_CON                          string             `bson:"contact_ContactNumbers_SearchIndex_CON" json:"contact_ContactNumbers_SearchIndex_CON"`
	Contact_ContactNumbers_SearchIndex_EQ                           string             `bson:"contact_ContactNumbers_SearchIndex_EQ" json:"contact_ContactNumbers_SearchIndex_EQ"`
	Contact_ContactNumbers_SearchIndex_REG                          string             `bson:"contact_ContactNumbers_SearchIndex_REG" json:"contact_ContactNumbers_SearchIndex_REG"`
	Contact_DateCreated_EQ                                          primitive.DateTime `bson:"contact_DateCreated_EQ" json:"contact_DateCreated_EQ"`
	Contact_DateCreated_GTE                                         primitive.DateTime `bson:"contact_DateCreated_GTE" json:"contact_DateCreated_GTE"`
	Contact_DateCreated_LTE                                         primitive.DateTime `bson:"contact_DateCreated_LTE" json:"contact_DateCreated_LTE"`
	Contact_DateUpdated_EQ                                          primitive.DateTime `bson:"contact_DateUpdated_EQ" json:"contact_DateUpdated_EQ"`
	Contact_DateUpdated_GTE                                         primitive.DateTime `bson:"contact_DateUpdated_GTE" json:"contact_DateUpdated_GTE"`
	Contact_DateUpdated_LTE                                         primitive.DateTime `bson:"contact_DateUpdated_LTE" json:"contact_DateUpdated_LTE"`
	Contact_FirstName_CON                                           string             `bson:"contact_FirstName_CON" json:"contact_FirstName_CON"`
	Contact_FirstName_EQ                                            string             `bson:"contact_FirstName_EQ" json:"contact_FirstName_EQ"`
	Contact_FirstName_REG                                           string             `bson:"contact_FirstName_REG" json:"contact_FirstName_REG"`
	Contact_Hash_CON                                                string             `bson:"contact_Hash_CON" json:"contact_Hash_CON"`
	Contact_Hash_EQ                                                 string             `bson:"contact_Hash_EQ" json:"contact_Hash_EQ"`
	Contact_Hash_REG                                                string             `bson:"contact_Hash_REG" json:"contact_Hash_REG"`
	Contact_Id_CON                                                  string             `bson:"contact_Id_CON" json:"contact_Id_CON"`
	Contact_Id_EQ                                                   string             `bson:"contact_Id_EQ" json:"contact_Id_EQ"`
	Contact_Id_REG                                                  string             `bson:"contact_Id_REG" json:"contact_Id_REG"`
	Contact_IdIdentityUser_CON                                      string             `bson:"contact_IdIdentityUser_CON" json:"contact_IdIdentityUser_CON"`
	Contact_IdIdentityUser_EQ                                       string             `bson:"contact_IdIdentityUser_EQ" json:"contact_IdIdentityUser_EQ"`
	Contact_IdIdentityUser_REG                                      string             `bson:"contact_IdIdentityUser_REG" json:"contact_IdIdentityUser_REG"`
	Contact_JobTittle_CON                                           string             `bson:"contact_JobTittle_CON" json:"contact_JobTittle_CON"`
	Contact_JobTittle_EQ                                            string             `bson:"contact_JobTittle_EQ" json:"contact_JobTittle_EQ"`
	Contact_JobTittle_REG                                           string             `bson:"contact_JobTittle_REG" json:"contact_JobTittle_REG"`
	Contact_LastName_CON                                            string             `bson:"contact_LastName_CON" json:"contact_LastName_CON"`
	Contact_LastName_EQ                                             string             `bson:"contact_LastName_EQ" json:"contact_LastName_EQ"`
	Contact_LastName_REG                                            string             `bson:"contact_LastName_REG" json:"contact_LastName_REG"`
	Contact_Notes_CON                                               string             `bson:"contact_Notes_CON" json:"contact_Notes_CON"`
	Contact_Notes_EQ                                                string             `bson:"contact_Notes_EQ" json:"contact_Notes_EQ"`
	Contact_Notes_REG                                               string             `bson:"contact_Notes_REG" json:"contact_Notes_REG"`
	Contact_Variables_JsonValue_CON                                 string             `bson:"contact_Variables_JsonValue_CON" json:"contact_Variables_JsonValue_CON"`
	Contact_Variables_JsonValue_EQ                                  string             `bson:"contact_Variables_JsonValue_EQ" json:"contact_Variables_JsonValue_EQ"`
	Contact_Variables_JsonValue_REG                                 string             `bson:"contact_Variables_JsonValue_REG" json:"contact_Variables_JsonValue_REG"`
	Contact_Variables_Name_CON                                      string             `bson:"contact_Variables_Name_CON" json:"contact_Variables_Name_CON"`
	Contact_Variables_Name_EQ                                       string             `bson:"contact_Variables_Name_EQ" json:"contact_Variables_Name_EQ"`
	Contact_Variables_Name_REG                                      string             `bson:"contact_Variables_Name_REG" json:"contact_Variables_Name_REG"`
	Country_EQ                                                      CountryIsoCode     `bson:"country_EQ" json:"country_EQ"`
	DateCreated_EQ                                                  primitive.DateTime `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                                                 primitive.DateTime `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                                                 primitive.DateTime `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateEnded_EQ                                                    primitive.DateTime `bson:"dateEnded_EQ" json:"dateEnded_EQ"`
	DateEnded_GTE                                                   primitive.DateTime `bson:"dateEnded_GTE" json:"dateEnded_GTE"`
	DateEnded_LTE                                                   primitive.DateTime `bson:"dateEnded_LTE" json:"dateEnded_LTE"`
	DateUpdated_EQ                                                  primitive.DateTime `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                                                 primitive.DateTime `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                                                 primitive.DateTime `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	DigitsSent_CON                                                  string             `bson:"digitsSent_CON" json:"digitsSent_CON"`
	DigitsSent_EQ                                                   string             `bson:"digitsSent_EQ" json:"digitsSent_EQ"`
	DigitsSent_REG                                                  string             `bson:"digitsSent_REG" json:"digitsSent_REG"`
	DisabledVideo_EQ                                                bool               `bson:"disabledVideo_EQ" json:"disabledVideo_EQ"`
	From_CON                                                        string             `bson:"from_CON" json:"from_CON"`
	From_EQ                                                         string             `bson:"from_EQ" json:"from_EQ"`
	From_REG                                                        string             `bson:"from_REG" json:"from_REG"`
	FromCountry_EQ                                                  CountryIsoCode     `bson:"fromCountry_EQ" json:"fromCountry_EQ"`
	Id_CON                                                          string             `bson:"id_CON" json:"id_CON"`
	Id_EQ                                                           string             `bson:"id_EQ" json:"id_EQ"`
	Id_REG                                                          string             `bson:"id_REG" json:"id_REG"`
	IdLineThatInitiatedCall_CON                                     string             `bson:"idLineThatInitiatedCall_CON" json:"idLineThatInitiatedCall_CON"`
	IdLineThatInitiatedCall_EQ                                      string             `bson:"idLineThatInitiatedCall_EQ" json:"idLineThatInitiatedCall_EQ"`
	IdLineThatInitiatedCall_REG                                     string             `bson:"idLineThatInitiatedCall_REG" json:"idLineThatInitiatedCall_REG"`
	IdVoicemail_CON                                                 string             `bson:"idVoicemail_CON" json:"idVoicemail_CON"`
	IdVoicemail_EQ                                                  string             `bson:"idVoicemail_EQ" json:"idVoicemail_EQ"`
	IdVoicemail_REG                                                 string             `bson:"idVoicemail_REG" json:"idVoicemail_REG"`
	IsInternational_EQ                                              bool               `bson:"isInternational_EQ" json:"isInternational_EQ"`
	Recording_ErrorMessage_CON                                      string             `bson:"recording_ErrorMessage_CON" json:"recording_ErrorMessage_CON"`
	Recording_ErrorMessage_EQ                                       string             `bson:"recording_ErrorMessage_EQ" json:"recording_ErrorMessage_EQ"`
	Recording_ErrorMessage_REG                                      string             `bson:"recording_ErrorMessage_REG" json:"recording_ErrorMessage_REG"`
	Recording_Id_CON                                                string             `bson:"recording_Id_CON" json:"recording_Id_CON"`
	Recording_Id_EQ                                                 string             `bson:"recording_Id_EQ" json:"recording_Id_EQ"`
	Recording_Id_REG                                                string             `bson:"recording_Id_REG" json:"recording_Id_REG"`
	Recording_RecordingDurationInSeconds_EQ                         int32              `bson:"recording_RecordingDurationInSeconds_EQ" json:"recording_RecordingDurationInSeconds_EQ"`
	Recording_RecordingDurationInSeconds_GTE                        int32              `bson:"recording_RecordingDurationInSeconds_GTE" json:"recording_RecordingDurationInSeconds_GTE"`
	Recording_RecordingDurationInSeconds_LTE                        int32              `bson:"recording_RecordingDurationInSeconds_LTE" json:"recording_RecordingDurationInSeconds_LTE"`
	Recording_RecordingMp3_FileSizeInBytes_EQ                       int32              `bson:"recording_RecordingMp3_FileSizeInBytes_EQ" json:"recording_RecordingMp3_FileSizeInBytes_EQ"`
	Recording_RecordingMp3_FileSizeInBytes_GTE                      int32              `bson:"recording_RecordingMp3_FileSizeInBytes_GTE" json:"recording_RecordingMp3_FileSizeInBytes_GTE"`
	Recording_RecordingMp3_FileSizeInBytes_LTE                      int32              `bson:"recording_RecordingMp3_FileSizeInBytes_LTE" json:"recording_RecordingMp3_FileSizeInBytes_LTE"`
	Recording_RecordingMp3_Id_CON                                   string             `bson:"recording_RecordingMp3_Id_CON" json:"recording_RecordingMp3_Id_CON"`
	Recording_RecordingMp3_Id_EQ                                    string             `bson:"recording_RecordingMp3_Id_EQ" json:"recording_RecordingMp3_Id_EQ"`
	Recording_RecordingMp3_Id_REG                                   string             `bson:"recording_RecordingMp3_Id_REG" json:"recording_RecordingMp3_Id_REG"`
	Recording_RecordingMp3_Md5Hash_CON                              string             `bson:"recording_RecordingMp3_Md5Hash_CON" json:"recording_RecordingMp3_Md5Hash_CON"`
	Recording_RecordingMp3_Md5Hash_EQ                               string             `bson:"recording_RecordingMp3_Md5Hash_EQ" json:"recording_RecordingMp3_Md5Hash_EQ"`
	Recording_RecordingMp3_Md5Hash_REG                              string             `bson:"recording_RecordingMp3_Md5Hash_REG" json:"recording_RecordingMp3_Md5Hash_REG"`
	Recording_RecordingMp3_Url_CON                                  string             `bson:"recording_RecordingMp3_Url_CON" json:"recording_RecordingMp3_Url_CON"`
	Recording_RecordingMp3_Url_EQ                                   string             `bson:"recording_RecordingMp3_Url_EQ" json:"recording_RecordingMp3_Url_EQ"`
	Recording_RecordingMp3_Url_REG                                  string             `bson:"recording_RecordingMp3_Url_REG" json:"recording_RecordingMp3_Url_REG"`
	SecondsItTookToAnswer_EQ                                        int32              `bson:"secondsItTookToAnswer_EQ" json:"secondsItTookToAnswer_EQ"`
	SecondsItTookToAnswer_GTE                                       int32              `bson:"secondsItTookToAnswer_GTE" json:"secondsItTookToAnswer_GTE"`
	SecondsItTookToAnswer_LTE                                       int32              `bson:"secondsItTookToAnswer_LTE" json:"secondsItTookToAnswer_LTE"`
	Status_CON                                                      string             `bson:"status_CON" json:"status_CON"`
	Status_EQ                                                       string             `bson:"status_EQ" json:"status_EQ"`
	Status_REG                                                      string             `bson:"status_REG" json:"status_REG"`
	TimesWhenCallPlacedOnHold_SecondsElapsedWhenPlacedOnHold_EQ     int32              `bson:"timesWhenCallPlacedOnHold_SecondsElapsedWhenPlacedOnHold_EQ" json:"timesWhenCallPlacedOnHold_SecondsElapsedWhenPlacedOnHold_EQ"`
	TimesWhenCallPlacedOnHold_SecondsElapsedWhenPlacedOnHold_GTE    int32              `bson:"timesWhenCallPlacedOnHold_SecondsElapsedWhenPlacedOnHold_GTE" json:"timesWhenCallPlacedOnHold_SecondsElapsedWhenPlacedOnHold_GTE"`
	TimesWhenCallPlacedOnHold_SecondsElapsedWhenPlacedOnHold_LTE    int32              `bson:"timesWhenCallPlacedOnHold_SecondsElapsedWhenPlacedOnHold_LTE" json:"timesWhenCallPlacedOnHold_SecondsElapsedWhenPlacedOnHold_LTE"`
	TimesWhenCallPlacedOnHold_SecondsElapsedWhenRemovedFromHold_EQ  int32              `bson:"timesWhenCallPlacedOnHold_SecondsElapsedWhenRemovedFromHold_EQ" json:"timesWhenCallPlacedOnHold_SecondsElapsedWhenRemovedFromHold_EQ"`
	TimesWhenCallPlacedOnHold_SecondsElapsedWhenRemovedFromHold_GTE int32              `bson:"timesWhenCallPlacedOnHold_SecondsElapsedWhenRemovedFromHold_GTE" json:"timesWhenCallPlacedOnHold_SecondsElapsedWhenRemovedFromHold_GTE"`
	TimesWhenCallPlacedOnHold_SecondsElapsedWhenRemovedFromHold_LTE int32              `bson:"timesWhenCallPlacedOnHold_SecondsElapsedWhenRemovedFromHold_LTE" json:"timesWhenCallPlacedOnHold_SecondsElapsedWhenRemovedFromHold_LTE"`
	To_CON                                                          string             `bson:"to_CON" json:"to_CON"`
	To_EQ                                                           string             `bson:"to_EQ" json:"to_EQ"`
	To_REG                                                          string             `bson:"to_REG" json:"to_REG"`
	ToCountry_EQ                                                    CountryIsoCode     `bson:"toCountry_EQ" json:"toCountry_EQ"`
	ToInternationalFormat_CON                                       string             `bson:"toInternationalFormat_CON" json:"toInternationalFormat_CON"`
	ToInternationalFormat_EQ                                        string             `bson:"toInternationalFormat_EQ" json:"toInternationalFormat_EQ"`
	ToInternationalFormat_REG                                       string             `bson:"toInternationalFormat_REG" json:"toInternationalFormat_REG"`
}

// BUILDER from bson map:
func BuildCallOutgoingToPSTNFilterRequest(m map[string]interface{}, x *CallOutgoingToPSTNFilterRequest) {
	if val, ok := m["callType_EQ"]; ok && val != nil {
		x.CallType_EQ = CallType("CallType_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["channelVariables_CallerIdName_CON"]; ok && val != nil {
		x.ChannelVariables_CallerIdName_CON = val.(string)
	}
	if val, ok := m["channelVariables_CallerIdName_EQ"]; ok && val != nil {
		x.ChannelVariables_CallerIdName_EQ = val.(string)
	}
	if val, ok := m["channelVariables_CallerIdName_REG"]; ok && val != nil {
		x.ChannelVariables_CallerIdName_REG = val.(string)
	}
	if val, ok := m["channelVariables_CallerIdNumber_CON"]; ok && val != nil {
		x.ChannelVariables_CallerIdNumber_CON = val.(string)
	}
	if val, ok := m["channelVariables_CallerIdNumber_EQ"]; ok && val != nil {
		x.ChannelVariables_CallerIdNumber_EQ = val.(string)
	}
	if val, ok := m["channelVariables_CallerIdNumber_REG"]; ok && val != nil {
		x.ChannelVariables_CallerIdNumber_REG = val.(string)
	}
	if val, ok := m["channelVariables_IdMusicOnHold_CON"]; ok && val != nil {
		x.ChannelVariables_IdMusicOnHold_CON = val.(string)
	}
	if val, ok := m["channelVariables_IdMusicOnHold_EQ"]; ok && val != nil {
		x.ChannelVariables_IdMusicOnHold_EQ = val.(string)
	}
	if val, ok := m["channelVariables_IdMusicOnHold_REG"]; ok && val != nil {
		x.ChannelVariables_IdMusicOnHold_REG = val.(string)
	}
	if val, ok := m["channelVariables_Language_EQ"]; ok && val != nil {
		x.ChannelVariables_Language_EQ = Language("ChannelVariables_Language_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["childCalls_ChildCallType_EQ"]; ok && val != nil {
		x.ChildCalls_ChildCallType_EQ = ChildCallType("ChildCalls_ChildCallType_EQ_" + val.(string))
	} // is NOT readonly obtained from map
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
	if val, ok := m["contact_ContactEmails_SearchIndex_CON"]; ok && val != nil {
		x.Contact_ContactEmails_SearchIndex_CON = val.(string)
	}
	if val, ok := m["contact_ContactEmails_SearchIndex_EQ"]; ok && val != nil {
		x.Contact_ContactEmails_SearchIndex_EQ = val.(string)
	}
	if val, ok := m["contact_ContactEmails_SearchIndex_REG"]; ok && val != nil {
		x.Contact_ContactEmails_SearchIndex_REG = val.(string)
	}
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
	if val, ok := m["contact_ContactNumbers_SearchIndex_CON"]; ok && val != nil {
		x.Contact_ContactNumbers_SearchIndex_CON = val.(string)
	}
	if val, ok := m["contact_ContactNumbers_SearchIndex_EQ"]; ok && val != nil {
		x.Contact_ContactNumbers_SearchIndex_EQ = val.(string)
	}
	if val, ok := m["contact_ContactNumbers_SearchIndex_REG"]; ok && val != nil {
		x.Contact_ContactNumbers_SearchIndex_REG = val.(string)
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
	if val, ok := m["contact_IdIdentityUser_CON"]; ok && val != nil {
		x.Contact_IdIdentityUser_CON = val.(string)
	}
	if val, ok := m["contact_IdIdentityUser_EQ"]; ok && val != nil {
		x.Contact_IdIdentityUser_EQ = val.(string)
	}
	if val, ok := m["contact_IdIdentityUser_REG"]; ok && val != nil {
		x.Contact_IdIdentityUser_REG = val.(string)
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
	if val, ok := m["country_EQ"]; ok && val != nil {
		x.Country_EQ = CountryIsoCode("Country_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["dateCreated_EQ"]; ok && val != nil {
		x.DateCreated_EQ = val.(primitive.DateTime)
	}
	if val, ok := m["dateCreated_GTE"]; ok && val != nil {
		x.DateCreated_GTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateCreated_LTE"]; ok && val != nil {
		x.DateCreated_LTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateEnded_EQ"]; ok && val != nil {
		x.DateEnded_EQ = val.(primitive.DateTime)
	}
	if val, ok := m["dateEnded_GTE"]; ok && val != nil {
		x.DateEnded_GTE = val.(primitive.DateTime)
	}
	if val, ok := m["dateEnded_LTE"]; ok && val != nil {
		x.DateEnded_LTE = val.(primitive.DateTime)
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
	if val, ok := m["digitsSent_CON"]; ok && val != nil {
		x.DigitsSent_CON = val.(string)
	}
	if val, ok := m["digitsSent_EQ"]; ok && val != nil {
		x.DigitsSent_EQ = val.(string)
	}
	if val, ok := m["digitsSent_REG"]; ok && val != nil {
		x.DigitsSent_REG = val.(string)
	}
	if val, ok := m["disabledVideo_EQ"]; ok && val != nil {
		x.DisabledVideo_EQ = val.(bool)
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
	if val, ok := m["fromCountry_EQ"]; ok && val != nil {
		x.FromCountry_EQ = CountryIsoCode("FromCountry_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["id_CON"]; ok && val != nil {
		x.Id_CON = val.(string)
	}
	if val, ok := m["id_EQ"]; ok && val != nil {
		x.Id_EQ = val.(string)
	}
	if val, ok := m["id_REG"]; ok && val != nil {
		x.Id_REG = val.(string)
	}
	if val, ok := m["idLineThatInitiatedCall_CON"]; ok && val != nil {
		x.IdLineThatInitiatedCall_CON = val.(string)
	}
	if val, ok := m["idLineThatInitiatedCall_EQ"]; ok && val != nil {
		x.IdLineThatInitiatedCall_EQ = val.(string)
	}
	if val, ok := m["idLineThatInitiatedCall_REG"]; ok && val != nil {
		x.IdLineThatInitiatedCall_REG = val.(string)
	}
	if val, ok := m["idVoicemail_CON"]; ok && val != nil {
		x.IdVoicemail_CON = val.(string)
	}
	if val, ok := m["idVoicemail_EQ"]; ok && val != nil {
		x.IdVoicemail_EQ = val.(string)
	}
	if val, ok := m["idVoicemail_REG"]; ok && val != nil {
		x.IdVoicemail_REG = val.(string)
	}
	if val, ok := m["isInternational_EQ"]; ok && val != nil {
		x.IsInternational_EQ = val.(bool)
	}
	if val, ok := m["recording_ErrorMessage_CON"]; ok && val != nil {
		x.Recording_ErrorMessage_CON = val.(string)
	}
	if val, ok := m["recording_ErrorMessage_EQ"]; ok && val != nil {
		x.Recording_ErrorMessage_EQ = val.(string)
	}
	if val, ok := m["recording_ErrorMessage_REG"]; ok && val != nil {
		x.Recording_ErrorMessage_REG = val.(string)
	}
	if val, ok := m["recording_Id_CON"]; ok && val != nil {
		x.Recording_Id_CON = val.(string)
	}
	if val, ok := m["recording_Id_EQ"]; ok && val != nil {
		x.Recording_Id_EQ = val.(string)
	}
	if val, ok := m["recording_Id_REG"]; ok && val != nil {
		x.Recording_Id_REG = val.(string)
	}
	if val, ok := m["recording_RecordingDurationInSeconds_EQ"]; ok && val != nil {
		x.Recording_RecordingDurationInSeconds_EQ = val.(int32)
	}
	if val, ok := m["recording_RecordingDurationInSeconds_GTE"]; ok && val != nil {
		x.Recording_RecordingDurationInSeconds_GTE = val.(int32)
	}
	if val, ok := m["recording_RecordingDurationInSeconds_LTE"]; ok && val != nil {
		x.Recording_RecordingDurationInSeconds_LTE = val.(int32)
	}
	if val, ok := m["recording_RecordingMp3_FileSizeInBytes_EQ"]; ok && val != nil {
		x.Recording_RecordingMp3_FileSizeInBytes_EQ = val.(int32)
	}
	if val, ok := m["recording_RecordingMp3_FileSizeInBytes_GTE"]; ok && val != nil {
		x.Recording_RecordingMp3_FileSizeInBytes_GTE = val.(int32)
	}
	if val, ok := m["recording_RecordingMp3_FileSizeInBytes_LTE"]; ok && val != nil {
		x.Recording_RecordingMp3_FileSizeInBytes_LTE = val.(int32)
	}
	if val, ok := m["recording_RecordingMp3_Id_CON"]; ok && val != nil {
		x.Recording_RecordingMp3_Id_CON = val.(string)
	}
	if val, ok := m["recording_RecordingMp3_Id_EQ"]; ok && val != nil {
		x.Recording_RecordingMp3_Id_EQ = val.(string)
	}
	if val, ok := m["recording_RecordingMp3_Id_REG"]; ok && val != nil {
		x.Recording_RecordingMp3_Id_REG = val.(string)
	}
	if val, ok := m["recording_RecordingMp3_Md5Hash_CON"]; ok && val != nil {
		x.Recording_RecordingMp3_Md5Hash_CON = val.(string)
	}
	if val, ok := m["recording_RecordingMp3_Md5Hash_EQ"]; ok && val != nil {
		x.Recording_RecordingMp3_Md5Hash_EQ = val.(string)
	}
	if val, ok := m["recording_RecordingMp3_Md5Hash_REG"]; ok && val != nil {
		x.Recording_RecordingMp3_Md5Hash_REG = val.(string)
	}
	if val, ok := m["recording_RecordingMp3_Url_CON"]; ok && val != nil {
		x.Recording_RecordingMp3_Url_CON = val.(string)
	}
	if val, ok := m["recording_RecordingMp3_Url_EQ"]; ok && val != nil {
		x.Recording_RecordingMp3_Url_EQ = val.(string)
	}
	if val, ok := m["recording_RecordingMp3_Url_REG"]; ok && val != nil {
		x.Recording_RecordingMp3_Url_REG = val.(string)
	}
	if val, ok := m["secondsItTookToAnswer_EQ"]; ok && val != nil {
		x.SecondsItTookToAnswer_EQ = val.(int32)
	}
	if val, ok := m["secondsItTookToAnswer_GTE"]; ok && val != nil {
		x.SecondsItTookToAnswer_GTE = val.(int32)
	}
	if val, ok := m["secondsItTookToAnswer_LTE"]; ok && val != nil {
		x.SecondsItTookToAnswer_LTE = val.(int32)
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
	if val, ok := m["timesWhenCallPlacedOnHold_SecondsElapsedWhenPlacedOnHold_EQ"]; ok && val != nil {
		x.TimesWhenCallPlacedOnHold_SecondsElapsedWhenPlacedOnHold_EQ = val.(int32)
	}
	if val, ok := m["timesWhenCallPlacedOnHold_SecondsElapsedWhenPlacedOnHold_GTE"]; ok && val != nil {
		x.TimesWhenCallPlacedOnHold_SecondsElapsedWhenPlacedOnHold_GTE = val.(int32)
	}
	if val, ok := m["timesWhenCallPlacedOnHold_SecondsElapsedWhenPlacedOnHold_LTE"]; ok && val != nil {
		x.TimesWhenCallPlacedOnHold_SecondsElapsedWhenPlacedOnHold_LTE = val.(int32)
	}
	if val, ok := m["timesWhenCallPlacedOnHold_SecondsElapsedWhenRemovedFromHold_EQ"]; ok && val != nil {
		x.TimesWhenCallPlacedOnHold_SecondsElapsedWhenRemovedFromHold_EQ = val.(int32)
	}
	if val, ok := m["timesWhenCallPlacedOnHold_SecondsElapsedWhenRemovedFromHold_GTE"]; ok && val != nil {
		x.TimesWhenCallPlacedOnHold_SecondsElapsedWhenRemovedFromHold_GTE = val.(int32)
	}
	if val, ok := m["timesWhenCallPlacedOnHold_SecondsElapsedWhenRemovedFromHold_LTE"]; ok && val != nil {
		x.TimesWhenCallPlacedOnHold_SecondsElapsedWhenRemovedFromHold_LTE = val.(int32)
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
	if val, ok := m["toCountry_EQ"]; ok && val != nil {
		x.ToCountry_EQ = CountryIsoCode("ToCountry_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["toInternationalFormat_CON"]; ok && val != nil {
		x.ToInternationalFormat_CON = val.(string)
	}
	if val, ok := m["toInternationalFormat_EQ"]; ok && val != nil {
		x.ToInternationalFormat_EQ = val.(string)
	}
	if val, ok := m["toInternationalFormat_REG"]; ok && val != nil {
		x.ToInternationalFormat_REG = val.(string)
	}
}
