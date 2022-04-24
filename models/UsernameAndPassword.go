package models

type UsernameAndPassword struct {
	Password string `bson:"password" json:"password"`
	Username string `bson:"username" json:"username"`
}

// BUILDER from bson map:
func BuildUsernameAndPassword(m map[string]interface{}, x *UsernameAndPassword) {
	if val, ok := m["password"]; ok && val != nil {
		x.Password = val.(string)
	}
	if val, ok := m["username"]; ok && val != nil {
		x.Username = val.(string)
	}
}
