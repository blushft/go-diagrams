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
	opts = append(opts, attr.AssetImage("assets/generic/os/linux-general.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("linux-general", opts...)
}

func (c *genericOSContainer) Suse(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/generic/os/suse.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("suse", opts...)
}

func (c *genericOSContainer) Ubuntu(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/generic/os/ubuntu.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ubuntu", opts...)
}

func (c *genericOSContainer) Windows(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/generic/os/windows.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("windows", opts...)
}

func (c *genericOSContainer) Android(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/generic/os/android.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("android", opts...)
}

func (c *genericOSContainer) Centos(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/generic/os/centos.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("centos", opts...)
}

func (c *genericOSContainer) Ios(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/generic/os/ios.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ios", opts...)
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