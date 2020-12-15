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
	opts = append(opts, attr.AssetImage("assets/aws/security/single-sign-on.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("single-sign-on", opts...)
}

func (c *awsSecurityContainer) Waf(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/waf.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("waf", opts...)
}

func (c *awsSecurityContainer) CloudDirectory(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/cloud-directory.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloud-directory", opts...)
}

func (c *awsSecurityContainer) IdentityAndAccessManagementIamAwsSts(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/identity-and-access-management-iam-aws-sts.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("identity-and-access-management-iam-aws-sts", opts...)
}

func (c *awsSecurityContainer) KeyManagementService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/key-management-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("key-management-service", opts...)
}

func (c *awsSecurityContainer) Detective(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/detective.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("detective", opts...)
}

func (c *awsSecurityContainer) Inspector(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/inspector.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("inspector", opts...)
}

func (c *awsSecurityContainer) DirectoryService(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/directory-service.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("directory-service", opts...)
}

func (c *awsSecurityContainer) Guardduty(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/guardduty.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("guardduty", opts...)
}

func (c *awsSecurityContainer) IdentityAndAccessManagementIam(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/identity-and-access-management-iam.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("identity-and-access-management-iam", opts...)
}

func (c *awsSecurityContainer) Macie(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/macie.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("macie", opts...)
}

func (c *awsSecurityContainer) SecretsManager(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/secrets-manager.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("secrets-manager", opts...)
}

func (c *awsSecurityContainer) CertificateManager(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/certificate-manager.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("certificate-manager", opts...)
}

func (c *awsSecurityContainer) Cloudhsm(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/cloudhsm.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloudhsm", opts...)
}

func (c *awsSecurityContainer) Cognito(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/cognito.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cognito", opts...)
}

func (c *awsSecurityContainer) SecurityHub(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/security-hub.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("security-hub", opts...)
}

func (c *awsSecurityContainer) IdentityAndAccessManagementIamPermissions(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/identity-and-access-management-iam-permissions.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("identity-and-access-management-iam-permissions", opts...)
}

func (c *awsSecurityContainer) IdentityAndAccessManagementIamRole(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/identity-and-access-management-iam-role.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("identity-and-access-management-iam-role", opts...)
}

func (c *awsSecurityContainer) ResourceAccessManager(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/resource-access-manager.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("resource-access-manager", opts...)
}

func (c *awsSecurityContainer) SecurityIdentityAndCompliance(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/security-identity-and-compliance.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("security-identity-and-compliance", opts...)
}

func (c *awsSecurityContainer) Shield(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/shield.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("shield", opts...)
}

func (c *awsSecurityContainer) Artifact(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/artifact.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("artifact", opts...)
}

func (c *awsSecurityContainer) FirewallManager(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/firewall-manager.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("firewall-manager", opts...)
}

func (c *awsSecurityContainer) IdentityAndAccessManagementIamAccessAnalyzer(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/security/identity-and-access-management-iam-access-analyzer.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("identity-and-access-management-iam-access-analyzer", opts...)
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
