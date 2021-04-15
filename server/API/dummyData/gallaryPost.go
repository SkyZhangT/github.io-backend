package dummyData

type Item struct {
	Title   string   `json:"title"`
	User    string   `json:"user"`
	Time    string   `json:"time"`
	Text    string   `json:"text"`
	Picture []string `json:"picture"`
}

type Posts struct {
	Items []Item
}

func New() *Posts {
	return &Posts{Items: []Item{}}
}

func (p *Posts) Add(i Item) {
	p.Items = append(p.Items, i)
}

func (p *Posts) Get() []Item {
	return p.Items
}