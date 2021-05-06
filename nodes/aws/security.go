package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type securityContainer struct {
	path  string
	attrs []attr.Attribute
}

var Security = &securityContainer{path: "assets/aws/security"}

func (c *securityContainer) Cloudhsm(opts ...attr.Attribute) *node.Node {
	return node.New("cloudhsm", attr.AssetImage("assets/aws/security/cloudhsm.png"), attr.Shape(attr.None))
}

func (c *securityContainer) FirewallManager(opts ...attr.Attribute) *node.Node {
	return node.New("firewall-manager", attr.AssetImage("assets/aws/security/firewall-manager.png"), attr.Shape(attr.None))
}

func (c *securityContainer) Macie(opts ...attr.Attribute) *node.Node {
	return node.New("macie", attr.AssetImage("assets/aws/security/macie.png"), attr.Shape(attr.None))
}

func (c *securityContainer) IdentityAndAccessManagementIamAwsSts(opts ...attr.Attribute) *node.Node {
	return node.New("identity-and-access-management-iam-aws-sts", attr.AssetImage("assets/aws/security/identity-and-access-management-iam-aws-sts.png"), attr.Shape(attr.None))
}

func (c *securityContainer) IdentityAndAccessManagementIamRole(opts ...attr.Attribute) *node.Node {
	return node.New("identity-and-access-management-iam-role", attr.AssetImage("assets/aws/security/identity-and-access-management-iam-role.png"), attr.Shape(attr.None))
}

func (c *securityContainer) IdentityAndAccessManagementIam(opts ...attr.Attribute) *node.Node {
	return node.New("identity-and-access-management-iam", attr.AssetImage("assets/aws/security/identity-and-access-management-iam.png"), attr.Shape(attr.None))
}

func (c *securityContainer) KeyManagementService(opts ...attr.Attribute) *node.Node {
	return node.New("key-management-service", attr.AssetImage("assets/aws/security/key-management-service.png"), attr.Shape(attr.None))
}

func (c *securityContainer) Artifact(opts ...attr.Attribute) *node.Node {
	return node.New("artifact", attr.AssetImage("assets/aws/security/artifact.png"), attr.Shape(attr.None))
}

func (c *securityContainer) CertificateManager(opts ...attr.Attribute) *node.Node {
	return node.New("certificate-manager", attr.AssetImage("assets/aws/security/certificate-manager.png"), attr.Shape(attr.None))
}

func (c *securityContainer) Detective(opts ...attr.Attribute) *node.Node {
	return node.New("detective", attr.AssetImage("assets/aws/security/detective.png"), attr.Shape(attr.None))
}

func (c *securityContainer) Guardduty(opts ...attr.Attribute) *node.Node {
	return node.New("guardduty", attr.AssetImage("assets/aws/security/guardduty.png"), attr.Shape(attr.None))
}

func (c *securityContainer) SingleSignOn(opts ...attr.Attribute) *node.Node {
	return node.New("single-sign-on", attr.AssetImage("assets/aws/security/single-sign-on.png"), attr.Shape(attr.None))
}

func (c *securityContainer) Waf(opts ...attr.Attribute) *node.Node {
	return node.New("waf", attr.AssetImage("assets/aws/security/waf.png"), attr.Shape(attr.None))
}

func (c *securityContainer) SecretsManager(opts ...attr.Attribute) *node.Node {
	return node.New("secrets-manager", attr.AssetImage("assets/aws/security/secrets-manager.png"), attr.Shape(attr.None))
}

func (c *securityContainer) SecurityHub(opts ...attr.Attribute) *node.Node {
	return node.New("security-hub", attr.AssetImage("assets/aws/security/security-hub.png"), attr.Shape(attr.None))
}

func (c *securityContainer) SecurityIdentityAndCompliance(opts ...attr.Attribute) *node.Node {
	return node.New("security-identity-and-compliance", attr.AssetImage("assets/aws/security/security-identity-and-compliance.png"), attr.Shape(attr.None))
}

func (c *securityContainer) CloudDirectory(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-directory", attr.AssetImage("assets/aws/security/cloud-directory.png"), attr.Shape(attr.None))
}

func (c *securityContainer) Cognito(opts ...attr.Attribute) *node.Node {
	return node.New("cognito", attr.AssetImage("assets/aws/security/cognito.png"), attr.Shape(attr.None))
}

func (c *securityContainer) IdentityAndAccessManagementIamAccessAnalyzer(opts ...attr.Attribute) *node.Node {
	return node.New("identity-and-access-management-iam-access-analyzer", attr.AssetImage("assets/aws/security/identity-and-access-management-iam-access-analyzer.png"), attr.Shape(attr.None))
}

func (c *securityContainer) ResourceAccessManager(opts ...attr.Attribute) *node.Node {
	return node.New("resource-access-manager", attr.AssetImage("assets/aws/security/resource-access-manager.png"), attr.Shape(attr.None))
}

func (c *securityContainer) DirectoryService(opts ...attr.Attribute) *node.Node {
	return node.New("directory-service", attr.AssetImage("assets/aws/security/directory-service.png"), attr.Shape(attr.None))
}

func (c *securityContainer) IdentityAndAccessManagementIamPermissions(opts ...attr.Attribute) *node.Node {
	return node.New("identity-and-access-management-iam-permissions", attr.AssetImage("assets/aws/security/identity-and-access-management-iam-permissions.png"), attr.Shape(attr.None))
}

func (c *securityContainer) Inspector(opts ...attr.Attribute) *node.Node {
	return node.New("inspector", attr.AssetImage("assets/aws/security/inspector.png"), attr.Shape(attr.None))
}

func (c *securityContainer) Shield(opts ...attr.Attribute) *node.Node {
	return node.New("shield", attr.AssetImage("assets/aws/security/shield.png"), attr.Shape(attr.None))
}
