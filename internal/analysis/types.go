package analysis

type Nodes struct {
	Nodes []*Node `json:"nodes"`
}

type Node struct {
	Name     string              `json:"name"`
	CalledBy map[string]struct{} `json:"calledBy"`
	Calls    map[string]struct{} `json:"calls"`
}

func NewNode() *Nodes {
	return &Nodes{}
}

func (ns *Nodes) InsertNode(n Node) {
	ns.Nodes = append(ns.Nodes, &n)
}

func (ns *Nodes) GetNode(n string) *Node {
	for _, i := range ns.Nodes {
		if i.Name == n {
			return i
		}
	}

	return nil
}
