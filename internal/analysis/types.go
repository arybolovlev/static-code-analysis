package analysis

type NodesSet struct {
	Nodes []*Node `json:"items"`
}

type Node struct {
	Name string `json:"name"`
	// FileName string              `json:"fileName"`
	Impact map[*Node]struct{} `json:"impact"`
}

func NewNodeSet() *NodesSet {
	return &NodesSet{}
}

func (s *NodesSet) InsertNode(n Node) {
	s.Nodes = append(s.Nodes, &n)
}

func (s *NodesSet) GetNode(n string) *Node {
	for _, i := range s.Nodes {
		if i.Name == n {
			return i
		}
	}

	return nil
}
