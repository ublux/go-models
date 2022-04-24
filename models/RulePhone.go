package models

import . "github.com/ublux/go-models/enums"
import "go.mongodb.org/mongo-driver/bson/primitive"

type RulePhone struct {
	DaysOfWeek  []DayOfWeek `bson:"daysOfWeek" json:"daysOfWeek"`
	IdCallFlow  string      `bson:"idCallFlow" json:"idCallFlow"`
	IdExtension string      `bson:"idExtension" json:"idExtension"`
}

// BUILDER from bson map:
func BuildRulePhone(m map[string]interface{}, x *RulePhone) {
	if val, ok := m["daysOfWeek"]; ok && val != nil {
		if array, ok := (val).(primitive.A); ok { // array case
			for _, val = range array {
				x.DaysOfWeek = append(x.DaysOfWeek, DayOfWeek("DaysOfWeek_"+val.(string)))
			} // is NOT readonly obtained from map
		}
	}
	if val, ok := m["idCallFlow"]; ok && val != nil {
		x.IdCallFlow = val.(string)
	}
	if val, ok := m["idExtension"]; ok && val != nil {
		x.IdExtension = val.(string)
	}
}
