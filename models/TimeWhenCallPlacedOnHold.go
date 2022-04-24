package models

type TimeWhenCallPlacedOnHold struct {
	SecondsElapsedWhenPlacedOnHold    int32 `bson:"secondsElapsedWhenPlacedOnHold" json:"secondsElapsedWhenPlacedOnHold"`
	SecondsElapsedWhenRemovedFromHold int32 `bson:"secondsElapsedWhenRemovedFromHold" json:"secondsElapsedWhenRemovedFromHold"`
}

// BUILDER from bson map:
func BuildTimeWhenCallPlacedOnHold(m map[string]interface{}, x *TimeWhenCallPlacedOnHold) {
	if val, ok := m["secondsElapsedWhenPlacedOnHold"]; ok && val != nil {
		x.SecondsElapsedWhenPlacedOnHold = val.(int32)
	}
	if val, ok := m["secondsElapsedWhenRemovedFromHold"]; ok && val != nil {
		x.SecondsElapsedWhenRemovedFromHold = val.(int32)
	}
}
