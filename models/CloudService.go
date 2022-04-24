package models

import . "github.com/ublux/go-models/enums"

type CloudService interface {
	IUbluxDocument
	IUbluxDocumentId

	GetIdIdentity() string
	GetCloudServiceType() CloudServiceType
	GetCountryIsoCode() CountryIsoCode
	GetLocalnet() string
	GetExternalIps() []string
	GetIsFailover() bool
	GetNAT() bool
	GetIsHealthy() bool
}
