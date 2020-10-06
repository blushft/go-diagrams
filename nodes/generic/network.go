package generic

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type networkContainer struct {
	path  string
	attrs []attr.Attribute
}

var Network = &networkContainer{path: "assets/generic/network"}

func (c *networkContainer) Firewall(opts ...attr.Attribute) *node.Node {
	return node.New("firewall", attr.AssetImage("assets/generic/network/firewall.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Router(opts ...attr.Attribute) *node.Node {
	return node.New("router", attr.AssetImage("assets/generic/network/router.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Switch(opts ...attr.Attribute) *node.Node {
	return node.New("switch", attr.AssetImage("assets/generic/network/switch.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Vpn(opts ...attr.Attribute) *node.Node {
	return node.New("vpn", attr.AssetImage("assets/generic/network/vpn.png"), attr.Shape(attr.None))
}
