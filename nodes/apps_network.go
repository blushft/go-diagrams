package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type appsNetworkContainer struct {
	path  string
	attrs []attr.Attribute
}

var AppsNetwork = &appsNetworkContainer{path: "assets/apps/network"}

func (c *appsNetworkContainer) Consul(opts ...attr.Attribute) *node.Node {
	return node.New("consul", attr.AssetImage("assets/apps/network/consul.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Etcd(opts ...attr.Attribute) *node.Node {
	return node.New("etcd", attr.AssetImage("assets/apps/network/etcd.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Internet(opts ...attr.Attribute) *node.Node {
	return node.New("internet", attr.AssetImage("assets/apps/network/internet.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Istio(opts ...attr.Attribute) *node.Node {
	return node.New("istio", attr.AssetImage("assets/apps/network/istio.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Traefik(opts ...attr.Attribute) *node.Node {
	return node.New("traefik", attr.AssetImage("assets/apps/network/traefik.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Caddy(opts ...attr.Attribute) *node.Node {
	return node.New("caddy", attr.AssetImage("assets/apps/network/caddy.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Envoy(opts ...attr.Attribute) *node.Node {
	return node.New("envoy", attr.AssetImage("assets/apps/network/envoy.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Haproxy(opts ...attr.Attribute) *node.Node {
	return node.New("haproxy", attr.AssetImage("assets/apps/network/haproxy.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Nginx(opts ...attr.Attribute) *node.Node {
	return node.New("nginx", attr.AssetImage("assets/apps/network/nginx.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Ocelot(opts ...attr.Attribute) *node.Node {
	return node.New("ocelot", attr.AssetImage("assets/apps/network/ocelot.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Vyos(opts ...attr.Attribute) *node.Node {
	return node.New("vyos", attr.AssetImage("assets/apps/network/vyos.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Zookeeper(opts ...attr.Attribute) *node.Node {
	return node.New("zookeeper", attr.AssetImage("assets/apps/network/zookeeper.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Linkerd(opts ...attr.Attribute) *node.Node {
	return node.New("linkerd", attr.AssetImage("assets/apps/network/linkerd.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Tomcat(opts ...attr.Attribute) *node.Node {
	return node.New("tomcat", attr.AssetImage("assets/apps/network/tomcat.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Apache(opts ...attr.Attribute) *node.Node {
	return node.New("apache", attr.AssetImage("assets/apps/network/apache.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Kong(opts ...attr.Attribute) *node.Node {
	return node.New("kong", attr.AssetImage("assets/apps/network/kong.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Pfsense(opts ...attr.Attribute) *node.Node {
	return node.New("pfsense", attr.AssetImage("assets/apps/network/pfsense.png"), attr.Shape(attr.None))
}

func (c *appsNetworkContainer) Pomerium(opts ...attr.Attribute) *node.Node {
	return node.New("pomerium", attr.AssetImage("assets/apps/network/pomerium.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.network"
	Register(namespace, "Consul", AppsNetwork.Consul)
	Register(namespace, "Etcd", AppsNetwork.Etcd)
	Register(namespace, "Internet", AppsNetwork.Internet)
	Register(namespace, "Istio", AppsNetwork.Istio)
	Register(namespace, "Traefik", AppsNetwork.Traefik)
	Register(namespace, "Caddy", AppsNetwork.Caddy)
	Register(namespace, "Envoy", AppsNetwork.Envoy)
	Register(namespace, "Haproxy", AppsNetwork.Haproxy)
	Register(namespace, "Nginx", AppsNetwork.Nginx)
	Register(namespace, "Ocelot", AppsNetwork.Ocelot)
	Register(namespace, "Vyos", AppsNetwork.Vyos)
	Register(namespace, "Zookeeper", AppsNetwork.Zookeeper)
	Register(namespace, "Linkerd", AppsNetwork.Linkerd)
	Register(namespace, "Tomcat", AppsNetwork.Tomcat)
	Register(namespace, "Apache", AppsNetwork.Apache)
	Register(namespace, "Kong", AppsNetwork.Kong)
	Register(namespace, "Pfsense", AppsNetwork.Pfsense)
	Register(namespace, "Pomerium", AppsNetwork.Pomerium)
}