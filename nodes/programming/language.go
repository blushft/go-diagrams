package programming

import "github.com/blushft/go-diagrams/diagram"

type languageContainer struct {
	path string
	opts []diagram.NodeOption
}

var Language = &languageContainer{
	opts: diagram.OptionSet{diagram.Provider("programming"), diagram.NodeShape("none")},
	path: "assets/programming/language",
}

func (c *languageContainer) Typescript(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/typescript.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) Bash(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/bash.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) Csharp(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/csharp.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) Matlab(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/matlab.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) Nodejs(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/nodejs.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) Php(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/php.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) Cpp(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/cpp.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) Go(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/go.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) R(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/r.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) Rust(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/rust.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) Swift(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/swift.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) C(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/c.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) Javascript(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/javascript.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) Kotlin(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/kotlin.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) Python(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/python.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) Ruby(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/ruby.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) Dart(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/dart.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *languageContainer) Java(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/programming/language/java.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
