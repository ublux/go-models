package enums

type ExtensionType string

const (
	ExtensionType_CallFlow   ExtensionType = "CallFlow"
	ExtensionType_Conference ExtensionType = "Conference"
	ExtensionType_Dial       ExtensionType = "Dial"
	ExtensionType_Queue      ExtensionType = "Queue"
	ExtensionType_Voicemail  ExtensionType = "Voicemail"
)
