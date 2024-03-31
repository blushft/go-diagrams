package programming

import "github.com/blushft/go-diagrams/diagram"

type frameworkContainer struct {
	path string
	opts []diagram.NodeOption
}

var Framework = &frameworkContainer{
	opts: diagram.OptionSet{diagram.Provider("programming"), diagram.NodeShape("none")},
	path: "assets/programming/framework",
}

func (c *frameworkContainer) Django(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/framework/django.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *frameworkContainer) Flutter(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/framework/flutter.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *frameworkContainer) React(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/framework/react.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *frameworkContainer) Laravel(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/framework/laravel.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *frameworkContainer) Rails(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/framework/rails.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *frameworkContainer) Spring(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/framework/spring.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *frameworkContainer) Vue(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/framework/vue.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *frameworkContainer) Angular(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/framework/angular.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *frameworkContainer) Backbone(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/framework/backbone.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *frameworkContainer) Ember(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/framework/ember.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *frameworkContainer) Flask(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/framework/flask.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
