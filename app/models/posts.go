package models

type Item struct {
	ID       interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	User     string      `json:"user" bson:"user"`
	Title    string      `json:"title" bson:"title"`
	Location string      `json:"location" bson:"location"`
	Time     int64       `json:"time" bson:"time"`
	Text     string      `json:"text" bson:"text"`
	Images   interface{} `json:"images" bson:"images"`
	Likes    int64       `json:"likes" bson:"likes"`
	Comments interface{} `json:"comments" bson:"comments"`
}
