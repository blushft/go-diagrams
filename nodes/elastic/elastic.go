package elastic

import attr "github.com/blushft/go-diagrams/attr"

type elasticContainer struct {
	path  string
	attrs []attr.Attribute
}

var Elastic = &elasticContainer{path: "assets/elastic/elastic"}
