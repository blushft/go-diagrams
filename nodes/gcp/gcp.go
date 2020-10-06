package gcp

import attr "github.com/blushft/go-diagrams/attr"

type gcpContainer struct {
	path  string
	attrs []attr.Attribute
}

var Gcp = &gcpContainer{path: "assets/gcp/gcp"}
