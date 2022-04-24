package enums

type EventActionType string

const (
	EventActionType_ForwardToExtension   EventActionType = "ForwardToExtension"
	EventActionType_ForwardToPhoneNumber EventActionType = "ForwardToPhoneNumber"
	EventActionType_LeaveVoicemail       EventActionType = "LeaveVoicemail"
)
