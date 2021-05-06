package node

import (
	"github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/attr/color"
)

type Nodes []*Node

type Node struct {
	id    string
	attrs attr.Attributes

	edgesFrom []string
	edgesTo   []string
}

func New(id string, attrs ...attr.Attribute) *Node {
	return &Node{
		id:    id,
		attrs: attr.NewAttributes(attrs...),
	}
}

func (n *Node) ID() string {
	return n.id
}

func (n *Node) Attributes() attr.Attributes {
	return n.attrs
}

func (n *Node) WithAttribute(attr attr.Attribute) *Node {
	n.attrs.Set(attr)

	return n
}

func (n *Node) WithAttributes(attrs ...attr.Attribute) *Node {
	n.attrs.Set(attrs...)

	return n
}

func (n *Node) URL(u string) *Node {
	return n.WithAttribute(attr.URL(u))
}

func (n *Node) Area(area float64) *Node {
	return n.WithAttribute(attr.Area(area))
}

func (n *Node) Class(c string) *Node {
	return n.WithAttribute(attr.Class(c))
}

func (n *Node) Color(c color.Color) *Node {
	return n.WithAttribute(attr.Color(c))
}

func (n *Node) ColorScheme(c string) *Node {
	return n.WithAttribute(attr.ColorScheme(c))
}

func (n *Node) Comment(c string) *Node {
	return n.WithAttribute(attr.Comment(c))
}

func (n *Node) Distortion(v float64) *Node {
	return n.WithAttribute(attr.Distortion(v))
}

func (n *Node) FillColor(c color.Color) *Node {
	return n.WithAttribute(attr.FillColor(c))
}

func (n *Node) FontSize(s float64) *Node {
	return n.WithAttribute(attr.FontSize(s))
}

func (n *Node) Image(img string) *Node {
	return n.WithAttribute(attr.Image(img))
}
