package programming

import "github.com/blushft/go-diagrams/node"

type frameworkContainer struct {
	path string
	opts []node.Option
}

var Framework = &frameworkContainer{
	opts: node.OptionSet{node.Provider("programming"), node.Shape("none")},
	path: "assets/programming/framework",
}

func (c *frameworkContainer) Spring(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/framework/spring.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *frameworkContainer) Vue(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/framework/vue.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *frameworkContainer) Flutter(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/framework/flutter.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *frameworkContainer) Laravel(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/framework/laravel.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *frameworkContainer) Rails(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/framework/rails.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *frameworkContainer) Angular(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/framework/angular.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *frameworkContainer) Backbone(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/framework/backbone.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *frameworkContainer) Django(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/framework/django.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *frameworkContainer) Ember(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/framework/ember.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *frameworkContainer) Flask(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/framework/flask.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *frameworkContainer) React(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/framework/react.png")}, c.opts, opts)
	return node.New(nopts...)
}
