package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type gitopsContainer struct {
	path  string
	attrs []attr.Attribute
}

var Gitops = &gitopsContainer{path: "assets/apps/gitops"}

func (c *gitopsContainer) Argocd(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/gitops/argocd.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("argocd", opts...)
}

func (c *gitopsContainer) Flagger(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/gitops/flagger.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("flagger", opts...)
}

func (c *gitopsContainer) Flux(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/gitops/flux.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("flux", opts...)
}

func init() {
	const namespace = "apps.gitops"
	Register(namespace, "Argocd", Gitops.Argocd)
	Register(namespace, "Flagger", Gitops.Flagger)
	Register(namespace, "Flux", Gitops.Flux)
}