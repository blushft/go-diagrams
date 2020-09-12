package alibabacloud

import "github.com/blushft/go-diagrams/node"

type communicationContainer struct {
	path string
	opts []node.Option
}

var Communication = &communicationContainer{
	opts: node.OptionSet{node.Provider("alibabacloud"), node.Shape("none")},
	path: "assets/alibabacloud/communication",
}

func (c *communicationContainer) DirectMail(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/communication/direct-mail.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *communicationContainer) MobilePush(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/alibabacloud/communication/mobile-push.png")}, c.opts, opts)
	return node.New(nopts...)
}
