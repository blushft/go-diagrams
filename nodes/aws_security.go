package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type awsSecurityContainer struct {
	path  string
	attrs []attr.Attribute
}

var AWSSecurity = &awsSecurityContainer{path: "assets/aws/security"}

func (c *awsSecurityContainer) SingleSignOn(opts ...attr.Attribute) *node.Node {
	return node.New("single-sign-on", attr.AssetImage("assets/aws/security/single-sign-on.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) Waf(opts ...attr.Attribute) *node.Node {
	return node.New("waf", attr.AssetImage("assets/aws/security/waf.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) CloudDirectory(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-directory", attr.AssetImage("assets/aws/security/cloud-directory.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) IdentityAndAccessManagementIamAwsSts(opts ...attr.Attribute) *node.Node {
	return node.New("identity-and-access-management-iam-aws-sts", attr.AssetImage("assets/aws/security/identity-and-access-management-iam-aws-sts.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) KeyManagementService(opts ...attr.Attribute) *node.Node {
	return node.New("key-management-service", attr.AssetImage("assets/aws/security/key-management-service.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) Detective(opts ...attr.Attribute) *node.Node {
	return node.New("detective", attr.AssetImage("assets/aws/security/detective.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) Inspector(opts ...attr.Attribute) *node.Node {
	return node.New("inspector", attr.AssetImage("assets/aws/security/inspector.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) DirectoryService(opts ...attr.Attribute) *node.Node {
	return node.New("directory-service", attr.AssetImage("assets/aws/security/directory-service.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) Guardduty(opts ...attr.Attribute) *node.Node {
	return node.New("guardduty", attr.AssetImage("assets/aws/security/guardduty.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) IdentityAndAccessManagementIam(opts ...attr.Attribute) *node.Node {
	return node.New("identity-and-access-management-iam", attr.AssetImage("assets/aws/security/identity-and-access-management-iam.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) Macie(opts ...attr.Attribute) *node.Node {
	return node.New("macie", attr.AssetImage("assets/aws/security/macie.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) SecretsManager(opts ...attr.Attribute) *node.Node {
	return node.New("secrets-manager", attr.AssetImage("assets/aws/security/secrets-manager.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) CertificateManager(opts ...attr.Attribute) *node.Node {
	return node.New("certificate-manager", attr.AssetImage("assets/aws/security/certificate-manager.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) Cloudhsm(opts ...attr.Attribute) *node.Node {
	return node.New("cloudhsm", attr.AssetImage("assets/aws/security/cloudhsm.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) Cognito(opts ...attr.Attribute) *node.Node {
	return node.New("cognito", attr.AssetImage("assets/aws/security/cognito.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) SecurityHub(opts ...attr.Attribute) *node.Node {
	return node.New("security-hub", attr.AssetImage("assets/aws/security/security-hub.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) IdentityAndAccessManagementIamPermissions(opts ...attr.Attribute) *node.Node {
	return node.New("identity-and-access-management-iam-permissions", attr.AssetImage("assets/aws/security/identity-and-access-management-iam-permissions.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) IdentityAndAccessManagementIamRole(opts ...attr.Attribute) *node.Node {
	return node.New("identity-and-access-management-iam-role", attr.AssetImage("assets/aws/security/identity-and-access-management-iam-role.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) ResourceAccessManager(opts ...attr.Attribute) *node.Node {
	return node.New("resource-access-manager", attr.AssetImage("assets/aws/security/resource-access-manager.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) SecurityIdentityAndCompliance(opts ...attr.Attribute) *node.Node {
	return node.New("security-identity-and-compliance", attr.AssetImage("assets/aws/security/security-identity-and-compliance.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) Shield(opts ...attr.Attribute) *node.Node {
	return node.New("shield", attr.AssetImage("assets/aws/security/shield.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) Artifact(opts ...attr.Attribute) *node.Node {
	return node.New("artifact", attr.AssetImage("assets/aws/security/artifact.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) FirewallManager(opts ...attr.Attribute) *node.Node {
	return node.New("firewall-manager", attr.AssetImage("assets/aws/security/firewall-manager.png"), attr.Shape(attr.None))
}

func (c *awsSecurityContainer) IdentityAndAccessManagementIamAccessAnalyzer(opts ...attr.Attribute) *node.Node {
	return node.New("identity-and-access-management-iam-access-analyzer", attr.AssetImage("assets/aws/security/identity-and-access-management-iam-access-analyzer.png"), attr.Shape(attr.None))
}

func init() {
  const namespace= "aws.security"
  Register(namespace, "SingleSignOn", AWSSecurity.SingleSignOn)
  Register(namespace, "Waf", AWSSecurity.Waf)
  Register(namespace, "CloudDirectory", AWSSecurity.CloudDirectory)
  Register(namespace, "IdentityAndAccessManagementIamAwsSts", AWSSecurity.IdentityAndAccessManagementIamAwsSts)
  Register(namespace, "KeyManagementService", AWSSecurity.KeyManagementService)
  Register(namespace, "Detective", AWSSecurity.Detective)
  Register(namespace, "Inspector", AWSSecurity.Inspector)
  Register(namespace, "DirectoryService", AWSSecurity.DirectoryService)
  Register(namespace, "Guardduty", AWSSecurity.Guardduty)
  Register(namespace, "IdentityAndAccessManagementIam", AWSSecurity.IdentityAndAccessManagementIam)
  Register(namespace, "Macie", AWSSecurity.Macie)
  Register(namespace, "SecretsManager", AWSSecurity.SecretsManager)
  Register(namespace, "CertificateManager", AWSSecurity.CertificateManager)
  Register(namespace, "Cloudhsm", AWSSecurity.Cloudhsm)
  Register(namespace, "Cognito", AWSSecurity.Cognito)
  Register(namespace, "SecurityHub", AWSSecurity.SecurityHub)
  Register(namespace, "IdentityAndAccessManagementIamPermissions", AWSSecurity.IdentityAndAccessManagementIamPermissions)
  Register(namespace, "IdentityAndAccessManagementIamRole", AWSSecurity.IdentityAndAccessManagementIamRole)
  Register(namespace, "ResourceAccessManager", AWSSecurity.ResourceAccessManager)
  Register(namespace, "SecurityIdentityAndCompliance", AWSSecurity.SecurityIdentityAndCompliance)
  Register(namespace, "Shield", AWSSecurity.Shield)
  Register(namespace, "Artifact", AWSSecurity.Artifact)
  Register(namespace, "FirewallManager", AWSSecurity.FirewallManager)
  Register(namespace, "IdentityAndAccessManagementIamAccessAnalyzer", AWSSecurity.IdentityAndAccessManagementIamAccessAnalyzer)
}
