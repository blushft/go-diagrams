package node

import (
	graphviz "github.com/awalterschulze/gographviz"
	"github.com/google/uuid"
)

type Edge struct {
	id      string
	start   string
	end     string
	Options EdgeOptions
}

func NewEdge(start, end string, opts ...EdgeOption) *Edge {
	options := DefaultEdgeOptions(opts...)

	return &Edge{
		id:      uuid.New().String(),
		start:   start,
		end:     end,
		Options: options,
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

func (e *Edge) Graph(graph *graphviz.Escape, start, end string) error {
	return graph.AddEdge(start, end, true, nil)

}

type EdgeOptions struct {
	Label string
	Color string
}

type EdgeOption func(*EdgeOptions)

func DefaultEdgeOptions(opts ...EdgeOption) EdgeOptions {
	eopts := EdgeOptions{
		Color: "#7B8894",
	}

	for _, o := range opts {
		o(&eopts)
	}

	return eopts
}
