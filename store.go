package sample

type Item struct {
	A int
}

type Store struct {
	i *Item
}

func (s *Store) Sell() int {
	return s.i.A
}

func NewStore(i *Item) (*Store, error) {
	return &Store{i: i}, nil
}

func NewOKItem() *Item {
	return &Item{A: 1}
}

func NewNGItem() (*Item, error) {
	return &Item{A: 0}, &e{}
}

type e struct{}

func (e *e) Error() string {
	return "NoItem"
}
