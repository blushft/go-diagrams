package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type appsCiContainer struct {
	path  string
	attrs []attr.Attribute
}

var Ci = &appsCiContainer{path: "assets/apps/ci"}

func (c *appsCiContainer) Travisci(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/ci/travisci.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("travisci", opts...)
}

func (c *appsCiContainer) Zuulci(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/ci/zuulci.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("zuulci", opts...)
}

func (c *appsCiContainer) Circleci(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/ci/circleci.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("circleci", opts...)
}

func (c *appsCiContainer) Concourseci(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/ci/concourseci.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("concourseci", opts...)
}

func (c *appsCiContainer) Droneci(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/ci/droneci.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("droneci", opts...)
}

func (c *appsCiContainer) Gitlabci(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/ci/gitlabci.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("gitlabci", opts...)
}

func (c *appsCiContainer) Jenkins(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/ci/jenkins.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("jenkins", opts...)
}

func (c *appsCiContainer) Teamcity(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/apps/ci/teamcity.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("teamcity", opts...)
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
