package analysis

type ItemSet struct {
	Items []*Item `json:"items"`
}

type Item struct {
	Name     string              `json:"name"`
	FileName string              `json:"fileName"`
	Impact   map[string]struct{} `json:"impact"`
}

func NewItemSet() *ItemSet {
	return &ItemSet{}
}

func (s *ItemSet) InsertItem(i Item) {
	s.Items = append(s.Items, &i)
}

func (s *ItemSet) Item(n string) *Item {
	for _, i := range s.Items {
		if i.Name == n {
			return i
		}
	}

	return nil
}
