package enums

type ChildCallType string

const (
	ChildCallType_ForwardToExtension           ChildCallType = "ForwardToExtension"
	ChildCallType_AttendantTransferToExtension ChildCallType = "AttendantTransferToExtension"
	ChildCallType_AttendantTransferToPSTN      ChildCallType = "AttendantTransferToPSTN"
	ChildCallType_BlindTransferToExtension     ChildCallType = "BlindTransferToExtension"
	ChildCallType_BlindTransferToPSTN          ChildCallType = "BlindTransferToPSTN"
)
