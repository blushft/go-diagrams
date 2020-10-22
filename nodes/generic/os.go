package generic

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/node"
)

type osContainer struct {
	path  string
	attrs []attr.Attribute
}

var Os = &osContainer{path: "assets/generic/os"}

func (c *osContainer) LinuxGeneral(opts ...attr.Attribute) *node.Node {
	return node.New("linux-general", attr.AssetImage("assets/generic/os/linux-general.png"), attr.Shape(attr.None))
}

func (c *osContainer) Suse(opts ...attr.Attribute) *node.Node {
	return node.New("suse", attr.AssetImage("assets/generic/os/suse.png"), attr.Shape(attr.None))
}

func (c *osContainer) Ubuntu(opts ...attr.Attribute) *node.Node {
	return node.New("ubuntu", attr.AssetImage("assets/generic/os/ubuntu.png"), attr.Shape(attr.None))
}

func (c *osContainer) Windows(opts ...attr.Attribute) *node.Node {
	return node.New("windows", attr.AssetImage("assets/generic/os/windows.png"), attr.Shape(attr.None))
}

func (c *osContainer) Android(opts ...attr.Attribute) *node.Node {
	return node.New("android", attr.AssetImage("assets/generic/os/android.png"), attr.Shape(attr.None))
}

func (c *osContainer) Centos(opts ...attr.Attribute) *node.Node {
	return node.New("centos", attr.AssetImage("assets/generic/os/centos.png"), attr.Shape(attr.None))
}

func (c *osContainer) Ios(opts ...attr.Attribute) *node.Node {
	return node.New("ios", attr.AssetImage("assets/generic/os/ios.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "generic.os"
	diagram.Register(namespace, "LinuxGeneral", Os.LinuxGeneral)
	diagram.Register(namespace, "Suse", Os.Suse)
	diagram.Register(namespace, "Ubuntu", Os.Ubuntu)
	diagram.Register(namespace, "Windows", Os.Windows)
	diagram.Register(namespace, "Android", Os.Android)
	diagram.Register(namespace, "Centos", Os.Centos)
	diagram.Register(namespace, "Ios", Os.Ios)
}