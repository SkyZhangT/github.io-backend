package models

type Item struct {
	User    string      `json:"user" bson:"user"`
	Title   string      `json:"title" bson:"title"`
	Time    int32       `json:"time" bson:"time"`
	Text    string      `json:"text" bson:"text"`
	Picture interface{} `json:"picture" bson:"picture"`
}