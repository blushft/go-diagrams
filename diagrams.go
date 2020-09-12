package diagrams

import (
	"github.com/blushft/go-diagrams/diagram"
)

func New(opts ...diagram.Option) (*diagram.Diagram, error) {
	return diagram.New(opts...)
}
