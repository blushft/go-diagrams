package diagram

import (
	"strconv"

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

func (e *Edge) render(start, end string, graph *graphviz.Escape) error {
	return graph.AddEdge(start, end, true, e.attrs())
}

func (e *Edge) attrs() map[string]string {
	attrs := map[string]string{
		"label":     e.Options.Label,
		"color":     e.Options.Color,
		"dir":       e.Options.edgeDirection(),
		"style":     e.Options.Style,
		"fontname":  e.Options.Font.Name,
		"fontcolor": e.Options.Font.Color,
		"fontsize":  strconv.FormatInt(int64(e.Options.Font.Size), 10),
	}

	for k, v := range e.Options.Attributes {
		attrs[k] = v
	}

	return trimAttrs(attrs)
}

type EdgeOptions struct {
	Label      string
	Color      string
	Forward    bool
	Reverse    bool
	Font       Font
	Style      string
	Attributes map[string]string
}

type EdgeOption func(*EdgeOptions)

func DefaultEdgeOptions(opts ...EdgeOption) EdgeOptions {
	eopts := EdgeOptions{
		Color: "#7B8894",
		Font: Font{
			Name:  "Sans-Serif",
			Size:  13,
			Color: "#2D3436",
		},
		Forward:    true,
		Attributes: make(map[string]string),
	}

	for _, o := range opts {
		o(&eopts)
	}

	return eopts
}

func (o EdgeOptions) edgeDirection() string {
	if o.Forward && o.Reverse {
		return "both"
	}

	if o.Forward {
		return "forward"
	}

	if o.Reverse {
		return "back"
	}

	return "none"
}

func Forward() EdgeOption {
	return func(o *EdgeOptions) {
		o.Forward = true
		o.Reverse = false
	}
}

func Reverse() EdgeOption {
	return func(o *EdgeOptions) {
		o.Reverse = true
		o.Forward = false
	}
}

func Bidirectional() EdgeOption {
	return func(o *EdgeOptions) {
		o.Forward = true
		o.Reverse = true
	}
}
