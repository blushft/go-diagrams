package programming

import "github.com/blushft/go-diagrams/node"

type languageContainer struct {
	path string
	opts []node.Option
}

var Language = &languageContainer{
	opts: node.OptionSet{node.Provider("programming"), node.Shape("none")},
	path: "assets/programming/language",
}

func (c *languageContainer) C(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/c.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) Python(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/python.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) Typescript(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/typescript.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) Bash(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/bash.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) Dart(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/dart.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) Javascript(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/javascript.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) Matlab(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/matlab.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) R(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/r.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) Rust(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/rust.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) Kotlin(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/kotlin.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) Php(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/php.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) Ruby(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/ruby.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) Cpp(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/cpp.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) Csharp(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/csharp.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) Go(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/go.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) Java(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/java.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) Nodejs(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/nodejs.png")}, c.opts, opts)
	return node.New(nopts...)
}

func (c *languageContainer) Swift(opts ...node.Option) *node.Node {
	nopts := node.MergeOptionSets(node.OptionSet{node.Icon("assets/programming/language/swift.png")}, c.opts, opts)
	return node.New(nopts...)
}
