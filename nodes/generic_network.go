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
	opts = append(opts, attr.AssetImage("assets/generic/network/firewall.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("firewall", opts...)
}

func (c *GenericNetworkContainer) Router(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/generic/network/router.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("router", opts...)
}

func (c *GenericNetworkContainer) Switch(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/generic/network/switch.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("switch", opts...)
}

func (c *GenericNetworkContainer) Vpn(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/generic/network/vpn.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vpn", opts...)
}

func init() {
	const namespace = "generic.network"
	Register(namespace, "Firewall", GenericNetwork.Firewall)
	Register(namespace, "Router", GenericNetwork.Router)
	Register(namespace, "Switch", GenericNetwork.Switch)
	Register(namespace, "Vpn", GenericNetwork.Vpn)
}
