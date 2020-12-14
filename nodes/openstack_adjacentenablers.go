package nodes

import attr "github.com/blushft/go-diagrams/attr"

type adjacentenablersContainer struct {
	path  string
	attrs []attr.Attribute
}

var Adjacentenablers = &adjacentenablersContainer{path: "assets/openstack/adjacentenablers"}
