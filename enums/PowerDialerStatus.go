package enums

type PowerDialerStatus string

const (
	PowerDialerStatus_ToBeStarted       PowerDialerStatus = "ToBeStarted"
	PowerDialerStatus_Pending           PowerDialerStatus = "Pending"
	PowerDialerStatus_InProgress        PowerDialerStatus = "InProgress"
	PowerDialerStatus_Failed            PowerDialerStatus = "Failed"
	PowerDialerStatus_ExceededDateLimit PowerDialerStatus = "ExceededDateLimit"
	PowerDialerStatus_Completed         PowerDialerStatus = "Completed"
	PowerDialerStatus_Canceled          PowerDialerStatus = "Canceled"
	PowerDialerStatus_Retrying          PowerDialerStatus = "Retrying"
	PowerDialerStatus_Paused            PowerDialerStatus = "Paused"
)
