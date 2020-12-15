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
	opts = append(opts, attr.AssetImage("assets/apps/network/consul.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("consul", opts...)
}

func (c *appsNetworkContainer) Etcd(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/etcd.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("etcd", opts...)
}

func (c *appsNetworkContainer) Internet(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/internet.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("internet", opts...)
}

func (c *appsNetworkContainer) Istio(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/istio.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("istio", opts...)
}

func (c *appsNetworkContainer) Traefik(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/traefik.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("traefik", opts...)
}

func (c *appsNetworkContainer) Caddy(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/caddy.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("caddy", opts...)
}

func (c *appsNetworkContainer) Envoy(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/envoy.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("envoy", opts...)
}

func (c *appsNetworkContainer) Haproxy(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/haproxy.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("haproxy", opts...)
}

func (c *appsNetworkContainer) Nginx(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/nginx.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("nginx", opts...)
}

func (c *appsNetworkContainer) Ocelot(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/ocelot.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ocelot", opts...)
}

func (c *appsNetworkContainer) Vyos(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/vyos.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vyos", opts...)
}

func (c *appsNetworkContainer) Zookeeper(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/zookeeper.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("zookeeper", opts...)
}

func (c *appsNetworkContainer) Linkerd(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/linkerd.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("linkerd", opts...)
}

func (c *appsNetworkContainer) Tomcat(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/tomcat.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("tomcat", opts...)
}

func (c *appsNetworkContainer) Apache(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/apache.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("apache", opts...)
}

func (c *appsNetworkContainer) Kong(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/kong.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("kong", opts...)
}

func (c *appsNetworkContainer) Pfsense(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/pfsense.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("pfsense", opts...)
}

func (c *appsNetworkContainer) Pomerium(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/network/pomerium.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("pomerium", opts...)
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