package programming

import attr "github.com/blushft/go-diagrams/attr"

type programmingContainer struct {
	path  string
	attrs []attr.Attribute
}

var Programming = &programmingContainer{path: "assets/programming/programming"}
