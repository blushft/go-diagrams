package nodes

import attr "github.com/blushft/go-diagrams/attr"

type adjacentenablersContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackAdjacentenablers =&adjacentenablersContainer{path: "assets/openstack/adjacentenablers"}

func init() {
  const namespace = "openstack.adjacentenablers"
}
