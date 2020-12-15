package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type communicationContainer struct {
	path  string
	attrs []attr.Attribute
}

var AlibabaCommunication = &communicationContainer{path: "assets/alibabacloud/communication"}

func (c *communicationContainer) MobilePush(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/communication/mobile-push.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("mobile-push", opts...)
}

func (c *communicationContainer) DirectMail(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/communication/direct-mail.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("direct-mail", opts...)
}

func init() {
	const namespace = "alibabacloud.communication"
	Register(namespace, "MobilePush", AlibabaCommunication.MobilePush)
	Register(namespace, "DirectMail", AlibabaCommunication.DirectMail)

}
