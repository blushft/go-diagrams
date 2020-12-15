package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type alibabaSecurityContainer struct {
	path  string
	attrs []attr.Attribute
}

var AlibabacloudSecurity =&alibabaSecurityContainer{path: "assets/alibabacloud/security"}

func (c *alibabaSecurityContainer) CloudSecurityScanner(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/cloud-security-scanner.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloud-security-scanner", opts...)
}

func (c *alibabaSecurityContainer) DataEncryptionService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/data-encryption-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-encryption-service", opts...)
}

func (c *alibabaSecurityContainer) IdVerification(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/id-verification.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("id-verification", opts...)
}

func (c *alibabaSecurityContainer) SecurityCenter(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/security-center.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("security-center", opts...)
}

func (c *alibabaSecurityContainer) SslCertificates(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/ssl-certificates.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ssl-certificates", opts...)
}

func (c *alibabaSecurityContainer) AntiBotService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/anti-bot-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("anti-bot-service", opts...)
}

func (c *alibabaSecurityContainer) BastionHost(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/bastion-host.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("bastion-host", opts...)
}

func (c *alibabaSecurityContainer) ManagedSecurityService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/managed-security-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("managed-security-service", opts...)
}

func (c *alibabaSecurityContainer) ServerGuard(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/server-guard.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("server-guard", opts...)
}

func (c *alibabaSecurityContainer) AntiDdosBasic(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/anti-ddos-basic.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("anti-ddos-basic", opts...)
}

func (c *alibabaSecurityContainer) AntiDdosPro(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/anti-ddos-pro.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("anti-ddos-pro", opts...)
}

func (c *alibabaSecurityContainer) CrowdsourcedSecurityTesting(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/crowdsourced-security-testing.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("crowdsourced-security-testing", opts...)
}

func (c *alibabaSecurityContainer) GameShield(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/game-shield.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("game-shield", opts...)
}

func (c *alibabaSecurityContainer) WebApplicationFirewall(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/web-application-firewall.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("web-application-firewall", opts...)
}

func (c *alibabaSecurityContainer) AntifraudService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/antifraud-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("antifraud-service", opts...)
}

func (c *alibabaSecurityContainer) CloudFirewall(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/cloud-firewall.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloud-firewall", opts...)
}

func (c *alibabaSecurityContainer) ContentModeration(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/content-moderation.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("content-moderation", opts...)
}

func (c *alibabaSecurityContainer) DbAudit(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/security/db-audit.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("db-audit", opts...)
}

func init() {
  const namespace = "alibabacloud.security"
  Register(namespace, "CloudSecurityScanner", AlibabacloudSecurity.CloudSecurityScanner)
  Register(namespace, "DataEncryptionService", AlibabacloudSecurity.DataEncryptionService)
  Register(namespace, "IdVerification", AlibabacloudSecurity.IdVerification)
  Register(namespace, "SecurityCenter", AlibabacloudSecurity.SecurityCenter)
  Register(namespace, "SslCertificates", AlibabacloudSecurity.SslCertificates)
  Register(namespace, "AntiBotService", AlibabacloudSecurity.AntiBotService)
  Register(namespace, "BastionHost", AlibabacloudSecurity.BastionHost)
  Register(namespace, "ManagedSecurityService", AlibabacloudSecurity.ManagedSecurityService)
  Register(namespace, "ServerGuard", AlibabacloudSecurity.ServerGuard)
  Register(namespace, "AntiDdosBasic", AlibabacloudSecurity.AntiDdosBasic)
  Register(namespace, "AntiDdosPro", AlibabacloudSecurity.AntiDdosPro)
  Register(namespace, "CrowdsourcedSecurityTesting", AlibabacloudSecurity.CrowdsourcedSecurityTesting)
  Register(namespace, "GameShield", AlibabacloudSecurity.GameShield)
  Register(namespace, "WebApplicationFirewall", AlibabacloudSecurity.WebApplicationFirewall)
  Register(namespace, "AntifraudService", AlibabacloudSecurity.AntifraudService)
  Register(namespace, "CloudFirewall", AlibabacloudSecurity.CloudFirewall)
  Register(namespace, "ContentModeration", AlibabacloudSecurity.ContentModeration)
  Register(namespace, "DbAudit", AlibabacloudSecurity.DbAudit)
}
