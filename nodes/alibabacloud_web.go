package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type webContainer struct {
	path  string
	attrs []attr.Attribute
}

var AlibabacloudWeb =&webContainer{path: "assets/alibabacloud/web"}

func (c *webContainer) Dns(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/web/dns.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dns", opts...)
}

func (c *webContainer) Domain(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/alibabacloud/web/domain.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("domain", opts...)
}

func init() {
  const namespace = "alibabacloud.web"
  Register(namespace, "Dns", AlibabacloudWeb.Dns)
  Register(namespace, "Domain", AlibabacloudWeb.Domain)
}
