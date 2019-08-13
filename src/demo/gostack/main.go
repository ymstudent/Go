package gostack

type Item interface {

}

type ItemStack struct {
	items []Item
}

func (s *ItemStack) New() *ItemStack {
	s.items = []Item{}
	return s
}

func (s *ItemStack) Push(t Item)  {
	s.items = append(s.items, t)
}

func (s *ItemStack) Pop() *Item {
	item := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items) - 1]
	return &item
}