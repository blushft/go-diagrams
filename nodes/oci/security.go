package oci

import "github.com/blushft/go-diagrams/diagram"

type securityContainer struct {
	path string
	opts []diagram.NodeOption
}

var Security = &securityContainer{
	opts: diagram.OptionSet{diagram.Provider("oci"), diagram.NodeShape("none")},
	path: "assets/oci/security",
}

func (c *securityContainer) MaxSecurityZone(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/security/max-security-zone.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) WafWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/security/waf-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) CloudGuardWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/security/cloud-guard-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) KeyManagementWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/security/key-management-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) KeyManagement(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/security/key-management.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) CloudGuard(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/security/cloud-guard.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) VaultWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/security/vault-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Vault(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/security/vault.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Waf(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/security/waf.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Ddos(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/security/ddos.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) Encryption(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/security/encryption.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdAccess(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/security/id-access.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) MaxSecurityZoneWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/security/max-security-zone-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) DdosWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/security/ddos-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) EncryptionWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/security/encryption-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdAccessWhite(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/oci/security/id-access-white.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
