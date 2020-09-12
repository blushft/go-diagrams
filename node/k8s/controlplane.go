package k8s

import "github.com/blushft/go-diagrams/node"

type controlplaneContainer struct {
	path string
	opts []node.Option
}

var Controlplane = &controlplaneContainer{
	opts: node.OptionSet{node.Provider("k8s"), node.Shape("none")},
	path: "assets/k8s/controlplane",
}

func (c *controlplaneContainer) Api(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/controlplane/api.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *controlplaneContainer) CCM(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/controlplane/c-c-m.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *controlplaneContainer) CM(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/controlplane/c-m.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *controlplaneContainer) KProxy(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/controlplane/k-proxy.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *controlplaneContainer) Kubelet(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/controlplane/kubelet.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *controlplaneContainer) Sched(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/k8s/controlplane/sched.png")}, c.opts, opts)
	return node.New(nopts...)
}
