package gcp

import "github.com/blushft/go-diagrams/diagram"

type securityContainer struct {
	path string
	opts []diagram.NodeOption
}

var Security = &securityContainer{
	opts: diagram.OptionSet{diagram.Provider("gcp"), diagram.NodeShape("none")},
	path: "assets/gcp/security",
}

func (c *securityContainer) Iap(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/security/iap.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) KeyManagementService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/security/key-management-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) ResourceManager(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/security/resource-manager.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) SecurityCommandCenter(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/security/security-command-center.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) SecurityScanner(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/security/security-scanner.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Iam(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/security/iam.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
