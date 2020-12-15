package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type managementContainer struct {
	path  string
	attrs []attr.Attribute
}

var Management = &managementContainer{path: "assets/aws/management"}

func (c *managementContainer) Cloudtrail(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/cloudtrail.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloudtrail", opts...)
}

func (c *managementContainer) Organizations(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/organizations.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("organizations", opts...)
}

func (c *managementContainer) SystemsManager(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/systems-manager.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("systems-manager", opts...)
}

func (c *managementContainer) Cloudwatch(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/cloudwatch.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloudwatch", opts...)
}

func (c *managementContainer) LicenseManager(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/license-manager.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("license-manager", opts...)
}

func (c *managementContainer) WellArchitectedTool(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/well-architected-tool.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("well-architected-tool", opts...)
}

func (c *managementContainer) ManagedServices(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/managed-services.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("managed-services", opts...)
}

func (c *managementContainer) Cloudformation(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/cloudformation.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cloudformation", opts...)
}

func (c *managementContainer) Codeguru(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/codeguru.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("codeguru", opts...)
}

func (c *managementContainer) Config(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/config.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("config", opts...)
}

func (c *managementContainer) ManagementConsole(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/management-console.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("management-console", opts...)
}

func (c *managementContainer) Opsworks(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/opsworks.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("opsworks", opts...)
}

func (c *managementContainer) ServiceCatalog(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/service-catalog.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("service-catalog", opts...)
}

func (c *managementContainer) SystemsManagerParameterStore(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/systems-manager-parameter-store.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("systems-manager-parameter-store", opts...)
}

func (c *managementContainer) TrustedAdvisor(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/trusted-advisor.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("trusted-advisor", opts...)
}

func (c *managementContainer) AutoScaling(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/auto-scaling.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("auto-scaling", opts...)
}

func (c *managementContainer) CommandLineInterface(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/command-line-interface.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("command-line-interface", opts...)
}

func (c *managementContainer) ControlTower(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/aws/management/control-tower.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("control-tower", opts...)
}

func init() {
  const namespace= "aws.management"
  Register(namespace, "Cloudtrail", Management.Cloudtrail)
  Register(namespace, "Organizations", Management.Organizations)
  Register(namespace, "SystemsManager", Management.SystemsManager)
  Register(namespace, "Cloudwatch", Management.Cloudwatch)
  Register(namespace, "LicenseManager", Management.LicenseManager)
  Register(namespace, "WellArchitectedTool", Management.WellArchitectedTool)
  Register(namespace, "ManagedServices", Management.ManagedServices)
  Register(namespace, "Cloudformation", Management.Cloudformation)
  Register(namespace, "Codeguru", Management.Codeguru)
  Register(namespace, "Config", Management.Config)
  Register(namespace, "ManagementConsole", Management.ManagementConsole)
  Register(namespace, "Opsworks", Management.Opsworks)
  Register(namespace, "ServiceCatalog", Management.ServiceCatalog)
  Register(namespace, "SystemsManagerParameterStore", Management.SystemsManagerParameterStore)
  Register(namespace, "TrustedAdvisor", Management.TrustedAdvisor)
  Register(namespace, "AutoScaling", Management.AutoScaling)
  Register(namespace, "CommandLineInterface", Management.CommandLineInterface)
  Register(namespace, "ControlTower", Management.ControlTower)
}
