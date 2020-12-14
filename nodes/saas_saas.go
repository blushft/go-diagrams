package nodes

import attr "github.com/blushft/go-diagrams/attr"

type saasContainer struct {
	path  string
	attrs []attr.Attribute
}

var Saas = &saasContainer{path: "assets/saas/saas"}
