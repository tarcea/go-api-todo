package data

type Item struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	CreatedAt   string `json:"createdAt"`
}

type List struct {
	Items []Item
}

func New() *List {
	return &List{}
}

func (r *List) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *List) GetAll() []Item {
	return r.Items
}
