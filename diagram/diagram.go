package diagram

import "github.com/blushft/go-diagrams/attr"

type Identifiable interface {
	ID() string
}

type Attributed interface {
	Attributes() attr.Attributes
}
type Diagram interface {
	Identifiable
	Attributed
	Nodes() []Node
	Edges() []Edge
	Children() []Diagram
}

type Node interface {
	Identifiable
	Attributed
}

type Edge interface {
	Identifiable
	Attributed
	Start() string
	End() string
}
