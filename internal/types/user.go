package types

type BasicUserInfo struct {
	UserID  string `json:"user_id" bson:"user_id"`
	Name    string `json:"name" bson:"name"`
	Picture string `json:"picture" bson:"picture"`
}
