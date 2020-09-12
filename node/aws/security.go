package aws

import "github.com/blushft/go-diagrams/node"

type securityContainer struct {
	path string
	opts []node.Option
}

var Security = &securityContainer{
	opts: node.OptionSet{node.Provider("aws"), node.Shape("none")},
	path: "assets/aws/security",
}

func (c *securityContainer) CloudDirectory(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/cloud-directory.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Cloudhsm(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/cloudhsm.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) DirectoryService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/directory-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) FirewallManager(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/firewall-manager.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagementIamAccessAnalyzer(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/identity-and-access-management-iam-access-analyzer.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagementIamRole(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/identity-and-access-management-iam-role.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Macie(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/macie.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Artifact(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/artifact.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Waf(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/waf.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) SingleSignOn(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/single-sign-on.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Inspector(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/inspector.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) SecretsManager(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/secrets-manager.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) SecurityIdentityAndCompliance(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/security-identity-and-compliance.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Guardduty(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/guardduty.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) SecurityHub(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/security-hub.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Shield(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/shield.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagementIamAwsSts(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/identity-and-access-management-iam-aws-sts.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Cognito(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/cognito.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) Detective(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/detective.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagementIamPermissions(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/identity-and-access-management-iam-permissions.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) IdentityAndAccessManagementIam(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/identity-and-access-management-iam.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) KeyManagementService(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/key-management-service.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) ResourceAccessManager(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/resource-access-manager.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *securityContainer) CertificateManager(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/aws/security/certificate-manager.png")}, c.opts, opts)
	return node.New(nopts...)
}
