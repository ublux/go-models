package models

import . "github.com/ublux/go-models/enums"
import "time"
import "go.mongodb.org/mongo-driver/bson/primitive"

type CallIncomingToExtension struct {
	CallType                  CallType                   `bson:"callType" json:"callType"`
	ChannelVariables          ChannelVariables           `bson:"channelVariables" json:"channelVariables"`
	ChildCalls                []ChildCall                `bson:"childCalls" json:"childCalls"`
	Contact                   Contact                    `bson:"contact" json:"contact"`
	DateCreated               time.Time                  `bson:"dateCreated" json:"dateCreated"`
	DateDeleted               time.Time                  `bson:"dateDeleted" json:"dateDeleted"`
	DateEnded                 time.Time                  `bson:"dateEnded" json:"dateEnded"`
	DateUpdated               time.Time                  `bson:"dateUpdated" json:"dateUpdated"`
	DigitsSent                []string                   `bson:"digitsSent" json:"digitsSent"`
	DisabledVideo             bool                       `bson:"disabledVideo" json:"disabledVideo"`
	ExtensionFriendlyName     string                     `bson:"extensionFriendlyName" json:"extensionFriendlyName"`
	ExtensionNumber           string                     `bson:"extensionNumber" json:"extensionNumber"`
	From                      string                     `bson:"from" json:"from"`
	FromCountry               CountryIsoCode             `bson:"fromCountry" json:"fromCountry"`
	FromInternationalFormat   string                     `bson:"fromInternationalFormat" json:"fromInternationalFormat"`
	Id                        string                     `bson:"id" json:"id"`
	IdAccount                 string                     `bson:"idAccount" json:"idAccount"`
	IdExtension               string                     `bson:"idExtension" json:"idExtension"`
	IdLineThatAnswered        string                     `bson:"idLineThatAnswered" json:"idLineThatAnswered"`
	IdsLinesThatDidNotRing    []string                   `bson:"idsLinesThatDidNotRing" json:"idsLinesThatDidNotRing"`
	IdsLinesThatRing          []string                   `bson:"idsLinesThatRing" json:"idsLinesThatRing"`
	IdVoicemail               string                     `bson:"idVoicemail" json:"idVoicemail"`
	IdVoipNumberPhone         string                     `bson:"idVoipNumberPhone" json:"idVoipNumberPhone"`
	IdVoipProvider            string                     `bson:"idVoipProvider" json:"idVoipProvider"`
	IsInternational           bool                       `bson:"isInternational" json:"isInternational"`
	ProviderSid               string                     `bson:"providerSid" json:"providerSid"`
	Recording                 Recording                  `bson:"recording" json:"recording"`
	SecondsItTookToAnswer     int32                      `bson:"secondsItTookToAnswer" json:"secondsItTookToAnswer"`
	Status                    string                     `bson:"status" json:"status"`
	TimesWhenCallPlacedOnHold []TimeWhenCallPlacedOnHold `bson:"timesWhenCallPlacedOnHold" json:"timesWhenCallPlacedOnHold"`
	To                        string                     `bson:"to" json:"to"`
	ToCountry                 CountryIsoCode             `bson:"toCountry" json:"toCountry"`
}

// Implementing interface IUbluxDocument
func (x CallIncomingToExtension) GetDateDeleted() time.Time {
	return x.DateDeleted
}
func (x CallIncomingToExtension) GetDateCreated() time.Time {
	return x.DateCreated
}
func (x CallIncomingToExtension) GetDateUpdated() time.Time {
	return x.DateUpdated
}

// Implementing interface IUbluxDocumentId
func (x CallIncomingToExtension) GetId() string {
	return x.Id
}

// Implementing interface IReferncesAccount
func (x CallIncomingToExtension) GetIdAccount() string {
	return x.IdAccount
}

// Implementing interface CallIncoming
func (x CallIncomingToExtension) GetProviderSid() string {
	return x.ProviderSid
}
func (x CallIncomingToExtension) GetIdVoipProvider() string {
	return x.IdVoipProvider
}
func (x CallIncomingToExtension) GetContact() Contact {
	return x.Contact
}
func (x CallIncomingToExtension) GetIdVoipNumberPhone() string {
	return x.IdVoipNumberPhone
}
func (x CallIncomingToExtension) GetFromInternationalFormat() string {
	return x.FromInternationalFormat
}

// Implementing interface Call
func (x CallIncomingToExtension) GetIdVoicemail() string {
	return x.IdVoicemail
}
func (x CallIncomingToExtension) GetChannelVariables() ChannelVariables {
	return x.ChannelVariables
}
func (x CallIncomingToExtension) GetChildCalls() []ChildCall {
	return x.ChildCalls
}
func (x CallIncomingToExtension) GetDateEnded() time.Time {
	return x.DateEnded
}
func (x CallIncomingToExtension) GetStatus() string {
	return x.Status
}
func (x CallIncomingToExtension) GetSecondsItTookToAnswer() int32 {
	return x.SecondsItTookToAnswer
}
func (x CallIncomingToExtension) GetTimesWhenCallPlacedOnHold() []TimeWhenCallPlacedOnHold {
	return x.TimesWhenCallPlacedOnHold
}
func (x CallIncomingToExtension) GetFrom() string {
	return x.From
}
func (x CallIncomingToExtension) GetFromCountry() CountryIsoCode {
	return x.FromCountry
}
func (x CallIncomingToExtension) GetTo() string {
	return x.To
}
func (x CallIncomingToExtension) GetToCountry() CountryIsoCode {
	return x.ToCountry
}
func (x CallIncomingToExtension) GetCallType() CallType {
	return x.CallType
}
func (x CallIncomingToExtension) GetRecording() Recording {
	return x.Recording
}
func (x CallIncomingToExtension) GetDisabledVideo() bool {
	return x.DisabledVideo
}
func (x CallIncomingToExtension) GetDigitsSent() []string {
	return x.DigitsSent
}
func (x CallIncomingToExtension) GetIsInternational() bool {
	return x.IsInternational
}

// Implementing interface UbluxDocument

// BUILDER from bson map:
func BuildCallIncomingToExtension(m map[string]interface{}, x *CallIncomingToExtension) {
	x.CallType = CallType_IncomingToExtension // readonly property
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
		x.DateCreated = val.(time.Time)
	}
	if val, ok := m["dateDeleted"]; ok && val != nil {
		x.DateDeleted = val.(time.Time)
	}
	if val, ok := m["dateEnded"]; ok && val != nil {
		x.DateEnded = val.(time.Time)
	}
	if val, ok := m["dateUpdated"]; ok && val != nil {
		x.DateUpdated = val.(time.Time)
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
	if val, ok := m["fromInternationalFormat"]; ok && val != nil {
		x.FromInternationalFormat = val.(string)
	}
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
	if val, ok := m["idVoipNumberPhone"]; ok && val != nil {
		x.IdVoipNumberPhone = val.(string)
	}
	if val, ok := m["idVoipProvider"]; ok && val != nil {
		x.IdVoipProvider = val.(string)
	}
	if val, ok := m["isInternational"]; ok && val != nil {
		x.IsInternational = val.(bool)
	}
	if val, ok := m["providerSid"]; ok && val != nil {
		x.ProviderSid = val.(string)
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
