package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type ociSecurityContainer struct {
	path  string
	attrs []attr.Attribute
}

var OciSecurity = &ociSecurityContainer{path: "assets/oci/security"}

func (c *ociSecurityContainer) DdosWhite(opts ...attr.Attribute) *node.Node {
	return node.New("ddos-white", attr.AssetImage("assets/oci/security/ddos-white.png"), attr.Shape(attr.None))
}

func (c *ociSecurityContainer) Waf(opts ...attr.Attribute) *node.Node {
	return node.New("waf", attr.AssetImage("assets/oci/security/waf.png"), attr.Shape(attr.None))
}

func (c *ociSecurityContainer) MaxSecurityZoneWhite(opts ...attr.Attribute) *node.Node {
	return node.New("max-security-zone-white", attr.AssetImage("assets/oci/security/max-security-zone-white.png"), attr.Shape(attr.None))
}

func (c *ociSecurityContainer) MaxSecurityZone(opts ...attr.Attribute) *node.Node {
	return node.New("max-security-zone", attr.AssetImage("assets/oci/security/max-security-zone.png"), attr.Shape(attr.None))
}

func (c *ociSecurityContainer) VaultWhite(opts ...attr.Attribute) *node.Node {
	return node.New("vault-white", attr.AssetImage("assets/oci/security/vault-white.png"), attr.Shape(attr.None))
}

func (c *ociSecurityContainer) WafWhite(opts ...attr.Attribute) *node.Node {
	return node.New("waf-white", attr.AssetImage("assets/oci/security/waf-white.png"), attr.Shape(attr.None))
}

func (c *ociSecurityContainer) CloudGuard(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-guard", attr.AssetImage("assets/oci/security/cloud-guard.png"), attr.Shape(attr.None))
}

func (c *ociSecurityContainer) IdAccessWhite(opts ...attr.Attribute) *node.Node {
	return node.New("id-access-white", attr.AssetImage("assets/oci/security/id-access-white.png"), attr.Shape(attr.None))
}

func (c *ociSecurityContainer) KeyManagement(opts ...attr.Attribute) *node.Node {
	return node.New("key-management", attr.AssetImage("assets/oci/security/key-management.png"), attr.Shape(attr.None))
}

func (c *ociSecurityContainer) KeyManagementWhite(opts ...attr.Attribute) *node.Node {
	return node.New("key-management-white", attr.AssetImage("assets/oci/security/key-management-white.png"), attr.Shape(attr.None))
}

func (c *ociSecurityContainer) Ddos(opts ...attr.Attribute) *node.Node {
	return node.New("ddos", attr.AssetImage("assets/oci/security/ddos.png"), attr.Shape(attr.None))
}

func (c *ociSecurityContainer) EncryptionWhite(opts ...attr.Attribute) *node.Node {
	return node.New("encryption-white", attr.AssetImage("assets/oci/security/encryption-white.png"), attr.Shape(attr.None))
}

func (c *ociSecurityContainer) Encryption(opts ...attr.Attribute) *node.Node {
	return node.New("encryption", attr.AssetImage("assets/oci/security/encryption.png"), attr.Shape(attr.None))
}

func (c *ociSecurityContainer) CloudGuardWhite(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-guard-white", attr.AssetImage("assets/oci/security/cloud-guard-white.png"), attr.Shape(attr.None))
}

func (c *ociSecurityContainer) IdAccess(opts ...attr.Attribute) *node.Node {
	return node.New("id-access", attr.AssetImage("assets/oci/security/id-access.png"), attr.Shape(attr.None))
}

func (c *ociSecurityContainer) Vault(opts ...attr.Attribute) *node.Node {
	return node.New("vault", attr.AssetImage("assets/oci/security/vault.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "oci.security"
  Register(namespace, "DdosWhite", OciSecurity.DdosWhite)
  Register(namespace, "Waf", OciSecurity.Waf)
  Register(namespace, "MaxSecurityZoneWhite", OciSecurity.MaxSecurityZoneWhite)
  Register(namespace, "MaxSecurityZone", OciSecurity.MaxSecurityZone)
  Register(namespace, "VaultWhite", OciSecurity.VaultWhite)
  Register(namespace, "WafWhite", OciSecurity.WafWhite)
  Register(namespace, "CloudGuard", OciSecurity.CloudGuard)
  Register(namespace, "IdAccessWhite", OciSecurity.IdAccessWhite)
  Register(namespace, "KeyManagement", OciSecurity.KeyManagement)
  Register(namespace, "KeyManagementWhite", OciSecurity.KeyManagementWhite)
  Register(namespace, "Ddos", OciSecurity.Ddos)
  Register(namespace, "EncryptionWhite", OciSecurity.EncryptionWhite)
  Register(namespace, "Encryption", OciSecurity.Encryption)
  Register(namespace, "CloudGuardWhite", OciSecurity.CloudGuardWhite)
  Register(namespace, "IdAccess", OciSecurity.IdAccess)
  Register(namespace, "Vault", OciSecurity.Vault)
}
