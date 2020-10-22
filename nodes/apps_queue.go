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
	return node.New("activemq", attr.AssetImage("assets/apps/queue/activemq.png"), attr.Shape(attr.None))
}

func (c *queueContainer) Celery(opts ...attr.Attribute) *node.Node {
	return node.New("celery", attr.AssetImage("assets/apps/queue/celery.png"), attr.Shape(attr.None))
}

func (c *queueContainer) Kafka(opts ...attr.Attribute) *node.Node {
	return node.New("kafka", attr.AssetImage("assets/apps/queue/kafka.png"), attr.Shape(attr.None))
}

func (c *queueContainer) Rabbitmq(opts ...attr.Attribute) *node.Node {
	return node.New("rabbitmq", attr.AssetImage("assets/apps/queue/rabbitmq.png"), attr.Shape(attr.None))
}

func (c *queueContainer) Zeromq(opts ...attr.Attribute) *node.Node {
	return node.New("zeromq", attr.AssetImage("assets/apps/queue/zeromq.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.queue"
	Register(namespace, "Activemq", Queue.Activemq)
	Register(namespace, "Celery", Queue.Celery)
	Register(namespace, "Kafka", Queue.Kafka)
	Register(namespace, "Rabbitmq", Queue.Rabbitmq)
	Register(namespace, "Zeromq", Queue.Zeromq)
}