package diagram

import (
	"errors"
	graphviz "github.com/awalterschulze/gographviz"
	"io/ioutil"
	"os"
	"path/filepath"
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

func (d *Diagram) Parse(file *os.File) error {
	return d.parse(file)
}

func (d *Diagram) Render() error {
	return d.render()
}

func (d *Diagram) render() error {
	outdir := d.options.Name
	if err := os.MkdirAll(outdir, os.ModePerm); err != nil {
		return err
	}

	for _, n := range d.root.nodes {
		err := n.render("root", outdir, d.g)
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
		if err := g.render(outdir, d.g); err != nil {
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
	fname := filepath.Join(d.options.Name, d.options.FileName+".dot")
	return ioutil.WriteFile(fname, []byte(d.g.String()), 0644)
}

func (d *Diagram) interpretYML(yml *DiagramData) error {
	if yml.Label != "" {
		d.options.Label = yml.Label
	}

	if yml.Direction != "" {
		d.options.Direction = yml.Direction
	}

	for _, v := range yml.Nodes {
		parsedNode := d.parseNode(v)
		switch parsedNode.(type) {
		case *Node:
			d.Add(parsedNode.(*Node))
		case *Group:
			d.Group(parsedNode.(*Group))
		}
	}

	return nil
}

func (d *Diagram) parseNode(v NodeData) interface{} {
	switch v.Type {
	case "Group":
		groupNode := NewGroup(v.Name)
		if v.Label != "" {
			groupNode.Label(v.Label)
		}

		if v.BackgroundColor != "" {
			groupNode.BackgroundColor(v.BackgroundColor)
		}

		for _, child := range v.Nodes {
			parsedNode := d.parseNode(child)
			switch parsedNode.(type) {
			case *Node:
				groupNode.Add(parsedNode.(*Node))
			case *Group:
				groupNode.Group(parsedNode.(*Group))
			}
		}
		return groupNode
	/*case "gcp.Network.DNS":
		// TODO: find a better way to do the mapping
		node := gcp.Network.Dns()
		if v.Label != "" {
			node.Label(v.Label)
		}
		return node*/
	default:
		panic("unimplemented " + v.Type)
	}

	return nil
}