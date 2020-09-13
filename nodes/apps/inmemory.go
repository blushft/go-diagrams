package apps

import "github.com/blushft/go-diagrams/diagram"

type inmemoryContainer struct {
	path string
	opts []diagram.NodeOption
}

var Inmemory = &inmemoryContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/inmemory",
}

func (c *inmemoryContainer) Hazelcast(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/inmemory/hazelcast.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *inmemoryContainer) Memcached(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/inmemory/memcached.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *inmemoryContainer) Redis(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/inmemory/redis.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *inmemoryContainer) Aerospike(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/inmemory/aerospike.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
