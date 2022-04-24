package enums

type CallType string

const (
	CallType_IncomingToCallFlow  CallType = "IncomingToCallFlow"
	CallType_IncomingToExtension CallType = "IncomingToExtension"
	CallType_OutgoingToExtension CallType = "OutgoingToExtension"
	CallType_OutgoingToPSTN      CallType = "OutgoingToPSTN"
	CallType_FeatureVoicemail    CallType = "FeatureVoicemail"
)
