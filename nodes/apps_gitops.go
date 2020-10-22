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
	return node.New("argocd", attr.AssetImage("assets/apps/gitops/argocd.png"), attr.Shape(attr.None))
}

func (c *gitopsContainer) Flagger(opts ...attr.Attribute) *node.Node {
	return node.New("flagger", attr.AssetImage("assets/apps/gitops/flagger.png"), attr.Shape(attr.None))
}

func (c *gitopsContainer) Flux(opts ...attr.Attribute) *node.Node {
	return node.New("flux", attr.AssetImage("assets/apps/gitops/flux.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.gitops"
	Register(namespace, "Argocd", Gitops.Argocd)
	Register(namespace, "Flagger", Gitops.Flagger)
	Register(namespace, "Flux", Gitops.Flux)
}