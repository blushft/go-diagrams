package gcp

import "github.com/blushft/go-diagrams/node"

type iotContainer struct {
	path string
	opts []node.Option
}

var Iot = &iotContainer{
	opts: node.OptionSet{node.Provider("gcp"), node.Shape("none")},
	path: "assets/gcp/iot",
}

func (c *iotContainer) IotCore(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/gcp/iot/iot-core.png")}, c.opts, opts)
	return node.New(nopts...)
}
