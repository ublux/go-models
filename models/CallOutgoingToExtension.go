package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type CallOutgoingToExtension struct {
	CallType                  CallType                   `bson:"callType" json:"callType"`
	ChannelVariables          ChannelVariables           `bson:"channelVariables" json:"channelVariables"`
	ChildCalls                []ChildCall                `bson:"childCalls" json:"childCalls"`
	Contact                   Contact                    `bson:"contact" json:"contact"`
	DateCreated               primitive.DateTime         `bson:"dateCreated" json:"dateCreated"`
	DateDeleted               primitive.DateTime         `bson:"dateDeleted" json:"dateDeleted"`
	DateEnded                 primitive.DateTime         `bson:"dateEnded" json:"dateEnded"`
	DateUpdated               primitive.DateTime         `bson:"dateUpdated" json:"dateUpdated"`
	DigitsSent                []string                   `bson:"digitsSent" json:"digitsSent"`
	DisabledVideo             bool                       `bson:"disabledVideo" json:"disabledVideo"`
	ExtensionFriendlyName     string                     `bson:"extensionFriendlyName" json:"extensionFriendlyName"`
	ExtensionNumber           string                     `bson:"extensionNumber" json:"extensionNumber"`
	From                      string                     `bson:"from" json:"from"`
	FromCountry               CountryIsoCode             `bson:"fromCountry" json:"fromCountry"`
	Id                        string                     `bson:"_id" json:"id"`
	IdAccount                 string                     `bson:"idAccount" json:"idAccount"`
	IdExtension               string                     `bson:"idExtension" json:"idExtension"`
	IdLineThatAnswered        string                     `bson:"idLineThatAnswered" json:"idLineThatAnswered"`
	IdLineThatInitiatedCall   string                     `bson:"idLineThatInitiatedCall" json:"idLineThatInitiatedCall"`
	IdsLinesThatDidNotRing    []string                   `bson:"idsLinesThatDidNotRing" json:"idsLinesThatDidNotRing"`
	IdsLinesThatRing          []string                   `bson:"idsLinesThatRing" json:"idsLinesThatRing"`
	IdVoicemail               string                     `bson:"idVoicemail" json:"idVoicemail"`
	IsInternational           bool                       `bson:"isInternational" json:"isInternational"`
	Recording                 Recording                  `bson:"recording" json:"recording"`
	SecondsItTookToAnswer     int32                      `bson:"secondsItTookToAnswer" json:"secondsItTookToAnswer"`
	Status                    string                     `bson:"status" json:"status"`
	TimesWhenCallPlacedOnHold []TimeWhenCallPlacedOnHold `bson:"timesWhenCallPlacedOnHold" json:"timesWhenCallPlacedOnHold"`
	To                        string                     `bson:"to" json:"to"`
	ToCountry                 CountryIsoCode             `bson:"toCountry" json:"toCountry"`
}

// Implementing interface IUbluxDocument
func (x CallOutgoingToExtension) GetDateDeleted() primitive.DateTime {
	return x.DateDeleted
}
func (x CallOutgoingToExtension) GetDateCreated() primitive.DateTime {
	return x.DateCreated
}
func (x CallOutgoingToExtension) GetDateUpdated() primitive.DateTime {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x CallOutgoingToExtension) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x CallOutgoingToExtension) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface CallOutgoing
func (x CallOutgoingToExtension) GetContact() Contact {
	return x.Contact
}
func (x CallOutgoingToExtension) GetIdLineThatInitiatedCall() string {
	return x.IdLineThatInitiatedCall
}

// Implementing interface Call
func (x CallOutgoingToExtension) GetIdVoicemail() string {
	return x.IdVoicemail
}
func (x CallOutgoingToExtension) GetChannelVariables() ChannelVariables {
	return x.ChannelVariables
}
func (x CallOutgoingToExtension) GetChildCalls() []ChildCall {
	return x.ChildCalls
}
func (x CallOutgoingToExtension) GetDateEnded() primitive.DateTime {
	return x.DateEnded
}
func (x CallOutgoingToExtension) GetStatus() string {
	return x.Status
}
func (x CallOutgoingToExtension) GetSecondsItTookToAnswer() int32 {
	return x.SecondsItTookToAnswer
}
func (x CallOutgoingToExtension) GetTimesWhenCallPlacedOnHold() []TimeWhenCallPlacedOnHold {
	return x.TimesWhenCallPlacedOnHold
}
func (x CallOutgoingToExtension) GetFrom() string {
	return x.From
}
func (x CallOutgoingToExtension) GetFromCountry() CountryIsoCode {
	return x.FromCountry
}
func (x CallOutgoingToExtension) GetTo() string {
	return x.To
}
func (x CallOutgoingToExtension) GetToCountry() CountryIsoCode {
	return x.ToCountry
}
func (x CallOutgoingToExtension) GetCallType() CallType {
	return x.CallType
}
func (x CallOutgoingToExtension) GetRecording() Recording {
	return x.Recording
}
func (x CallOutgoingToExtension) GetDisabledVideo() bool {
	return x.DisabledVideo
}
func (x CallOutgoingToExtension) GetDigitsSent() []string {
	return x.DigitsSent
}
func (x CallOutgoingToExtension) GetIsInternational() bool {
	return x.IsInternational
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildCallOutgoingToExtension(m map[string]interface{}, x *CallOutgoingToExtension) {
	x.CallType = CallType_OutgoingToExtension // readonly property
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
	if val, ok := m["dateCreated"]; ok && val != nil {
		x.DateCreated = val.(primitive.DateTime)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(primitive.DateTime)
	}
	if val, ok := m["dateEnded"]; ok && val != nil {
		x.DateEnded = val.(primitive.DateTime)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(primitive.DateTime)
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
	if val, ok := m["extensionFriendlyName"]; ok && val != nil {
		x.ExtensionFriendlyName = val.(string)
	}
	if val, ok := m["extensionNumber"]; ok && val != nil {
		x.ExtensionNumber = val.(string)
	}
	if val, ok := m["from"]; ok && val != nil {
		x.From = val.(string)
	}
	if val, ok := m["fromCountry"]; ok && val != nil {
		x.FromCountry = CountryIsoCode("FromCountry_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["id"]; ok && val != nil {
		x.Id = val.(string)
	}
	if val, ok := m["idAccount"]; ok && val != nil {
		x.IdAccount = val.(string)
	}
	if val, ok := m["idExtension"]; ok && val != nil {
		x.IdExtension = val.(string)
	}
	if val, ok := m["idLineThatAnswered"]; ok && val != nil {
		x.IdLineThatAnswered = val.(string)
	}
	if val, ok := m["idLineThatInitiatedCall"]; ok && val != nil {
		x.IdLineThatInitiatedCall = val.(string)
	}
	if val, ok := m["idsLinesThatDidNotRing"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.IdsLinesThatDidNotRing = append(x.IdsLinesThatDidNotRing, val.(string))
				}
			}
		}
	}
	if val, ok := m["idsLinesThatRing"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				if val != nil {
					x.IdsLinesThatRing = append(x.IdsLinesThatRing, val.(string))
				}
			}
		}
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
}
