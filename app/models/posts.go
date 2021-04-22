package models

type Item struct {
	ID     interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	User   string      `json:"user" bson:"user"`
	Title  string      `json:"title" bson:"title"`
	Time   int64       `json:"time" bson:"time"`
	Text   string      `json:"text" bson:"text"`
	Images interface{} `json:"images" bson:"images"`
}