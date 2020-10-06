package alibabacloud

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type securityContainer struct {
	path  string
	attrs []attr.Attribute
}

var Security = &securityContainer{path: "assets/alibabacloud/security"}

func (c *securityContainer) CloudSecurityScanner(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-security-scanner", attr.AssetImage("assets/alibabacloud/security/cloud-security-scanner.png"), attr.Shape(attr.None))
}

func (c *securityContainer) DataEncryptionService(opts ...attr.Attribute) *node.Node {
	return node.New("data-encryption-service", attr.AssetImage("assets/alibabacloud/security/data-encryption-service.png"), attr.Shape(attr.None))
}

func (c *securityContainer) IdVerification(opts ...attr.Attribute) *node.Node {
	return node.New("id-verification", attr.AssetImage("assets/alibabacloud/security/id-verification.png"), attr.Shape(attr.None))
}

func (c *securityContainer) SecurityCenter(opts ...attr.Attribute) *node.Node {
	return node.New("security-center", attr.AssetImage("assets/alibabacloud/security/security-center.png"), attr.Shape(attr.None))
}

func (c *securityContainer) SslCertificates(opts ...attr.Attribute) *node.Node {
	return node.New("ssl-certificates", attr.AssetImage("assets/alibabacloud/security/ssl-certificates.png"), attr.Shape(attr.None))
}

func (c *securityContainer) AntiBotService(opts ...attr.Attribute) *node.Node {
	return node.New("anti-bot-service", attr.AssetImage("assets/alibabacloud/security/anti-bot-service.png"), attr.Shape(attr.None))
}

func (c *securityContainer) BastionHost(opts ...attr.Attribute) *node.Node {
	return node.New("bastion-host", attr.AssetImage("assets/alibabacloud/security/bastion-host.png"), attr.Shape(attr.None))
}

func (c *securityContainer) ManagedSecurityService(opts ...attr.Attribute) *node.Node {
	return node.New("managed-security-service", attr.AssetImage("assets/alibabacloud/security/managed-security-service.png"), attr.Shape(attr.None))
}

func (c *securityContainer) ServerGuard(opts ...attr.Attribute) *node.Node {
	return node.New("server-guard", attr.AssetImage("assets/alibabacloud/security/server-guard.png"), attr.Shape(attr.None))
}

func (c *securityContainer) AntiDdosBasic(opts ...attr.Attribute) *node.Node {
	return node.New("anti-ddos-basic", attr.AssetImage("assets/alibabacloud/security/anti-ddos-basic.png"), attr.Shape(attr.None))
}

func (c *securityContainer) AntiDdosPro(opts ...attr.Attribute) *node.Node {
	return node.New("anti-ddos-pro", attr.AssetImage("assets/alibabacloud/security/anti-ddos-pro.png"), attr.Shape(attr.None))
}

func (c *securityContainer) CrowdsourcedSecurityTesting(opts ...attr.Attribute) *node.Node {
	return node.New("crowdsourced-security-testing", attr.AssetImage("assets/alibabacloud/security/crowdsourced-security-testing.png"), attr.Shape(attr.None))
}

func (c *securityContainer) GameShield(opts ...attr.Attribute) *node.Node {
	return node.New("game-shield", attr.AssetImage("assets/alibabacloud/security/game-shield.png"), attr.Shape(attr.None))
}

func (c *securityContainer) WebApplicationFirewall(opts ...attr.Attribute) *node.Node {
	return node.New("web-application-firewall", attr.AssetImage("assets/alibabacloud/security/web-application-firewall.png"), attr.Shape(attr.None))
}

func (c *securityContainer) AntifraudService(opts ...attr.Attribute) *node.Node {
	return node.New("antifraud-service", attr.AssetImage("assets/alibabacloud/security/antifraud-service.png"), attr.Shape(attr.None))
}

func (c *securityContainer) CloudFirewall(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-firewall", attr.AssetImage("assets/alibabacloud/security/cloud-firewall.png"), attr.Shape(attr.None))
}

func (c *securityContainer) ContentModeration(opts ...attr.Attribute) *node.Node {
	return node.New("content-moderation", attr.AssetImage("assets/alibabacloud/security/content-moderation.png"), attr.Shape(attr.None))
}

func (c *securityContainer) DbAudit(opts ...attr.Attribute) *node.Node {
	return node.New("db-audit", attr.AssetImage("assets/alibabacloud/security/db-audit.png"), attr.Shape(attr.None))
}
