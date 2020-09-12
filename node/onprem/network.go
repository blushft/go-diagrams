package onprem

import "github.com/blushft/go-diagrams/node"

type networkContainer struct {
	path string
	opts []node.Option
}

var Network = &networkContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/network",
}

func (c *networkContainer) Pfsense(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/pfsense.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Vyos(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/vyos.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Caddy(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/caddy.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Haproxy(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/haproxy.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Internet(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/internet.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Ocelot(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/ocelot.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Linkerd(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/linkerd.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Nginx(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/nginx.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Zookeeper(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/zookeeper.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Tomcat(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/tomcat.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Envoy(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/envoy.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Etcd(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/etcd.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Istio(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/istio.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Kong(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/kong.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Apache(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/apache.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Consul(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/consul.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Pomerium(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/pomerium.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *networkContainer) Traefik(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/network/traefik.png")}, c.opts, opts)
	return node.New(nopts...)
}
