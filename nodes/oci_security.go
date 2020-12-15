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
	opts = append(opts, attr.AssetImage("assets/oci/security/ddos-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ddos-white", opts...)
}

func (c *ociSecurityContainer) Waf(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/security/waf.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("waf", opts...)
}

func (c *ociSecurityContainer) MaxSecurityZoneWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/security/max-security-zone-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("max-security-zone-white", opts...)
}

func (c *ociSecurityContainer) MaxSecurityZone(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/security/max-security-zone.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("max-security-zone", opts...)
}

func (c *ociSecurityContainer) VaultWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/security/vault-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vault-white", opts...)
}

func (c *ociSecurityContainer) WafWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/security/waf-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("waf-white", opts...)
}

func (c *ociSecurityContainer) CloudGuard(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/security/cloud-guard.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloud-guard", opts...)
}

func (c *ociSecurityContainer) IdAccessWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/security/id-access-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("id-access-white", opts...)
}

func (c *ociSecurityContainer) KeyManagement(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/security/key-management.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("key-management", opts...)
}

func (c *ociSecurityContainer) KeyManagementWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/security/key-management-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("key-management-white", opts...)
}

func (c *ociSecurityContainer) Ddos(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/security/ddos.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ddos", opts...)
}

func (c *ociSecurityContainer) EncryptionWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/security/encryption-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("encryption-white", opts...)
}

func (c *ociSecurityContainer) Encryption(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/security/encryption.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("encryption", opts...)
}

func (c *ociSecurityContainer) CloudGuardWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/security/cloud-guard-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloud-guard-white", opts...)
}

func (c *ociSecurityContainer) IdAccess(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/security/id-access.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("id-access", opts...)
}

func (c *ociSecurityContainer) Vault(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/security/vault.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vault", opts...)
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
