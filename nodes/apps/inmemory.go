package apps

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type inmemoryContainer struct {
	path  string
	attrs []attr.Attribute
}

var Inmemory = &inmemoryContainer{path: "assets/apps/inmemory"}

func (c *inmemoryContainer) Aerospike(opts ...attr.Attribute) *node.Node {
	return node.New("aerospike", attr.AssetImage("assets/apps/inmemory/aerospike.png"), attr.Shape(attr.None))
}

func (c *inmemoryContainer) Hazelcast(opts ...attr.Attribute) *node.Node {
	return node.New("hazelcast", attr.AssetImage("assets/apps/inmemory/hazelcast.png"), attr.Shape(attr.None))
}

func (c *inmemoryContainer) Memcached(opts ...attr.Attribute) *node.Node {
	return node.New("memcached", attr.AssetImage("assets/apps/inmemory/memcached.png"), attr.Shape(attr.None))
}

func (c *inmemoryContainer) Redis(opts ...attr.Attribute) *node.Node {
	return node.New("redis", attr.AssetImage("assets/apps/inmemory/redis.png"), attr.Shape(attr.None))
}
