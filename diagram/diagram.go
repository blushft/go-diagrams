package diagram

import (
	"errors"
	"io/ioutil"
	"os"

	graphviz "github.com/awalterschulze/gographviz"
	"github.com/davecgh/go-spew/spew"
)

type Connector interface {
	Connect(start, end *Node, opts ...EdgeOption) Connector
	ConnectByID(start, end string, opts ...EdgeOption) Connector
}

type Diagram struct {
	options Options

	g *graphviz.Escape

	root *Group
}

func New(opts ...Option) (*Diagram, error) {
	options := DefaultOptions(opts...)
	g := graphviz.NewEscape()
	g.SetName("root")
	g.SetDir(true)

	for k, v := range options.attrs() {
		if err := g.AddAttr("root", k, v); err != nil {
			return nil, err
		}
	}

	return new(g, options), nil
}

func new(g *graphviz.Escape, options Options) *Diagram {
	return &Diagram{
		g:       g,
		options: options,
		root:    newGroup("root", 0, nil),
	}
}

func (d *Diagram) Nodes() []*Node {
	return d.root.Nodes()
}

func (d *Diagram) Edges() []*Edge {
	return d.root.Edges()
}

func (d *Diagram) Groups() []*Group {
	return d.root.Children()
}

func (d *Diagram) Add(ns ...*Node) *Diagram {
	d.root.Add(ns...)
	return d
}

func (d *Diagram) Connect(start, end *Node, opts ...EdgeOption) *Diagram {
	d.Add(start, end)
	return d.ConnectByID(start.ID(), end.ID(), opts...)
}

func (d *Diagram) ConnectByID(start, end string, opts ...EdgeOption) *Diagram {
	d.root.ConnectByID(start, end, opts...)

	return d
}

func (d *Diagram) Group(g *Group) *Diagram {
	d.root.Group(g)
	return d
}

func (d *Diagram) Close() error {
	return nil
}

func (d *Diagram) Render() error {
	return d.render()
}

func (d *Diagram) render() error {
	for _, n := range d.root.nodes {
		err := n.render("root", d.g)
		if err != nil {
			return err
		}
	}

	for _, e := range d.root.edges {
		err := e.render(e.Start(), e.End(), d.g)
		if err != nil {
			return err
		}
	}

	for _, g := range d.root.children {
		if err := g.render(d.g); err != nil {
			return err
		}
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
	spew.Dump(d.g.Relations)
	return ioutil.WriteFile(fname, []byte(d.g.String()), os.ModePerm)
}
