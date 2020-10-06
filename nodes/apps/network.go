package apps

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type networkContainer struct {
	path  string
	attrs []attr.Attribute
}

var Network = &networkContainer{path: "assets/apps/network"}

func (c *networkContainer) Consul(opts ...attr.Attribute) *node.Node {
	return node.New("consul", attr.AssetImage("assets/apps/network/consul.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Etcd(opts ...attr.Attribute) *node.Node {
	return node.New("etcd", attr.AssetImage("assets/apps/network/etcd.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Internet(opts ...attr.Attribute) *node.Node {
	return node.New("internet", attr.AssetImage("assets/apps/network/internet.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Istio(opts ...attr.Attribute) *node.Node {
	return node.New("istio", attr.AssetImage("assets/apps/network/istio.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Traefik(opts ...attr.Attribute) *node.Node {
	return node.New("traefik", attr.AssetImage("assets/apps/network/traefik.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Caddy(opts ...attr.Attribute) *node.Node {
	return node.New("caddy", attr.AssetImage("assets/apps/network/caddy.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Envoy(opts ...attr.Attribute) *node.Node {
	return node.New("envoy", attr.AssetImage("assets/apps/network/envoy.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Haproxy(opts ...attr.Attribute) *node.Node {
	return node.New("haproxy", attr.AssetImage("assets/apps/network/haproxy.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Nginx(opts ...attr.Attribute) *node.Node {
	return node.New("nginx", attr.AssetImage("assets/apps/network/nginx.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Ocelot(opts ...attr.Attribute) *node.Node {
	return node.New("ocelot", attr.AssetImage("assets/apps/network/ocelot.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Vyos(opts ...attr.Attribute) *node.Node {
	return node.New("vyos", attr.AssetImage("assets/apps/network/vyos.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Zookeeper(opts ...attr.Attribute) *node.Node {
	return node.New("zookeeper", attr.AssetImage("assets/apps/network/zookeeper.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Linkerd(opts ...attr.Attribute) *node.Node {
	return node.New("linkerd", attr.AssetImage("assets/apps/network/linkerd.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Tomcat(opts ...attr.Attribute) *node.Node {
	return node.New("tomcat", attr.AssetImage("assets/apps/network/tomcat.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Apache(opts ...attr.Attribute) *node.Node {
	return node.New("apache", attr.AssetImage("assets/apps/network/apache.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Kong(opts ...attr.Attribute) *node.Node {
	return node.New("kong", attr.AssetImage("assets/apps/network/kong.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Pfsense(opts ...attr.Attribute) *node.Node {
	return node.New("pfsense", attr.AssetImage("assets/apps/network/pfsense.png"), attr.Shape(attr.None))
}

func (c *networkContainer) Pomerium(opts ...attr.Attribute) *node.Node {
	return node.New("pomerium", attr.AssetImage("assets/apps/network/pomerium.png"), attr.Shape(attr.None))
}
