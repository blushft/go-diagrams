package edge

import (
	"github.com/blushft/go-diagrams/attr"
)

func defaultAttributes(attrs ...attr.Attribute) attr.Attributes {
	return attr.NewAttributes(
		attr.FontName("Sans Serif"),
		attr.FontSize(13),
		attr.Direction(attr.DirEdgeNone),
		attr.WithAttributes(attrs...),
	)
}
