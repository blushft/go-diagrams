package oci

import "github.com/blushft/go-diagrams/node"

type securityContainer struct {
	path string
	opts []node.Option
}

var Security = &securityContainer{
	opts: node.OptionSet{node.Provider("oci"), node.Shape("none")},
	path: "assets/oci/security",
}

func (c *securityContainer) EncryptionWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/security/encryption-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) MaxSecurityZoneWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/security/max-security-zone-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Vault(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/security/vault.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Ddos(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/security/ddos.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Encryption(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/security/encryption.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) IdAccess(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/security/id-access.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) KeyManagement(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/security/key-management.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) DdosWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/security/ddos-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) MaxSecurityZone(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/security/max-security-zone.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) WafWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/security/waf-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Waf(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/security/waf.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) VaultWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/security/vault-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) CloudGuardWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/security/cloud-guard-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) CloudGuard(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/security/cloud-guard.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) IdAccessWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/security/id-access-white.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) KeyManagementWhite(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/oci/security/key-management-white.png")}, c.opts, opts)
	return node.New(nopts...)
}
