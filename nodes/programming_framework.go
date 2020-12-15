package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type frameworkContainer struct {
	path  string
	attrs []attr.Attribute
}

var Framework = &frameworkContainer{path: "assets/programming/framework"}

func (c *frameworkContainer) Rails(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/framework/rails.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("rails", opts...)
}

func (c *frameworkContainer) Spring(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/framework/spring.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("spring", opts...)
}

func (c *frameworkContainer) Angular(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/framework/angular.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("angular", opts...)
}

func (c *frameworkContainer) Backbone(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/framework/backbone.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("backbone", opts...)
}

func (c *frameworkContainer) Ember(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/framework/ember.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ember", opts...)
}

func (c *frameworkContainer) Flask(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/framework/flask.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("flask", opts...)
}

func (c *frameworkContainer) Laravel(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/framework/laravel.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("laravel", opts...)
}

func (c *frameworkContainer) Django(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/framework/django.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("django", opts...)
}

func (c *frameworkContainer) Flutter(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/framework/flutter.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("flutter", opts...)
}

func (c *frameworkContainer) React(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/framework/react.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("react", opts...)
}

func (c *frameworkContainer) Vue(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/framework/vue.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("vue", opts...)
}

func init() {
	const namespace = "programming.framework"
	Register(namespace, "Rails", Framework.Rails)
	Register(namespace, "Spring", Framework.Spring)
	Register(namespace, "Angular", Framework.Angular)
	Register(namespace, "Backbone", Framework.Backbone)
	Register(namespace, "Ember", Framework.Ember)
	Register(namespace, "Flask", Framework.Flask)
	Register(namespace, "Laravel", Framework.Laravel)
	Register(namespace, "Django", Framework.Django)
	Register(namespace, "Flutter", Framework.Flutter)
	Register(namespace, "React", Framework.React)
	Register(namespace, "Vue", Framework.Vue)
}