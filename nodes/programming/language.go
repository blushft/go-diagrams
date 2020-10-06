package programming

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type languageContainer struct {
	path  string
	attrs []attr.Attribute
}

var Language = &languageContainer{path: "assets/programming/language"}

func (c *languageContainer) Bash(opts ...attr.Attribute) *node.Node {
	return node.New("bash", attr.AssetImage("assets/programming/language/bash.png"), attr.Shape(attr.None))
}

func (c *languageContainer) Nodejs(opts ...attr.Attribute) *node.Node {
	return node.New("nodejs", attr.AssetImage("assets/programming/language/nodejs.png"), attr.Shape(attr.None))
}

func (c *languageContainer) Matlab(opts ...attr.Attribute) *node.Node {
	return node.New("matlab", attr.AssetImage("assets/programming/language/matlab.png"), attr.Shape(attr.None))
}

func (c *languageContainer) Python(opts ...attr.Attribute) *node.Node {
	return node.New("python", attr.AssetImage("assets/programming/language/python.png"), attr.Shape(attr.None))
}

func (c *languageContainer) Rust(opts ...attr.Attribute) *node.Node {
	return node.New("rust", attr.AssetImage("assets/programming/language/rust.png"), attr.Shape(attr.None))
}

func (c *languageContainer) Swift(opts ...attr.Attribute) *node.Node {
	return node.New("swift", attr.AssetImage("assets/programming/language/swift.png"), attr.Shape(attr.None))
}

func (c *languageContainer) Typescript(opts ...attr.Attribute) *node.Node {
	return node.New("typescript", attr.AssetImage("assets/programming/language/typescript.png"), attr.Shape(attr.None))
}

func (c *languageContainer) Csharp(opts ...attr.Attribute) *node.Node {
	return node.New("csharp", attr.AssetImage("assets/programming/language/csharp.png"), attr.Shape(attr.None))
}

func (c *languageContainer) Java(opts ...attr.Attribute) *node.Node {
	return node.New("java", attr.AssetImage("assets/programming/language/java.png"), attr.Shape(attr.None))
}

func (c *languageContainer) Javascript(opts ...attr.Attribute) *node.Node {
	return node.New("javascript", attr.AssetImage("assets/programming/language/javascript.png"), attr.Shape(attr.None))
}

func (c *languageContainer) Kotlin(opts ...attr.Attribute) *node.Node {
	return node.New("kotlin", attr.AssetImage("assets/programming/language/kotlin.png"), attr.Shape(attr.None))
}

func (c *languageContainer) C(opts ...attr.Attribute) *node.Node {
	return node.New("c", attr.AssetImage("assets/programming/language/c.png"), attr.Shape(attr.None))
}

func (c *languageContainer) Go(opts ...attr.Attribute) *node.Node {
	return node.New("go", attr.AssetImage("assets/programming/language/go.png"), attr.Shape(attr.None))
}

func (c *languageContainer) Php(opts ...attr.Attribute) *node.Node {
	return node.New("php", attr.AssetImage("assets/programming/language/php.png"), attr.Shape(attr.None))
}

func (c *languageContainer) R(opts ...attr.Attribute) *node.Node {
	return node.New("r", attr.AssetImage("assets/programming/language/r.png"), attr.Shape(attr.None))
}

func (c *languageContainer) Ruby(opts ...attr.Attribute) *node.Node {
	return node.New("ruby", attr.AssetImage("assets/programming/language/ruby.png"), attr.Shape(attr.None))
}

func (c *languageContainer) Cpp(opts ...attr.Attribute) *node.Node {
	return node.New("cpp", attr.AssetImage("assets/programming/language/cpp.png"), attr.Shape(attr.None))
}

func (c *languageContainer) Dart(opts ...attr.Attribute) *node.Node {
	return node.New("dart", attr.AssetImage("assets/programming/language/dart.png"), attr.Shape(attr.None))
}
