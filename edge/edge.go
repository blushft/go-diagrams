package edge

import (
	"github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type Edge struct {
	id    string
	start string
	end   string
	attrs attr.Attributes
}

func New(id string, start, end *node.Node, attrs ...attr.Attribute) *Edge {
	a := attr.NewAttributes(attrs...)

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
