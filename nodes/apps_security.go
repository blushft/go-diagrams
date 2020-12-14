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
	return node.New("trivy", attr.AssetImage("assets/apps/security/trivy.png"), attr.Shape(attr.None))
}

func (c *appsSecurityContainer) Vault(opts ...attr.Attribute) *node.Node {
	return node.New("vault", attr.AssetImage("assets/apps/security/vault.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.security"
	Register(namespace, "Trivy", Security.Trivy)
	Register(namespace, "Vault", Security.Vault)
}
