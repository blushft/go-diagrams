package onprem

import "github.com/blushft/go-diagrams/node"

type gitopsContainer struct {
	path string
	opts []node.Option
}

var Gitops = &gitopsContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/gitops",
}

func (c *gitopsContainer) Flux(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/gitops/flux.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *gitopsContainer) Argocd(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/gitops/argocd.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *gitopsContainer) Flagger(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/gitops/flagger.png")}, c.opts, opts)
	return node.New(nopts...)
}
