package apps

import "github.com/blushft/go-diagrams/diagram"

type queueContainer struct {
	path string
	opts []diagram.NodeOption
}

var Queue = &queueContainer{
	opts: diagram.OptionSet{diagram.Provider("apps"), diagram.NodeShape("none")},
	path: "assets/apps/queue",
}

func (c *queueContainer) Activemq(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/queue/activemq.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *queueContainer) Celery(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/queue/celery.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *queueContainer) Kafka(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/queue/kafka.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *queueContainer) Rabbitmq(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/queue/rabbitmq.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *queueContainer) Zeromq(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/apps/queue/zeromq.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
