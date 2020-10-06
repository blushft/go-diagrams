package outscale

import attr "github.com/blushft/go-diagrams/attr"

type outscaleContainer struct {
	path  string
	attrs []attr.Attribute
}

var Outscale = &outscaleContainer{path: "assets/outscale/outscale"}
