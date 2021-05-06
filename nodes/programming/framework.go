package programming

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type frameworkContainer struct {
	path  string
	attrs []attr.Attribute
}

var Framework = &frameworkContainer{path: "assets/programming/framework"}

func (c *frameworkContainer) Ember(opts ...attr.Attribute) *node.Node {
	return node.New("ember", attr.AssetImage("assets/programming/framework/ember.png"), attr.Shape(attr.None))
}

func (c *frameworkContainer) Flask(opts ...attr.Attribute) *node.Node {
	return node.New("flask", attr.AssetImage("assets/programming/framework/flask.png"), attr.Shape(attr.None))
}

func (c *frameworkContainer) Flutter(opts ...attr.Attribute) *node.Node {
	return node.New("flutter", attr.AssetImage("assets/programming/framework/flutter.png"), attr.Shape(attr.None))
}

func (c *frameworkContainer) Laravel(opts ...attr.Attribute) *node.Node {
	return node.New("laravel", attr.AssetImage("assets/programming/framework/laravel.png"), attr.Shape(attr.None))
}

func (c *frameworkContainer) Spring(opts ...attr.Attribute) *node.Node {
	return node.New("spring", attr.AssetImage("assets/programming/framework/spring.png"), attr.Shape(attr.None))
}

func (c *frameworkContainer) Vue(opts ...attr.Attribute) *node.Node {
	return node.New("vue", attr.AssetImage("assets/programming/framework/vue.png"), attr.Shape(attr.None))
}

func (c *frameworkContainer) Angular(opts ...attr.Attribute) *node.Node {
	return node.New("angular", attr.AssetImage("assets/programming/framework/angular.png"), attr.Shape(attr.None))
}

func (c *frameworkContainer) Backbone(opts ...attr.Attribute) *node.Node {
	return node.New("backbone", attr.AssetImage("assets/programming/framework/backbone.png"), attr.Shape(attr.None))
}

func (c *frameworkContainer) React(opts ...attr.Attribute) *node.Node {
	return node.New("react", attr.AssetImage("assets/programming/framework/react.png"), attr.Shape(attr.None))
}

func (c *frameworkContainer) Django(opts ...attr.Attribute) *node.Node {
	return node.New("django", attr.AssetImage("assets/programming/framework/django.png"), attr.Shape(attr.None))
}

func (c *frameworkContainer) Rails(opts ...attr.Attribute) *node.Node {
	return node.New("rails", attr.AssetImage("assets/programming/framework/rails.png"), attr.Shape(attr.None))
}
