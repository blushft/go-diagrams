package k8s

import "github.com/blushft/go-diagrams/diagram"

type controlplaneContainer struct {
	path string
	opts []diagram.NodeOption
}

var Controlplane = &controlplaneContainer{
	opts: diagram.OptionSet{diagram.Provider("k8s"), diagram.NodeShape("none")},
	path: "assets/k8s/controlplane",
}

func (c *controlplaneContainer) KProxy(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/controlplane/k-proxy.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *controlplaneContainer) Kubelet(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/controlplane/kubelet.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *controlplaneContainer) Sched(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/controlplane/sched.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *controlplaneContainer) Api(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/controlplane/api.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *controlplaneContainer) CCM(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/controlplane/c-c-m.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *controlplaneContainer) CM(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/k8s/controlplane/c-m.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
