package aws

import attr "github.com/blushft/go-diagrams/attr"

type awsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Aws = &awsContainer{path: "assets/aws/aws"}
