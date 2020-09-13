package alibabacloud

import "github.com/blushft/go-diagrams/diagram"

type communicationContainer struct {
	path string
	opts []diagram.NodeOption
}

var Communication = &communicationContainer{
	opts: diagram.OptionSet{diagram.Provider("alibabacloud"), diagram.NodeShape("none")},
	path: "assets/alibabacloud/communication",
}

func (c *communicationContainer) DirectMail(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/communication/direct-mail.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *communicationContainer) MobilePush(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/alibabacloud/communication/mobile-push.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
