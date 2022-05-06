package models

import "time"
import . "github.com/ublux/go-models/enums"

type IdentityFilterRequest struct {
	AllowConnectingFromIpRegex_CON  string       `bson:"allowConnectingFromIpRegex_CON" json:"allowConnectingFromIpRegex_CON"`
	AllowConnectingFromIpRegex_EQ   string       `bson:"allowConnectingFromIpRegex_EQ" json:"allowConnectingFromIpRegex_EQ"`
	AllowConnectingFromIpRegex_REG  string       `bson:"allowConnectingFromIpRegex_REG" json:"allowConnectingFromIpRegex_REG"`
	DateAuthenticated_EQ            time.Time    `bson:"dateAuthenticated_EQ" json:"dateAuthenticated_EQ"`
	DateAuthenticated_GTE           time.Time    `bson:"dateAuthenticated_GTE" json:"dateAuthenticated_GTE"`
	DateAuthenticated_LTE           time.Time    `bson:"dateAuthenticated_LTE" json:"dateAuthenticated_LTE"`
	DateCreated_EQ                  time.Time    `bson:"dateCreated_EQ" json:"dateCreated_EQ"`
	DateCreated_GTE                 time.Time    `bson:"dateCreated_GTE" json:"dateCreated_GTE"`
	DateCreated_LTE                 time.Time    `bson:"dateCreated_LTE" json:"dateCreated_LTE"`
	DateUpdated_EQ                  time.Time    `bson:"dateUpdated_EQ" json:"dateUpdated_EQ"`
	DateUpdated_GTE                 time.Time    `bson:"dateUpdated_GTE" json:"dateUpdated_GTE"`
	DateUpdated_LTE                 time.Time    `bson:"dateUpdated_LTE" json:"dateUpdated_LTE"`
	Id_CON                          string       `bson:"id_CON" json:"id_CON"`
	Id_EQ                           string       `bson:"id_EQ" json:"id_EQ"`
	Id_REG                          string       `bson:"id_REG" json:"id_REG"`
	IdentityType_EQ                 IdentityType `bson:"identityType_EQ" json:"identityType_EQ"`
	IpAddressWhereAuthenticated_CON string       `bson:"ipAddressWhereAuthenticated_CON" json:"ipAddressWhereAuthenticated_CON"`
	IpAddressWhereAuthenticated_EQ  string       `bson:"ipAddressWhereAuthenticated_EQ" json:"ipAddressWhereAuthenticated_EQ"`
	IpAddressWhereAuthenticated_REG string       `bson:"ipAddressWhereAuthenticated_REG" json:"ipAddressWhereAuthenticated_REG"`
	Password_CON                    string       `bson:"password_CON" json:"password_CON"`
	Password_EQ                     string       `bson:"password_EQ" json:"password_EQ"`
	Password_REG                    string       `bson:"password_REG" json:"password_REG"`
	PreventConnectingIfIpChanges_EQ bool         `bson:"preventConnectingIfIpChanges_EQ" json:"preventConnectingIfIpChanges_EQ"`
	UbluxRoles_CON                  UbluxRole    `bson:"ubluxRoles_CON" json:"ubluxRoles_CON"`
	Username_CON                    string       `bson:"username_CON" json:"username_CON"`
	Username_EQ                     string       `bson:"username_EQ" json:"username_EQ"`
	Username_REG                    string       `bson:"username_REG" json:"username_REG"`
}

// BUILDER from bson map:
func BuildIdentityFilterRequest(m map[string]interface{}, x *IdentityFilterRequest) {
	if val, ok := m["allowConnectingFromIpRegex_CON"]; ok && val != nil {
		x.AllowConnectingFromIpRegex_CON = val.(string)
	}
	if val, ok := m["allowConnectingFromIpRegex_EQ"]; ok && val != nil {
		x.AllowConnectingFromIpRegex_EQ = val.(string)
	}
	if val, ok := m["allowConnectingFromIpRegex_REG"]; ok && val != nil {
		x.AllowConnectingFromIpRegex_REG = val.(string)
	}
	if val, ok := m["dateAuthenticated_EQ"]; ok && val != nil {
		x.DateAuthenticated_EQ = val.(time.Time)
	}
	if val, ok := m["dateAuthenticated_GTE"]; ok && val != nil {
		x.DateAuthenticated_GTE = val.(time.Time)
	}
	if val, ok := m["dateAuthenticated_LTE"]; ok && val != nil {
		x.DateAuthenticated_LTE = val.(time.Time)
	}
	if val, ok := m["dateCreated_EQ"]; ok && val != nil {
		x.DateCreated_EQ = val.(time.Time)
	}
	if val, ok := m["dateCreated_GTE"]; ok && val != nil {
		x.DateCreated_GTE = val.(time.Time)
	}
	if val, ok := m["dateCreated_LTE"]; ok && val != nil {
		x.DateCreated_LTE = val.(time.Time)
	}
	if val, ok := m["dateUpdated_EQ"]; ok && val != nil {
		x.DateUpdated_EQ = val.(time.Time)
	}
	if val, ok := m["dateUpdated_GTE"]; ok && val != nil {
		x.DateUpdated_GTE = val.(time.Time)
	}
	if val, ok := m["dateUpdated_LTE"]; ok && val != nil {
		x.DateUpdated_LTE = val.(time.Time)
	}
	if val, ok := m["id_CON"]; ok && val != nil {
		x.Id_CON = val.(string)
	}
	if val, ok := m["id_EQ"]; ok && val != nil {
		x.Id_EQ = val.(string)
	}
	if val, ok := m["id_REG"]; ok && val != nil {
		x.Id_REG = val.(string)
	}
	if val, ok := m["identityType_EQ"]; ok && val != nil {
		x.IdentityType_EQ = IdentityType("IdentityType_EQ_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["ipAddressWhereAuthenticated_CON"]; ok && val != nil {
		x.IpAddressWhereAuthenticated_CON = val.(string)
	}
	if val, ok := m["ipAddressWhereAuthenticated_EQ"]; ok && val != nil {
		x.IpAddressWhereAuthenticated_EQ = val.(string)
	}
	if val, ok := m["ipAddressWhereAuthenticated_REG"]; ok && val != nil {
		x.IpAddressWhereAuthenticated_REG = val.(string)
	}
	if val, ok := m["password_CON"]; ok && val != nil {
		x.Password_CON = val.(string)
	}
	if val, ok := m["password_EQ"]; ok && val != nil {
		x.Password_EQ = val.(string)
	}
	if val, ok := m["password_REG"]; ok && val != nil {
		x.Password_REG = val.(string)
	}
	if val, ok := m["preventConnectingIfIpChanges_EQ"]; ok && val != nil {
		x.PreventConnectingIfIpChanges_EQ = val.(bool)
	}
	if val, ok := m["ubluxRoles_CON"]; ok && val != nil {
		x.UbluxRoles_CON = UbluxRole("UbluxRoles_CON_" + val.(string))
	} // is NOT readonly obtained from map
	if val, ok := m["username_CON"]; ok && val != nil {
		x.Username_CON = val.(string)
	}
	if val, ok := m["username_EQ"]; ok && val != nil {
		x.Username_EQ = val.(string)
	}
	if val, ok := m["username_REG"]; ok && val != nil {
		x.Username_REG = val.(string)
	}
}
