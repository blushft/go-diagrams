package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type managementContainer struct {
	path  string
	attrs []attr.Attribute
}

var Management = &managementContainer{path: "assets/aws/management"}

func (c *managementContainer) Config(opts ...attr.Attribute) *node.Node {
	return node.New("config", attr.AssetImage("assets/aws/management/config.png"), attr.Shape(attr.None))
}

func (c *managementContainer) ManagedServices(opts ...attr.Attribute) *node.Node {
	return node.New("managed-services", attr.AssetImage("assets/aws/management/managed-services.png"), attr.Shape(attr.None))
}

func (c *managementContainer) Organizations(opts ...attr.Attribute) *node.Node {
	return node.New("organizations", attr.AssetImage("assets/aws/management/organizations.png"), attr.Shape(attr.None))
}

func (c *managementContainer) CommandLineInterface(opts ...attr.Attribute) *node.Node {
	return node.New("command-line-interface", attr.AssetImage("assets/aws/management/command-line-interface.png"), attr.Shape(attr.None))
}

func (c *managementContainer) Cloudtrail(opts ...attr.Attribute) *node.Node {
	return node.New("cloudtrail", attr.AssetImage("assets/aws/management/cloudtrail.png"), attr.Shape(attr.None))
}

func (c *managementContainer) ServiceCatalog(opts ...attr.Attribute) *node.Node {
	return node.New("service-catalog", attr.AssetImage("assets/aws/management/service-catalog.png"), attr.Shape(attr.None))
}

func (c *managementContainer) SystemsManagerParameterStore(opts ...attr.Attribute) *node.Node {
	return node.New("systems-manager-parameter-store", attr.AssetImage("assets/aws/management/systems-manager-parameter-store.png"), attr.Shape(attr.None))
}

func (c *managementContainer) TrustedAdvisor(opts ...attr.Attribute) *node.Node {
	return node.New("trusted-advisor", attr.AssetImage("assets/aws/management/trusted-advisor.png"), attr.Shape(attr.None))
}

func (c *managementContainer) WellArchitectedTool(opts ...attr.Attribute) *node.Node {
	return node.New("well-architected-tool", attr.AssetImage("assets/aws/management/well-architected-tool.png"), attr.Shape(attr.None))
}

func (c *managementContainer) AutoScaling(opts ...attr.Attribute) *node.Node {
	return node.New("auto-scaling", attr.AssetImage("assets/aws/management/auto-scaling.png"), attr.Shape(attr.None))
}

func (c *managementContainer) Cloudwatch(opts ...attr.Attribute) *node.Node {
	return node.New("cloudwatch", attr.AssetImage("assets/aws/management/cloudwatch.png"), attr.Shape(attr.None))
}

func (c *managementContainer) Codeguru(opts ...attr.Attribute) *node.Node {
	return node.New("codeguru", attr.AssetImage("assets/aws/management/codeguru.png"), attr.Shape(attr.None))
}

func (c *managementContainer) Cloudformation(opts ...attr.Attribute) *node.Node {
	return node.New("cloudformation", attr.AssetImage("assets/aws/management/cloudformation.png"), attr.Shape(attr.None))
}

func (c *managementContainer) LicenseManager(opts ...attr.Attribute) *node.Node {
	return node.New("license-manager", attr.AssetImage("assets/aws/management/license-manager.png"), attr.Shape(attr.None))
}

func (c *managementContainer) ManagementConsole(opts ...attr.Attribute) *node.Node {
	return node.New("management-console", attr.AssetImage("assets/aws/management/management-console.png"), attr.Shape(attr.None))
}

func (c *managementContainer) Opsworks(opts ...attr.Attribute) *node.Node {
	return node.New("opsworks", attr.AssetImage("assets/aws/management/opsworks.png"), attr.Shape(attr.None))
}

func (c *managementContainer) SystemsManager(opts ...attr.Attribute) *node.Node {
	return node.New("systems-manager", attr.AssetImage("assets/aws/management/systems-manager.png"), attr.Shape(attr.None))
}

func (c *managementContainer) ControlTower(opts ...attr.Attribute) *node.Node {
	return node.New("control-tower", attr.AssetImage("assets/aws/management/control-tower.png"), attr.Shape(attr.None))
}
