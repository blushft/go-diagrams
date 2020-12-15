package nodes

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
	opts = append(opts, attr.AssetImage("assets/apps/inmemory/aerospike.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("aerospike", opts...)
}

func (c *inmemoryContainer) Hazelcast(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/inmemory/hazelcast.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("hazelcast", opts...)
}

func (c *inmemoryContainer) Memcached(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/inmemory/memcached.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("memcached", opts...)
}

func (c *inmemoryContainer) Redis(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/inmemory/redis.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("redis", opts...)
}

func init() {
	const namespace = "apps.inmemory"
	Register(namespace, "Aerospike", Inmemory.Aerospike)
	Register(namespace, "Hazelcast", Inmemory.Hazelcast)
	Register(namespace, "Memcached", Inmemory.Memcached)
	Register(namespace, "Redis", Inmemory.Redis)
}