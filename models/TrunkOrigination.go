package models

import . "github.com/ublux/go-models/enums"

type TrunkOrigination interface {
	IUbluxDocument
	IUbluxDocumentId

	GetIdVoipProvider() string
	GetIdCloudServicePbx() string
	GetIdCloudServicePbxFailover() string
	GetIdsVoipNumbers() []string
	GetTrunkOriginationType() TrunkOriginationType
	GetProviderId() string
	GetFriendlyName() string
}
