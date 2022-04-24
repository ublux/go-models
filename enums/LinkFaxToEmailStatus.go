package enums

type LinkFaxToEmailStatus string

const (
	LinkFaxToEmailStatus_None                   LinkFaxToEmailStatus = "None"
	LinkFaxToEmailStatus_AskedUserForAccountPin LinkFaxToEmailStatus = "AskedUserForAccountPin"
	LinkFaxToEmailStatus_SendWelcomeEmail       LinkFaxToEmailStatus = "SendWelcomeEmail"
)
