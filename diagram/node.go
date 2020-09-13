package diagram

import (
	"strconv"
	"strings"

	graphviz "github.com/awalterschulze/gographviz"
	"github.com/blushft/go-diagrams/pkg/rand"
)

type Node struct {
	id string

	conn    Connector
	Options NodeOptions
}

func NewNode(opts ...NodeOption) *Node {
	options := DefaultNodeOptions(opts...)

	return &Node{
		id:      rand.String(8),
		Options: options,
	}
}

func (n *Node) ID() string {
	return n.id
}

func (n *Node) Label(l string) *Node {
	n.Options.Label = l
	return n
}

func (n *Node) setConnector(c Connector) {
	n.conn = c
}

func (n *Node) render(name string, graph *graphviz.Escape) error {
	return graph.AddNode(name, n.id, n.attrs())
}

func (n *Node) attrs() map[string]string {

	if n.Options.Image != "" {
		labelH := float64(len(strings.Split(n.Options.Label, "/n")))
		n.Options.Height = n.Options.Height + (0.4 * labelH)
		n.Options.Shape = "none"
	}

	attrs := map[string]string{
		"label":      n.Options.Label,
		"image":      n.Options.Image,
		"height":     strconv.FormatFloat(n.Options.Height, 'f', -1, 64),
		"width":      strconv.FormatFloat(n.Options.Width, 'f', -1, 64),
		"shape":      n.Options.Shape,
		"style":      n.Options.Style,
		"labelloc":   n.Options.LabelLocation,
		"imagescale": strconv.FormatBool(n.Options.ImageScale),
		"fixedsize":  strconv.FormatBool(n.Options.FixedSize),
		"fontsize":   strconv.FormatInt(int64(n.Options.Font.Size), 10),
		"fontname":   n.Options.Font.Name,
		"fontcolor":  n.Options.Font.Color,
	}

	for k, v := range n.Options.Attributes {
		attrs[k] = v
	}

	return trimAttrs(attrs)
}

func renderNodes(name string, graph *graphviz.Escape, nodes ...*Node) error {
	for _, n := range nodes {
		err := n.render(name, graph)
		if err != nil {
			return err
		}
	}

	return nil
}

type NodeOptions struct {
	Name          string
	Label         string
	Provider      string
	Image         string
	ImageScale    bool
	Shape         string
	Style         string
	FixedSize     bool
	Width         float64
	Height        float64
	LabelLocation string
	Font          Font
	Attributes    map[string]string
}

type NodeOption func(*NodeOptions)

type OptionSet []NodeOption

func DefaultNodeOptions(opts ...NodeOption) NodeOptions {
	nopts := NodeOptions{
		Name:          "node",
		Label:         "node",
		Shape:         "box",
		Style:         "rounded",
		FixedSize:     true,
		Width:         1.4,
		Height:        1.4,
		LabelLocation: "b",
		ImageScale:    true,
		Font: Font{
			Name:  "Sans-Serif",
			Size:  13,
			Color: "#@D3436",
		},
		Attributes: make(map[string]string),
	}

	for _, o := range opts {
		o(&nopts)
	}

	return nopts
}

func MergeOptionSets(sets ...OptionSet) OptionSet {
	ns := OptionSet{}

	for _, set := range sets {
		for _, opt := range set {
			ns = append(ns, opt)
		}
	}

	return ns
}

func Name(n string) NodeOption {
	return func(o *NodeOptions) {
		o.Name = n
	}
}

func NodeLabel(l string) NodeOption {
	return func(o *NodeOptions) {
		o.Label = l
	}
}

func Provider(p string) NodeOption {
	return func(o *NodeOptions) {
		o.Provider = p
	}
}

func Icon(i string) NodeOption {
	return func(o *NodeOptions) {
		o.Image = i
	}
}

func NodeShape(s string) NodeOption {
	return func(o *NodeOptions) {
		o.Shape = s
	}
}

func NodeStyle(s string) NodeOption {
	return func(o *NodeOptions) {
		o.Style = s
	}
}

func FixedSize(b bool) NodeOption {
	return func(o *NodeOptions) {
		o.FixedSize = b
	}
}

func Width(w float64) NodeOption {
	return func(o *NodeOptions) {
		o.Width = w
	}
}

func Height(h float64) NodeOption {
	return func(o *NodeOptions) {
		o.Height = h
	}
}

func LabelLocation(l string) NodeOption {
	return func(o *NodeOptions) {
		o.LabelLocation = l
	}
}

func ImageScale(b bool) NodeOption {
	return func(o *NodeOptions) {
		o.ImageScale = b
	}
}

func SetFontOptions(f Font) NodeOption {
	return func(o *NodeOptions) {
		o.Font = f
	}
}
