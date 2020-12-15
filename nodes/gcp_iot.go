package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type iotContainer struct {
	path  string
	attrs []attr.Attribute
}

var GcpIot = &iotContainer{path: "assets/gcp/iot"}

func (c *iotContainer) IotCore(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/iot/iot-core.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("iot-core", opts...)
}

func init(){
	const namespace = "gcp.iot"
	Register(namespace, "IotCore", GcpIot.IotCore)
}