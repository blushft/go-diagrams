package node

import (
	"github.com/blushft/go-diagrams/attr"
)

func defaultAttributes(attrs ...attr.Attribute) attr.Attributes {
	return attr.NewAttributes(
		attr.Shape(attr.Box),
		attr.Style(attr.Rounded),
		attr.FixedSize(true),
		attr.Width(1.4),
		attr.Height(1.4),
		attr.LabelLocation("b"),
		attr.ImageScale(true),
		attr.FontName("Sans Serif"),
		attr.FontSize(13),
		attr.WithAttributes(attrs...),
	)
}
