package nodes

import attr "github.com/blushft/go-diagrams/attr"

type lifecyclemanagementContainer struct {
	path  string
	attrs []attr.Attribute
}

var OpenstackLifecyclemanagement =&lifecyclemanagementContainer{path: "assets/openstack/lifecyclemanagement"}

func init() {
  const namespace = "openstack.lifecyclemanagement"
}
