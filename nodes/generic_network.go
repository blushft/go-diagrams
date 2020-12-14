package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type GenericNetworkContainer struct {
	path  string
	attrs []attr.Attribute
}

var GenericNetwork = &GenericNetworkContainer{path: "assets/generic/network"}

func (c *GenericNetworkContainer) Firewall(opts ...attr.Attribute) *node.Node {
	return node.New("firewall", attr.AssetImage("assets/generic/network/firewall.png"), attr.Shape(attr.None))
}

func (c *GenericNetworkContainer) Router(opts ...attr.Attribute) *node.Node {
	return node.New("router", attr.AssetImage("assets/generic/network/router.png"), attr.Shape(attr.None))
}

func (c *GenericNetworkContainer) Switch(opts ...attr.Attribute) *node.Node {
	return node.New("switch", attr.AssetImage("assets/generic/network/switch.png"), attr.Shape(attr.None))
}

func (c *GenericNetworkContainer) Vpn(opts ...attr.Attribute) *node.Node {
	return node.New("vpn", attr.AssetImage("assets/generic/network/vpn.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "generic.network"
	Register(namespace, "Firewall", GenericNetwork.Firewall)
	Register(namespace, "Router", GenericNetwork.Router)
	Register(namespace, "Switch", GenericNetwork.Switch)
	Register(namespace, "Vpn", GenericNetwork.Vpn)
}
