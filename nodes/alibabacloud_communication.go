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
	return node.New("mobile-push", attr.AssetImage("assets/alibabacloud/communication/mobile-push.png"), attr.Shape(attr.None))
}

func (c *communicationContainer) DirectMail(opts ...attr.Attribute) *node.Node {
	return node.New("direct-mail", attr.AssetImage("assets/alibabacloud/communication/direct-mail.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "alibabacloud.communication"
	Register(namespace, "MobilePush", AlibabaCommunication.MobilePush)
	Register(namespace, "DirectMail", AlibabaCommunication.DirectMail)

}
