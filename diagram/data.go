package diagram

import "github.com/blushft/go-diagrams/attr"

type Meta struct {
	Name            string   `yaml:"name"`
	Label           string   `yaml:"label"`
	Type            string   `yaml:"type"`
	ConnectTo       []string `yaml:"connectTo"`
	ConnectFrom     []string `yaml:"connectFrom"`
	ConnectAllTo    []string `yaml:"connectAllTo"`
	ConnectAllFrom  []string `yaml:"connectAllFrom"`
	BackgroundColor string   `yaml:"backgroundColor"`
}

type EdgeData struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
}

type GroupData struct {
	Meta `yaml:",inline"`

	Group []GroupData `yaml:"groups"`
	Nodes []NodeData  `yaml:"nodes"`
	Edges []EdgeData  `yaml:"edges"`
}

type NodeData struct {
	Meta `yaml:",inline"`
}

type DiagramData struct {
	Direction attr.EdgeDirection `yaml:"direction" default:"forward"`
	Label     string             `yaml:"label"`
	Groups    []GroupData        `yaml:"groups"`
	Nodes     []NodeData         `yaml:"nodes"`
}
