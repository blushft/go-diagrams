package node

import (
	"github.com/blushft/go-diagrams/attr"
)

func DefaultAttributes(attrs ...attr.Attribute) []attr.Attribute {
	var newAttrs []attr.Attribute

	newAttrs = append(newAttrs,
		attr.Shape(attr.Box),
		attr.Style(attr.Rounded),
		attr.FixedSize(true),
		attr.Width(1.4),
		attr.Height(1.7),
		attr.LabelLocation("b"),
		attr.LabelDistance(5.0),
		attr.ImageScale("both"),
		attr.FontName("Sans Serif"),
		attr.FontSize(12),
		attr.WithAttributes(attrs...),
	)
	return newAttrs
}
