package nodes

import attr "github.com/blushft/go-diagrams/attr"

type openstackOperationsContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackOperations = &openstackOperationsContainer{path: "assets/openstack/operations"}

func init() {
  const namespace = "openstack.operations"
}
