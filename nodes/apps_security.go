package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type appsSecurityContainer struct {
	path  string
	attrs []attr.Attribute
}

var Security = &appsSecurityContainer{path: "assets/apps/security"}

func (c *appsSecurityContainer) Trivy(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/security/trivy.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("trivy", opts...)
}

func (c *appsSecurityContainer) Vault(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/security/vault.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vault", opts...)
}

func init() {
	const namespace = "apps.security"
	Register(namespace, "Trivy", Security.Trivy)
	Register(namespace, "Vault", Security.Vault)
}
