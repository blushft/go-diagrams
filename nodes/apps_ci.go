package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type ciContainer struct {
	path  string
	attrs []attr.Attribute
}

var Ci = &ciContainer{path: "assets/apps/ci"}

func (c *ciContainer) Travisci(opts ...attr.Attribute) *node.Node {
	return node.New("travisci", attr.AssetImage("assets/apps/ci/travisci.png"), attr.Shape(attr.None))
}

func (c *ciContainer) Zuulci(opts ...attr.Attribute) *node.Node {
	return node.New("zuulci", attr.AssetImage("assets/apps/ci/zuulci.png"), attr.Shape(attr.None))
}

func (c *ciContainer) Circleci(opts ...attr.Attribute) *node.Node {
	return node.New("circleci", attr.AssetImage("assets/apps/ci/circleci.png"), attr.Shape(attr.None))
}

func (c *ciContainer) Concourseci(opts ...attr.Attribute) *node.Node {
	return node.New("concourseci", attr.AssetImage("assets/apps/ci/concourseci.png"), attr.Shape(attr.None))
}

func (c *ciContainer) Droneci(opts ...attr.Attribute) *node.Node {
	return node.New("droneci", attr.AssetImage("assets/apps/ci/droneci.png"), attr.Shape(attr.None))
}

func (c *ciContainer) Gitlabci(opts ...attr.Attribute) *node.Node {
	return node.New("gitlabci", attr.AssetImage("assets/apps/ci/gitlabci.png"), attr.Shape(attr.None))
}

func (c *ciContainer) Jenkins(opts ...attr.Attribute) *node.Node {
	return node.New("jenkins", attr.AssetImage("assets/apps/ci/jenkins.png"), attr.Shape(attr.None))
}

func (c *ciContainer) Teamcity(opts ...attr.Attribute) *node.Node {
	return node.New("teamcity", attr.AssetImage("assets/apps/ci/teamcity.png"), attr.Shape(attr.None))
}

func init() {
	const namespace = "apps.ci"
	Register(namespace, "Travisci", Ci.Travisci)
	Register(namespace, "Zuulci", Ci.Zuulci)
	Register(namespace, "Circleci", Ci.Circleci)
	Register(namespace, "Concourseci", Ci.Concourseci)
	Register(namespace, "Droneci", Ci.Droneci)
	Register(namespace, "Gitlabci", Ci.Gitlabci)
	Register(namespace, "Jenkins", Ci.Jenkins)
	Register(namespace, "Teamcity", Ci.Teamcity)
}
