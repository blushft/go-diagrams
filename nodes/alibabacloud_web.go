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
	return node.New("dns", attr.AssetImage("assets/alibabacloud/web/dns.png"), attr.Shape(attr.None))
}

func (c *webContainer) Domain(opts ...attr.Attribute) *node.Node {
	return node.New("domain", attr.AssetImage("assets/alibabacloud/web/domain.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "alibabacloud.web"
  Register(namespace, "Dns", AlibabacloudWeb.Dns)
  Register(namespace, "Domain", AlibabacloudWeb.Domain)
}
