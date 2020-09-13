package apps

import "github.com/blushft/go-diagrams/diagram"

type networkContainer struct {
	path string
	opts []diagram.NodeOption
}

var Network = &networkContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/network",
}

func (c *networkContainer) Internet(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/internet.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Istio(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/istio.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Linkerd(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/linkerd.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Ocelot(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/ocelot.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Tomcat(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/tomcat.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Vyos(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/vyos.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Envoy(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/envoy.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Nginx(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/nginx.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Pfsense(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/pfsense.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Apache(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/apache.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Etcd(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/etcd.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Haproxy(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/haproxy.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Pomerium(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/pomerium.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Caddy(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/caddy.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Kong(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/kong.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Traefik(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/traefik.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Zookeeper(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/zookeeper.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *networkContainer) Consul(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/network/consul.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
