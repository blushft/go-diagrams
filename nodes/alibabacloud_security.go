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
	return node.New("cloud-security-scanner", attr.AssetImage("assets/alibabacloud/security/cloud-security-scanner.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) DataEncryptionService(opts ...attr.Attribute) *node.Node {
	return node.New("data-encryption-service", attr.AssetImage("assets/alibabacloud/security/data-encryption-service.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) IdVerification(opts ...attr.Attribute) *node.Node {
	return node.New("id-verification", attr.AssetImage("assets/alibabacloud/security/id-verification.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) SecurityCenter(opts ...attr.Attribute) *node.Node {
	return node.New("security-center", attr.AssetImage("assets/alibabacloud/security/security-center.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) SslCertificates(opts ...attr.Attribute) *node.Node {
	return node.New("ssl-certificates", attr.AssetImage("assets/alibabacloud/security/ssl-certificates.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) AntiBotService(opts ...attr.Attribute) *node.Node {
	return node.New("anti-bot-service", attr.AssetImage("assets/alibabacloud/security/anti-bot-service.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) BastionHost(opts ...attr.Attribute) *node.Node {
	return node.New("bastion-host", attr.AssetImage("assets/alibabacloud/security/bastion-host.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) ManagedSecurityService(opts ...attr.Attribute) *node.Node {
	return node.New("managed-security-service", attr.AssetImage("assets/alibabacloud/security/managed-security-service.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) ServerGuard(opts ...attr.Attribute) *node.Node {
	return node.New("server-guard", attr.AssetImage("assets/alibabacloud/security/server-guard.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) AntiDdosBasic(opts ...attr.Attribute) *node.Node {
	return node.New("anti-ddos-basic", attr.AssetImage("assets/alibabacloud/security/anti-ddos-basic.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) AntiDdosPro(opts ...attr.Attribute) *node.Node {
	return node.New("anti-ddos-pro", attr.AssetImage("assets/alibabacloud/security/anti-ddos-pro.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) CrowdsourcedSecurityTesting(opts ...attr.Attribute) *node.Node {
	return node.New("crowdsourced-security-testing", attr.AssetImage("assets/alibabacloud/security/crowdsourced-security-testing.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) GameShield(opts ...attr.Attribute) *node.Node {
	return node.New("game-shield", attr.AssetImage("assets/alibabacloud/security/game-shield.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) WebApplicationFirewall(opts ...attr.Attribute) *node.Node {
	return node.New("web-application-firewall", attr.AssetImage("assets/alibabacloud/security/web-application-firewall.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) AntifraudService(opts ...attr.Attribute) *node.Node {
	return node.New("antifraud-service", attr.AssetImage("assets/alibabacloud/security/antifraud-service.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) CloudFirewall(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-firewall", attr.AssetImage("assets/alibabacloud/security/cloud-firewall.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) ContentModeration(opts ...attr.Attribute) *node.Node {
	return node.New("content-moderation", attr.AssetImage("assets/alibabacloud/security/content-moderation.png"), attr.Shape(attr.None))
}

func (c *alibabaSecurityContainer) DbAudit(opts ...attr.Attribute) *node.Node {
	return node.New("db-audit", attr.AssetImage("assets/alibabacloud/security/db-audit.png"), attr.Shape(attr.None))
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
