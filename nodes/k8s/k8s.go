package k8s

import attr "github.com/blushft/go-diagrams/attr"

type k8sContainer struct {
	path  string
	attrs []attr.Attribute
}

var K8S = &k8sContainer{path: "assets/k8s/k8s"}
