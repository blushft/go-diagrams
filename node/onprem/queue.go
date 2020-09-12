package onprem

import "github.com/blushft/go-diagrams/node"

type queueContainer struct {
	path string
	opts []node.Option
}

var Queue = &queueContainer{
	opts: node.OptionSet{node.Provider("onprem"), node.Shape("none")},
	path: "assets/onprem/queue",
}

func (c *queueContainer) Zeromq(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/queue/zeromq.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *queueContainer) Activemq(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/queue/activemq.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *queueContainer) Celery(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/queue/celery.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *queueContainer) Kafka(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/queue/kafka.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *queueContainer) Rabbitmq(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/onprem/queue/rabbitmq.png")}, c.opts, opts)
	return node.New(nopts...)
}
