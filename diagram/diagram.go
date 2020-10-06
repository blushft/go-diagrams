package diagram

import (
	"strings"

	"github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/edge"
	"github.com/blushft/go-diagrams/node"
	"github.com/google/uuid"
)

type GroupType string

const (
	RootGraph GroupType = "G"
	Subgraph  GroupType = "S"
	Cluster   GroupType = "C"
)

func New(id string, attrs ...attr.Attribute) *Group {
	return new(id, RootGraph, defaultRootAttrs(attrs...))
}

type Group struct {
	t        GroupType
	id       string
	attrs    attr.Attributes
	parent   *Group
	children map[string]*Group

	nodes map[string]*node.Node
	edges map[string]*edge.Edge
}

func NewGroup(id string, attrs ...attr.Attribute) *Group {
	return new(id, Subgraph, defaultGroupAttrs(attrs...))
}

func NewCluster(id string, attrs ...attr.Attribute) *Group {
	return new("cluster"+id, Cluster, defaultClusterAttrs(attrs...))
}

func new(id string, t GroupType, attrs attr.Attributes) *Group {
	return &Group{
		t:        t,
		id:       id,
		attrs:    attrs,
		children: make(map[string]*Group),
		nodes:    make(map[string]*node.Node),
		edges:    make(map[string]*edge.Edge),
	}
}

func (g *Group) ID() string {
	return g.id
}

func (g *Group) Attributes() (map[string]string, error) {
	return g.attrs.Render()
}

func (g *Group) Nodes() []Node {
	ns := make([]Node, 0, len(g.nodes))

	for _, n := range g.nodes {
		ns = append(ns, n)
	}

	return ns
}

func (g *Group) Edges() []Edge {
	es := make([]Edge, 0, len(g.edges))

	for _, e := range g.edges {
		es = append(es, e)
	}

	return es
}

func (g *Group) Children() []Diagram {
	cs := make([]Diagram, 0, len(g.children))

	for _, c := range g.children {
		cs = append(cs, c)
	}

	return cs
}

func (g *Group) AddNode(n *node.Node) *Group {
	g.nodes[n.ID()] = n
	return g
}

func (g *Group) AddNodes(ns ...*node.Node) *Group {
	for _, n := range ns {
		g.AddNode(n)
	}

	return g
}

func (g *Group) Connect(start, end *node.Node) *Group {
	g.AddNodes(start, end)
	eid := strings.SplitAfter(strings.Replace(uuid.New().String(), "-", "", 1), "-")[0]
	g.edges[eid] = edge.New(eid, start, end)

	return g
}
