package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type queueContainer struct {
	path  string
	attrs []attr.Attribute
}

var Queue = &queueContainer{path: "assets/apps/queue"}

func (c *queueContainer) Activemq(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/queue/activemq.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("activemq", opts...)
}

func (c *queueContainer) Celery(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/queue/celery.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("celery", opts...)
}

func (c *queueContainer) Kafka(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/queue/kafka.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("kafka", opts...)
}

func (c *queueContainer) Rabbitmq(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/queue/rabbitmq.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("rabbitmq", opts...)
}

func (c *queueContainer) Zeromq(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/queue/zeromq.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("zeromq", opts...)
}

func init() {
	const namespace = "apps.queue"
	Register(namespace, "Activemq", Queue.Activemq)
	Register(namespace, "Celery", Queue.Celery)
	Register(namespace, "Kafka", Queue.Kafka)
	Register(namespace, "Rabbitmq", Queue.Rabbitmq)
	Register(namespace, "Zeromq", Queue.Zeromq)
}