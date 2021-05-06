package oci

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type securityContainer struct {
	path  string
	attrs []attr.Attribute
}

var Security = &securityContainer{path: "assets/oci/security"}

func (c *securityContainer) Vault(opts ...attr.Attribute) *node.Node {
	return node.New("vault", attr.AssetImage("assets/oci/security/vault.png"), attr.Shape(attr.None))
}

func (c *securityContainer) CloudGuard(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-guard", attr.AssetImage("assets/oci/security/cloud-guard.png"), attr.Shape(attr.None))
}

func (c *securityContainer) Ddos(opts ...attr.Attribute) *node.Node {
	return node.New("ddos", attr.AssetImage("assets/oci/security/ddos.png"), attr.Shape(attr.None))
}

func (c *securityContainer) EncryptionWhite(opts ...attr.Attribute) *node.Node {
	return node.New("encryption-white", attr.AssetImage("assets/oci/security/encryption-white.png"), attr.Shape(attr.None))
}

func (c *securityContainer) KeyManagement(opts ...attr.Attribute) *node.Node {
	return node.New("key-management", attr.AssetImage("assets/oci/security/key-management.png"), attr.Shape(attr.None))
}

func (c *securityContainer) MaxSecurityZoneWhite(opts ...attr.Attribute) *node.Node {
	return node.New("max-security-zone-white", attr.AssetImage("assets/oci/security/max-security-zone-white.png"), attr.Shape(attr.None))
}

func (c *securityContainer) CloudGuardWhite(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-guard-white", attr.AssetImage("assets/oci/security/cloud-guard-white.png"), attr.Shape(attr.None))
}

func (c *securityContainer) Encryption(opts ...attr.Attribute) *node.Node {
	return node.New("encryption", attr.AssetImage("assets/oci/security/encryption.png"), attr.Shape(attr.None))
}

func (c *securityContainer) KeyManagementWhite(opts ...attr.Attribute) *node.Node {
	return node.New("key-management-white", attr.AssetImage("assets/oci/security/key-management-white.png"), attr.Shape(attr.None))
}

func (c *securityContainer) VaultWhite(opts ...attr.Attribute) *node.Node {
	return node.New("vault-white", attr.AssetImage("assets/oci/security/vault-white.png"), attr.Shape(attr.None))
}

func (c *securityContainer) IdAccessWhite(opts ...attr.Attribute) *node.Node {
	return node.New("id-access-white", attr.AssetImage("assets/oci/security/id-access-white.png"), attr.Shape(attr.None))
}

func (c *securityContainer) IdAccess(opts ...attr.Attribute) *node.Node {
	return node.New("id-access", attr.AssetImage("assets/oci/security/id-access.png"), attr.Shape(attr.None))
}

func (c *securityContainer) MaxSecurityZone(opts ...attr.Attribute) *node.Node {
	return node.New("max-security-zone", attr.AssetImage("assets/oci/security/max-security-zone.png"), attr.Shape(attr.None))
}

func (c *securityContainer) WafWhite(opts ...attr.Attribute) *node.Node {
	return node.New("waf-white", attr.AssetImage("assets/oci/security/waf-white.png"), attr.Shape(attr.None))
}

func (c *securityContainer) DdosWhite(opts ...attr.Attribute) *node.Node {
	return node.New("ddos-white", attr.AssetImage("assets/oci/security/ddos-white.png"), attr.Shape(attr.None))
}

func (c *securityContainer) Waf(opts ...attr.Attribute) *node.Node {
	return node.New("waf", attr.AssetImage("assets/oci/security/waf.png"), attr.Shape(attr.None))
}
