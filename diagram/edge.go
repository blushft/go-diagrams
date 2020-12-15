package diagram

import (
	"github.com/blushft/go-diagrams/attr"
)

type Edge struct {
	id    string
	start string
	end   string
	attrs attr.Attributes
}

var _ IEdge = &Edge{}

func NewEdge(id string, start, end INode, attrs ...attr.Attribute) IEdge {
	a := defaultEdgeAttrs(attrs...)

	return &Edge{
		id:    id,
		start: start.ID(),
		end:   end.ID(),
		attrs: a,
	}
}

func (e *Edge) ID() string {
	return e.id
}

func (e *Edge) Start() string {
	return e.start
}

func (e *Edge) End() string {
	return e.end
}

func (e *Edge) Attributes() (map[string]string, error) {
	return e.attrs.Render()
}
