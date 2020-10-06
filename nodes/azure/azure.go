package azure

import attr "github.com/blushft/go-diagrams/attr"

type azureContainer struct {
	path  string
	attrs []attr.Attribute
}

var Azure = &azureContainer{path: "assets/azure/azure"}
