package models

import . "github.com/ublux/go-models/enums"

import "go.mongodb.org/mongo-driver/bson/primitive"

type Call interface {
	IUbluxDocument
	IUbluxDocumentId
	IReferncesAccount

	GetIdAccount() string
	GetIdVoicemail() string
	GetChannelVariables() ChannelVariables
	GetChildCalls() []ChildCall
	GetContact() Contact
	GetDateEnded() primitive.DateTime
	GetStatus() string
	GetSecondsItTookToAnswer() int32
	GetTimesWhenCallPlacedOnHold() []TimeWhenCallPlacedOnHold
	GetFrom() string
	GetFromCountry() CountryIsoCode
	GetTo() string
	GetToCountry() CountryIsoCode
	GetCallType() CallType
	GetRecording() Recording
	GetDisabledVideo() bool
	GetDigitsSent() []string
	GetIsInternational() bool
}
