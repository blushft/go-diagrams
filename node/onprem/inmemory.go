package onprem

import "github.com/blushft/go-diagrams/node"

type inmemoryContainer struct {
	path string
	opts []node.Option
}

var Inmemory = &inmemoryContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/inmemory",
}

func (c *inmemoryContainer) Aerospike(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/inmemory/aerospike.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *inmemoryContainer) Hazelcast(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/inmemory/hazelcast.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *inmemoryContainer) Memcached(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/inmemory/memcached.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *inmemoryContainer) Redis(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/inmemory/redis.png")}, c.opts, opts)
	return node.New(nopts...)
}
