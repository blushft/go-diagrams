package diagram

import (
	"strconv"

	graphviz "github.com/awalterschulze/gographviz"
)

type Group struct {
	idx      int
	id       string
	options  GroupOptions
	parent   *Group
	children map[string]*Group

	nodes map[string]*Node
	edges map[string]*Edge
}

func NewGroup(name string, opts ...GroupOption) *Group {
	return newGroup("cluster_"+name, 0, nil, opts...)

}

func newGroup(name string, idx int, parent *Group, opts ...GroupOption) *Group {
	options := defaultGroupOptions(idx, opts...)

	return &Group{
		id:       name,
		idx:      idx,
		options:  options,
		parent:   parent,
		children: make(map[string]*Group),
		nodes:    make(map[string]*Node),
		edges:    make(map[string]*Edge),
	}
}

func (g *Group) ID() string {
	return g.id
}

func (g *Group) Nodes() []*Node {
	nodes := make([]*Node, 0, len(g.nodes))
	for _, n := range g.nodes {
		nodes = append(nodes, n)
	}

	return nodes
}

func (g *Group) Edges() []*Edge {
	edges := make([]*Edge, 0, len(g.edges))
	for _, e := range g.edges {
		edges = append(edges, e)
	}

	return edges
}

func (g *Group) Children() []*Group {
	gs := make([]*Group, 0, len(g.children))

	for _, c := range g.children {
		gs = append(gs, c)
	}

	return gs
}

func (g *Group) Add(nodes ...*Node) *Group {
	for _, n := range nodes {
		g.nodes[n.ID()] = n
	}

	return g
}

func (g *Group) Connect(start, end *Node, opts ...EdgeOption) *Group {
	g.Add(start, end)
	return g.ConnectByID(start.ID(), end.ID(), opts...)
}

func (g *Group) ConnectByID(start, end string, opts ...EdgeOption) *Group {
	e := NewEdge(start, end, opts...)
	g.edges[e.ID()] = e

	return g
}

func (g *Group) ConnectAllTo(end string, opts ...EdgeOption) *Group {
	for _, n := range g.nodes {
		g.ConnectByID(n.ID(), end, opts...)
	}

	return g
}

func (g *Group) ConnectAllFrom(start string, opts ...EdgeOption) *Group {
	for _, n := range g.nodes {
		g.ConnectByID(start, n.ID(), opts...)
	}

	return g
}

func (g *Group) attrs() map[string]string {
	attrs := map[string]string{
		"label":     g.options.Label,
		"labeljust": g.options.LabelJustify,
		"pencolor":  g.options.PenColor,
		"bgcolor":   g.options.BackgroundColor,
		"shape":     g.options.Shape,
		"style":     g.options.Style,
		"fontname":  g.options.Font.Name,
		"fontsize":  strconv.FormatInt(int64(g.options.Font.Size), 10),
		"fontcolor": g.options.Font.Color,
	}

	for k, v := range g.options.Attributes {
		attrs[k] = v
	}

	return trimAttrs(attrs)
}

func (g *Group) Group(ng *Group) *Group {
	g.children[ng.id] = ng
	ng.parent = g
	return ng
}

func (g *Group) NewGroup(name string, opts ...GroupOption) *Group {
	idx := g.idx + 1

	ng := newGroup("cluster"+name, idx, g, opts...)
	g.children[ng.id] = ng
	ng.parent = g

	return ng
}

func (g *Group) Label(l string) *Group {
	g.options.Label = l
	return g
}

func (g *Group) BackgroundColor(c string) *Group {
	g.options.BackgroundColor = c
	return g
}

func (g *Group) render(outdir string, graph *graphviz.Escape) error {
	if err := graph.AddSubGraph(g.parent.id, g.id, g.attrs()); err != nil {
		return err
	}

	for _, n := range g.nodes {
		if err := n.render(g.id, outdir, graph); err != nil {
			return err
		}
	}

	for _, e := range g.edges {
		if err := e.render(e.start, e.end, graph); err != nil {
			return err
		}
	}

	for _, child := range g.children {
		if err := child.render(outdir, graph); err != nil {
			return err
		}
	}

	return nil
}

type GroupOptions struct {
	Label           string
	LabelJustify    string
	Direction       string
	PenColor        string
	BackgroundColor string
	Shape           string
	Style           string
	Font            Font
	Attributes      map[string]string
}

func DefaultGroupOptions(opts ...GroupOption) GroupOptions {
	return defaultGroupOptions(0, opts...)
}

func defaultGroupOptions(idx int, opts ...GroupOption) GroupOptions {
	options := GroupOptions{
		LabelJustify: "l",
		Direction:    string(LeftToRight),
		PenColor:     "#AEB6BE",
		Shape:        "box",
		Style:        "rounded",
		Font: Font{
			Name:  "Sans-Serif",
			Size:  12,
			Color: "#2D3436",
		},
		Attributes: make(map[string]string),
	}

	IndexedBackground(idx)(&options)

	for _, o := range opts {
		o(&options)
	}

	return options
}

type GroupOption func(*GroupOptions)

func BackgroundColor(c string) GroupOption {
	return func(o *GroupOptions) {
		o.BackgroundColor = c
	}
}

func IndexedBackground(idx int) GroupOption {
	bgcs := []string{"#E5F5FD", "#EBF3E7", "#ECE8F6", "#FDF7E3"}
	if idx-1 > len(bgcs) {
		idx = 0
	}

	return BackgroundColor(bgcs[idx])
}

func GroupLabel(l string) GroupOption {
	return func(o *GroupOptions) {
		o.Label = l
	}
}
