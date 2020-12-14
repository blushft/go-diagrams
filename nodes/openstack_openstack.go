package nodes

import attr "github.com/blushft/go-diagrams/attr"

type openstackContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackOpenstack =&openstackContainer{path: "assets/openstack/openstack"}

func init() {
  const namespace = "openstack.openstack"
}
