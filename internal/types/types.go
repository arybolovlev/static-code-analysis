package types

type Package struct {
	Name    string    `json:"name"`
	Nodes   []*Node   `json:"nodes"`
	Structs []*Struct `json:"structs"`
}

type Struct struct {
	Name string `json:"name"`
}

type Node struct {
	Name     string            `json:"name"`
	Imports  map[string]string `json:"imports"`
	CalledBy map[string]*Func  `json:"calledBy"`
	Calls    map[string]*Func  `json:"calls"`
}

type Func struct {
	// Name string `json:"name"`
	PkgName    string `json:"pkgName,omitempty"`
	PkgPath    string `json:"pkgPath,omitempty"`
	StructName string `json:"structName,omitempty"`
}

func NewPackage() *Package {
	return &Package{
		Nodes: make([]*Node, 0),
	}
}

func NewNode(n string) *Node {
	return &Node{
		Name:     n,
		CalledBy: make(map[string]*Func),
		Calls:    make(map[string]*Func),
	}
}

func (p *Package) InsertNode(n Node) {
	p.Nodes = append(p.Nodes, &n)
}

func (p *Package) InsertStruct(s Struct) {
	p.Structs = append(p.Structs, &s)
}

func (p *Package) GetNode(n string) *Node {
	for _, i := range p.Nodes {
		if i.Name == n {
			return i
		}
	}

	return nil
}
