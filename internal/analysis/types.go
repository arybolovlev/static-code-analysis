package analysis

type Module struct {
	Path     string     `json:"path"`
	Packages []*Package `json:"packages"`
}

type Package struct {
	Name      string              `json:"name"`
	Imports   map[string]string   `json:"imports,omitempty"`
	Types     map[string]struct{} `json:"types,omitempty"`
	Functions []*Function         `json:"functions,omitempty"`
}

type Function struct {
	Name     string              `json:"name"`
	CalledBy map[string]struct{} `json:"calledBy,omitempty"`
	Calls    map[string]struct{} `json:"calls,omitempty"`
}

// func NewPackage() *Package {
// 	return &Packages{}
// }

// func (ns *Nodes) InsertNode(n Node) {
// 	ns.Nodes = append(ns.Nodes, &n)
// }

func (p *Package) GetFunction(fName string) *Function {
	for _, fn := range p.Functions {
		if fn.Name == fName {
			return fn
		}
	}

	return nil
}
