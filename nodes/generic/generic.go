package generic

import attr "github.com/blushft/go-diagrams/attr"

type genericContainer struct {
	path  string
	attrs []attr.Attribute
}

var Generic = &genericContainer{path: "assets/generic/generic"}
