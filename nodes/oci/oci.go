package oci

import attr "github.com/blushft/go-diagrams/attr"

type ociContainer struct {
	path  string
	attrs []attr.Attribute
}

var Oci = &ociContainer{path: "assets/oci/oci"}
