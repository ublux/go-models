package enums

type PowerDialerGroupStatus string

const (
	PowerDialerGroupStatus_ToBeStarted PowerDialerGroupStatus = "ToBeStarted"
	PowerDialerGroupStatus_Started     PowerDialerGroupStatus = "Started"
	PowerDialerGroupStatus_Paused      PowerDialerGroupStatus = "Paused"
	PowerDialerGroupStatus_Canceled    PowerDialerGroupStatus = "Canceled"
	PowerDialerGroupStatus_Failed      PowerDialerGroupStatus = "Failed"
	PowerDialerGroupStatus_Completed   PowerDialerGroupStatus = "Completed"
)
