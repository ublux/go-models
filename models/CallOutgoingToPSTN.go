package models

import (
	. "github.com/ublux/go-models/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CallOutgoingToPSTN struct {
	Id                        string                     `bson:"_id" json:"id"`
	CallType                  CallType                   `bson:"callType" json:"callType"`
	ChannelVariables          ChannelVariables           `bson:"channelVariables" json:"channelVariables"`
	ChildCalls                []ChildCall                `bson:"childCalls" json:"childCalls"`
	Contact                   Contact                    `bson:"contact" json:"contact"`
	Country                   CountryIsoCode             `bson:"country" json:"country"`
	DateCreated               string                     `bson:"dateCreated" json:"dateCreated"`
	DateDeleted               string                     `bson:"dateDeleted" json:"dateDeleted"`
	DateEnded                 string                     `bson:"dateEnded" json:"dateEnded"`
	DateUpdated               string                     `bson:"dateUpdated" json:"dateUpdated"`
	DigitsSent                []string                   `bson:"digitsSent" json:"digitsSent"`
	DisabledVideo             bool                       `bson:"disabledVideo" json:"disabledVideo"`
	From                      string                     `bson:"from" json:"from"`
	FromCountry               CountryIsoCode             `bson:"fromCountry" json:"fromCountry"`
	IdAccount                 string                     `bson:"idAccount" json:"idAccount"`
	IdLineThatInitiatedCall   string                     `bson:"idLineThatInitiatedCall" json:"idLineThatInitiatedCall"`
	IdTrunkTermination        string                     `bson:"idTrunkTermination" json:"idTrunkTermination"`
	IdVoicemail               string                     `bson:"idVoicemail" json:"idVoicemail"`
	IsInternational           bool                       `bson:"isInternational" json:"isInternational"`
	Recording                 Recording                  `bson:"recording" json:"recording"`
	SecondsItTookToAnswer     int32                      `bson:"secondsItTookToAnswer" json:"secondsItTookToAnswer"`
	Status                    string                     `bson:"status" json:"status"`
	TimesWhenCallPlacedOnHold []TimeWhenCallPlacedOnHold `bson:"timesWhenCallPlacedOnHold" json:"timesWhenCallPlacedOnHold"`
	To                        string                     `bson:"to" json:"to"`
	ToCountry                 CountryIsoCode             `bson:"toCountry" json:"toCountry"`
	ToInternationalFormat     string                     `bson:"toInternationalFormat" json:"toInternationalFormat"`
}

// Implementing interface IUbluxDocument
func (x CallOutgoingToPSTN) GetDateCreated() string {
	return x.DateCreated
}
func (x CallOutgoingToPSTN) GetDateDeleted() string {
	return x.DateDeleted
}
func (x CallOutgoingToPSTN) GetDateUpdated() string {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x CallOutgoingToPSTN) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x CallOutgoingToPSTN) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface CallOutgoing
func (x CallOutgoingToPSTN) GetContact() Contact {
	return x.Contact
}
func (x CallOutgoingToPSTN) GetIdLineThatInitiatedCall() string {
	return x.IdLineThatInitiatedCall
}

// Implementing interface Call
func (x CallOutgoingToPSTN) GetIdVoicemail() string {
	return x.IdVoicemail
}
func (x CallOutgoingToPSTN) GetChannelVariables() ChannelVariables {
	return x.ChannelVariables
}
func (x CallOutgoingToPSTN) GetChildCalls() []ChildCall {
	return x.ChildCalls
}
func (x CallOutgoingToPSTN) GetDateEnded() string {
	return x.DateEnded
}
func (x CallOutgoingToPSTN) GetStatus() string {
	return x.Status
}
func (x CallOutgoingToPSTN) GetSecondsItTookToAnswer() int32 {
	return x.SecondsItTookToAnswer
}
func (x CallOutgoingToPSTN) GetTimesWhenCallPlacedOnHold() []TimeWhenCallPlacedOnHold {
	return x.TimesWhenCallPlacedOnHold
}
func (x CallOutgoingToPSTN) GetFrom() string {
	return x.From
}
func (x CallOutgoingToPSTN) GetFromCountry() CountryIsoCode {
	return x.FromCountry
}
func (x CallOutgoingToPSTN) GetTo() string {
	return x.To
}
func (x CallOutgoingToPSTN) GetToCountry() CountryIsoCode {
	return x.ToCountry
}
func (x CallOutgoingToPSTN) GetCallType() CallType {
	return x.CallType
}
func (x CallOutgoingToPSTN) GetRecording() Recording {
	return x.Recording
}
func (x CallOutgoingToPSTN) GetDisabledVideo() bool {
	return x.DisabledVideo
}
func (x CallOutgoingToPSTN) GetDigitsSent() []string {
	return x.DigitsSent
}
func (x CallOutgoingToPSTN) GetIsInternational() bool {
	return x.IsInternational
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildCallOutgoingToPSTN(m map[string]interface{}, x *CallOutgoingToPSTN) {
	if val, ok := m["_id"]; ok && val != nil {
		x.Id = val.(string)
	}
	x.CallType = CallType_OutgoingToPSTN // readonly property
	if val, ok := m["channelVariables"]; ok && val != nil {
		BuildChannelVariables(val.(map[string]interface{}), &x.ChannelVariables)
	}
	if val, ok := m["childCalls"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok {
			for _, val = range array {
				if val != nil {
					// determine the type to see what type of object we instantiate
					t := val.(map[string]interface{})["_t"].(string)
					switch t {
					case "ChildCallAttendantTransferToExtension":
						tmp := new(ChildCallAttendantTransferToExtension)
						BuildChildCallAttendantTransferToExtension(val.(map[string]interface{}), tmp)
						x.ChildCalls = append(x.ChildCalls, *tmp)
					case "ChildCallAttendantTransferToPSTN":
						tmp := new(ChildCallAttendantTransferToPSTN)
						BuildChildCallAttendantTransferToPSTN(val.(map[string]interface{}), tmp)
						x.ChildCalls = append(x.ChildCalls, *tmp)
					case "ChildCallBlindTransferToExtension":
						tmp := new(ChildCallBlindTransferToExtension)
						BuildChildCallBlindTransferToExtension(val.(map[string]interface{}), tmp)
						x.ChildCalls = append(x.ChildCalls, *tmp)
					case "ChildCallBlindTransferToPSTN":
						tmp := new(ChildCallBlindTransferToPSTN)
						BuildChildCallBlindTransferToPSTN(val.(map[string]interface{}), tmp)
						x.ChildCalls = append(x.ChildCalls, *tmp)
					case "ChildCallForwardToExtension":
						tmp := new(ChildCallForwardToExtension)
						BuildChildCallForwardToExtension(val.(map[string]interface{}), tmp)
						x.ChildCalls = append(x.ChildCalls, *tmp)
					default:
						// error
					}
				}
			}
		}
	}
	if val, ok := m["contact"]; ok && val != nil {
		BuildContact(val.(map[string]interface{}), &x.Contact)
	}
	if val, ok := m["country"]; ok && val != nil {
		x.Country = CountryIsoCode("Country_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(string)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(string)
	}
	if val, ok := m["dateEnded"]; ok && val != nil {
		x.DateEnded = val.(string)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(string)
	}
	if val, ok := m["digitsSent"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.DigitsSent = append(x.DigitsSent, val.(string))
				}
			}
		}
	}
	if val, ok := m["disabledVideo"]; ok && val != nil {
		x.DisabledVideo = val.(bool)
	}
	if val, ok := m["from"]; ok && val != nil {
		x.From = val.(string)
	}
	if val, ok := m["fromCountry"]; ok && val != nil {
		x.FromCountry = CountryIsoCode("FromCountry_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["idLineThatInitiatedCall"]; ok && val != nil {
		x.IdLineThatInitiatedCall = val.(string)
	}
	if val, ok := m["idTrunkTermination"]; ok && val != nil {
		x.IdTrunkTermination = val.(string)
	}
	if val, ok := m["idVoicemail"]; ok && val != nil {
		x.IdVoicemail = val.(string)
	}
	if val, ok := m["isInternational"]; ok && val != nil {
		x.IsInternational = val.(bool)
	}
	if val, ok := m["recording"]; ok && val != nil {
		BuildRecording(val.(map[string]interface{}), &x.Recording)
	}
	if val, ok := m["secondsItTookToAnswer"]; ok && val != nil {
		x.SecondsItTookToAnswer = val.(int32)
	}
	if val, ok := m["status"]; ok && val != nil {
		x.Status = val.(string)
	}
	if val, ok := m["timesWhenCallPlacedOnHold"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok {
			for _, val = range array {
				if val != nil {
					tmp := new(TimeWhenCallPlacedOnHold)
					BuildTimeWhenCallPlacedOnHold(val.(map[string]interface{}), tmp)
					x.TimesWhenCallPlacedOnHold = append(x.TimesWhenCallPlacedOnHold, *tmp)
				}
			}
		}
	}
	if val, ok := m["to"]; ok && val != nil {
		x.To = val.(string)
	}
	if val, ok := m["toCountry"]; ok && val != nil {
		x.ToCountry = CountryIsoCode("ToCountry_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["toInternationalFormat"]; ok && val != nil {
		x.ToInternationalFormat = val.(string)
	}
}
