package alibabacloud

import "github.com/blushft/go-diagrams/diagram"

type securityContainer struct {
	path string
	opts []diagram.NodeOption
}

var Security = &securityContainer{
	opts: diagram.OptionSet{diagram.Provider("alibabacloud"), diagram.NodeShape("none")},
	path: "assets/alibabacloud/security",
}

func (c *securityContainer) AntiBotService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/anti-bot-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) AntiDdosPro(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/anti-ddos-pro.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) BastionHost(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/bastion-host.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) CloudFirewall(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/cloud-firewall.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) CrowdsourcedSecurityTesting(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/crowdsourced-security-testing.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) GameShield(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/game-shield.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) SecurityCenter(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/security-center.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) AntifraudService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/antifraud-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) CloudSecurityScanner(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/cloud-security-scanner.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) ContentModeration(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/content-moderation.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) DbAudit(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/db-audit.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) AntiDdosBasic(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/anti-ddos-basic.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) DataEncryptionService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/data-encryption-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) IdVerification(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/id-verification.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) ManagedSecurityService(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/managed-security-service.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) ServerGuard(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/server-guard.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) SslCertificates(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/ssl-certificates.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *securityContainer) WebApplicationFirewall(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/security/web-application-firewall.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
