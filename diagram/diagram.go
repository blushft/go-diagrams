package diagram

import (
	"errors"
	"io/ioutil"
	"os"

	graphviz "github.com/awalterschulze/gographviz"
	"github.com/blushft/go-diagrams/node"
)

type Diagram struct {
	idx     int
	options Options
	parent  *Diagram
	child   *Diagram

	g *graphviz.Escape

	nodes map[string]*node.Node
	edges map[string]*node.Edge
}

func New(opts ...Option) (*Diagram, error) {
	options := DefaultOptions(opts...)
	g := graphviz.NewEscape()
	g.SetName(options.Name)
	g.SetDir(true)

	for k, v := range options.attrs() {
		if err := g.AddAttr(options.Name, k, v); err != nil {
			return nil, err
		}
	}

	return new(0, g, options, nil), nil
}

func new(idx int, g *graphviz.Escape, options Options, parent *Diagram) *Diagram {
	return &Diagram{
		idx:     idx,
		g:       g,
		options: options,
		nodes:   make(map[string]*node.Node),
		edges:   make(map[string]*node.Edge),
	}
}

func (d *Diagram) Nodes() []*node.Node {
	ns := make([]*node.Node, 0, len(d.nodes))
	for _, n := range d.nodes {
		ns = append(ns, n)
	}

	return ns
}

func (d *Diagram) Edges() []*node.Edge {
	es := make([]*node.Edge, 0, len(d.edges))
	for _, e := range d.edges {
		es = append(es, e)
	}

	return es
}

func (d *Diagram) Root() *Diagram {
	if d.parent != nil {
		return d.parent.Root()
	}

	return d
}

func (d *Diagram) Next() *Diagram {
	return d.child
}

func (d *Diagram) Previous() *Diagram {
	return d.parent
}

func (d *Diagram) AddNode(n *node.Node) *Diagram {
	d.nodes[n.ID()] = n

	return d
}

func (d *Diagram) AddNodes(ns ...*node.Node) *Diagram {
	for _, n := range ns {
		d.nodes[n.ID()] = n
	}

	return d
}

func (d *Diagram) Connect(start, end *node.Node, opts ...node.EdgeOption) *Diagram {
	d.AddNodes(start, end)
	return d.ConnectByID(start.ID(), end.ID(), opts...)
}

func (d *Diagram) ConnectByID(start, end string, opts ...node.EdgeOption) *Diagram {
	e := node.NewEdge(start, end, opts...)
	d.edges[e.ID()] = e

	return d
}

func (d *Diagram) Group(name string, opts ...Option) *Diagram {
	idx := d.idx + 1
	options := groupOptions("cluster_"+name, idx, d.options, opts...)

	gr := new(idx, d.g, options, d)
	gr.parent = d
	d.child = gr

	return gr
}

func (d *Diagram) Close() error {
	return nil
}

func (d *Diagram) Render() error {
	return render(d)
}

func (d *Diagram) render() error {
	if d.parent != nil {
		if err := d.g.AddSubGraph(d.parent.options.Name, d.options.Name, d.options.attrs()); err != nil {
			return err
		}
	}

	for _, n := range d.nodes {
		err := n.Graph(d.options.Name, d.g)
		if err != nil {
			return err
		}
	}

	for _, e := range d.edges {
		err := e.Graph(d.g, e.Start(), e.End())
		if err != nil {
			return err
		}
	}

	if d.parent != nil {
		return d.parent.render()
	}

	return d.renderOutput()
}

func (d *Diagram) renderOutput() error {
	switch d.options.OutFormat {
	case "dot":
		return d.saveDot()
	default:
		return errors.New("invalid output format")
	}
}

func (d *Diagram) saveDot() error {
	fname := d.options.FileName + ".dot"
	return ioutil.WriteFile(fname, []byte(d.g.String()), os.ModePerm)
}

func render(d *Diagram) error {
	lastChild := d
	for lastChild.Next() != nil {
		lastChild = lastChild.Next()
	}

	return lastChild.render()
}
