package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type azureIdentityContainer struct {
	path  string
	attrs []attr.Attribute
}

var azureidentityContainer = &azureIdentityContainer{path: "assets/azure/identity"}

func (c *azureIdentityContainer) ManagedIdentities(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/identity/managed-identities.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("managed-identities", opts...)
}

func (c *azureIdentityContainer) AccessReview(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/identity/access-review.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("access-review", opts...)
}

func (c *azureIdentityContainer) ActiveDirectory(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/identity/active-directory.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("active-directory", opts...)
}

func (c *azureIdentityContainer) AppRegistrations(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/identity/app-registrations.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("app-registrations", opts...)
}

func (c *azureIdentityContainer) AdIdentityProtection(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/identity/ad-identity-protection.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ad-identity-protection", opts...)
}

func (c *azureIdentityContainer) AdPrivilegedIdentityManagement(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/identity/ad-privileged-identity-management.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ad-privileged-identity-management", opts...)
}

func (c *azureIdentityContainer) ConditionalAccess(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/identity/conditional-access.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("conditional-access", opts...)
}

func (c *azureIdentityContainer) EnterpriseApplications(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/identity/enterprise-applications.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("enterprise-applications", opts...)
}

func (c *azureIdentityContainer) IdentityGovernance(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/identity/identity-governance.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("identity-governance", opts...)
}

func (c *azureIdentityContainer) ActiveDirectoryConnectHealth(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/identity/active-directory-connect-health.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("active-directory-connect-health", opts...)
}

func (c *azureIdentityContainer) AdB2C(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/identity/ad-b2c.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ad-b2c", opts...)
}

func (c *azureIdentityContainer) AdDomainServices(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/identity/ad-domain-services.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ad-domain-services", opts...)
}

func (c *azureIdentityContainer) InformationProtection(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/azure/identity/information-protection.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("information-protection", opts...)
}

func init() {
  const namespace = "azure.identity"
  Register(namespace, "ManagedIdentities", azureidentityContainer.ManagedIdentities)
  Register(namespace, "AccessReview", azureidentityContainer.AccessReview)
  Register(namespace, "ActiveDirectory", azureidentityContainer.ActiveDirectory)
  Register(namespace, "AppRegistrations", azureidentityContainer.AppRegistrations)
  Register(namespace, "AdIdentityProtection", azureidentityContainer.AdIdentityProtection)
  Register(namespace, "AdPrivilegedIdentityManagement", azureidentityContainer.AdPrivilegedIdentityManagement)
  Register(namespace, "ConditionalAccess", azureidentityContainer.ConditionalAccess)
  Register(namespace, "EnterpriseApplications", azureidentityContainer.EnterpriseApplications)
  Register(namespace, "IdentityGovernance", azureidentityContainer.IdentityGovernance)
  Register(namespace, "ActiveDirectoryConnectHealth", azureidentityContainer.ActiveDirectoryConnectHealth)
  Register(namespace, "AdB2C", azureidentityContainer.AdB2C)
  Register(namespace, "AdDomainServices", azureidentityContainer.AdDomainServices)
  Register(namespace, "InformationProtection", azureidentityContainer.InformationProtection)
}
