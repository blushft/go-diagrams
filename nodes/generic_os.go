package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type genericOSContainer struct {
	path  string
	attrs []attr.Attribute
}

var GenericOS = &genericOSContainer{path: "assets/generic/os"}

func (c *genericOSContainer) LinuxGeneral(opts ...attr.Attribute) *node.Node {
	return node.New("linux-general", attr.AssetImage("assets/generic/os/linux-general.png"), attr.Shape(attr.None))
}

func (c *genericOSContainer) Suse(opts ...attr.Attribute) *node.Node {
	return node.New("suse", attr.AssetImage("assets/generic/os/suse.png"), attr.Shape(attr.None))
}

func (c *genericOSContainer) Ubuntu(opts ...attr.Attribute) *node.Node {
	return node.New("ubuntu", attr.AssetImage("assets/generic/os/ubuntu.png"), attr.Shape(attr.None))
}

func (c *genericOSContainer) Windows(opts ...attr.Attribute) *node.Node {
	return node.New("windows", attr.AssetImage("assets/generic/os/windows.png"), attr.Shape(attr.None))
}

func (c *genericOSContainer) Android(opts ...attr.Attribute) *node.Node {
	return node.New("android", attr.AssetImage("assets/generic/os/android.png"), attr.Shape(attr.None))
}

func (c *genericOSContainer) Centos(opts ...attr.Attribute) *node.Node {
	return node.New("centos", attr.AssetImage("assets/generic/os/centos.png"), attr.Shape(attr.None))
}

func (c *genericOSContainer) Ios(opts ...attr.Attribute) *node.Node {
	return node.New("ios", attr.AssetImage("assets/generic/os/ios.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "generic.os"
	Register(namespace, "LinuxGeneral", GenericOS.LinuxGeneral)
	Register(namespace, "Suse", GenericOS.Suse)
	Register(namespace, "Ubuntu", GenericOS.Ubuntu)
	Register(namespace, "Windows", GenericOS.Windows)
	Register(namespace, "Android", GenericOS.Android)
	Register(namespace, "Centos", GenericOS.Centos)
	Register(namespace, "Ios", GenericOS.Ios)
}