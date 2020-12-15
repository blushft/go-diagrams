package diagram

import (
	"github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/attr/color"
)

func defaultRootAttrs(attrs ...attr.Attribute) attr.Attributes {
	return attr.NewAttributes(
		attr.Pad(2.0),
		attr.Splines(attr.SplineOrtho),
		attr.RankDir(attr.LeftToRight),
		attr.NodeSeparation(0.8),
		attr.RankSeparation(1.2),
		attr.FontName("Sans Serif"),
		attr.FontColor(color.Gray53()),
		attr.FontSize(16),
		attr.WithAttributes(attrs...),
	)
}

func defaultGroupAttrs(attrs ...attr.Attribute) attr.Attributes {
	return attr.NewAttributes(
		attr.Pad(2.0),
		attr.Splines(attr.SplineOrtho),
		attr.RankDir(attr.LeftToRight),
		attr.NodeSeparation(0.8),
		attr.RankSeparation(1.2),
		attr.FontName("Sans Serif"),
		attr.FontSize(12),
		attr.WithAttributes(attrs...),
	)
}

func defaultClusterAttrs(attrs ...attr.Attribute) attr.Attributes {
	return attr.NewAttributes(
		attr.RankDir(attr.LeftToRight),
		attr.Shape(attr.Box),
		attr.Style(attr.Rounded),
		attr.FontName("Sans Serif"),
		attr.FontSize(12),
		attr.WithAttributes(attrs...),
	)
}


func defaultEdgeAttrs(attrs ...attr.Attribute) attr.Attributes {
	return attr.NewAttributes(
		attr.FontName("Sans Serif"),
		attr.FontSize(12),
		attr.Pad(10.0),
		attr.LabelDistance(4.0),
		attr.LabelJustify(attr.JustifyCenter),
		attr.LabelHeight(5.0),
		attr.Direction(attr.EdgeNone),
		attr.WithAttributes(attrs...),
	)
}