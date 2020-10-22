package nodes

import attr "github.com/blushft/go-diagrams/attr"

type appsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Apps = &appsContainer{path: "assets/apps/apps"}