package alibabacloud

import "github.com/blushft/go-diagrams/node"

type securityContainer struct {
	path string
	opts []node.Option
}

var Security = &securityContainer{
	opts: node.OptionSet{node.Provider("alibabacloud"), node.Shape("none")},
	path: "assets/alibabacloud/security",
}

func (c *securityContainer) AntifraudService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/antifraud-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) DataEncryptionService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/data-encryption-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) WebApplicationFirewall(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/web-application-firewall.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) BastionHost(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/bastion-host.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) CloudFirewall(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/cloud-firewall.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) CloudSecurityScanner(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/cloud-security-scanner.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) ContentModeration(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/content-moderation.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) SecurityCenter(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/security-center.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) ServerGuard(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/server-guard.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) ManagedSecurityService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/managed-security-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) SslCertificates(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/ssl-certificates.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) AntiBotService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/anti-bot-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) AntiDdosBasic(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/anti-ddos-basic.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) AntiDdosPro(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/anti-ddos-pro.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) CrowdsourcedSecurityTesting(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/crowdsourced-security-testing.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) DbAudit(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/db-audit.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) GameShield(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/game-shield.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) IdVerification(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/security/id-verification.png")}, c.opts, opts)
	return node.New(nopts...)
}
