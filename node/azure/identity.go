package azure

import "github.com/blushft/go-diagrams/node"

type identityContainer struct {
	path string
	opts []node.Option
}

var Identity = &identityContainer{
	opts: node.OptionSet{node.Provider("azure"), node.Shape("none")},
	path: "assets/azure/identity",
}

func (c *identityContainer) AccessReview(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/identity/access-review.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *identityContainer) AdB2C(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/identity/ad-b2c.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *identityContainer) AppRegistrations(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/identity/app-registrations.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *identityContainer) EnterpriseApplications(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/identity/enterprise-applications.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *identityContainer) ManagedIdentities(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/identity/managed-identities.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *identityContainer) ConditionalAccess(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/identity/conditional-access.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *identityContainer) IdentityGovernance(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/identity/identity-governance.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *identityContainer) InformationProtection(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/identity/information-protection.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *identityContainer) ActiveDirectoryConnectHealth(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/identity/active-directory-connect-health.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *identityContainer) ActiveDirectory(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/identity/active-directory.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *identityContainer) AdDomainServices(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/identity/ad-domain-services.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *identityContainer) AdIdentityProtection(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/identity/ad-identity-protection.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *identityContainer) AdPrivilegedIdentityManagement(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/azure/identity/ad-privileged-identity-management.png")}, c.opts, opts)
	return node.New(nopts...)
}
