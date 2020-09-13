package gcp

import "github.com/blushft/go-diagrams/diagram"

type iotContainer struct {
	path string
	opts []diagram.NodeOption
}

var Iot = &iotContainer{
	opts: diagram.OptionSet{diagram.Provider("gcp"), diagram.NodeShape("none")},
	path: "assets/gcp/iot",
}

func (c *iotContainer) IotCore(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/gcp/iot/iot-core.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
