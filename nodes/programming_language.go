package nodes

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
	opts = append(opts, attr.AssetImage("assets/programming/language/bash.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("bash", opts...)
}

func (c *languageContainer) Nodejs(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/nodejs.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("nodejs", opts...)
}

func (c *languageContainer) Matlab(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/matlab.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("matlab", opts...)
}

func (c *languageContainer) Python(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/python.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("python", opts...)
}

func (c *languageContainer) Rust(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/rust.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("rust", opts...)
}

func (c *languageContainer) Swift(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/swift.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("swift", opts...)
}

func (c *languageContainer) Typescript(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/typescript.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("typescript", opts...)
}

func (c *languageContainer) Csharp(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/csharp.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("csharp", opts...)
}

func (c *languageContainer) Java(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/java.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("java", opts...)
}

func (c *languageContainer) Javascript(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/javascript.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("javascript", opts...)
}

func (c *languageContainer) Kotlin(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/kotlin.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("kotlin", opts...)
}

func (c *languageContainer) C(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/c.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("c", opts...)
}

func (c *languageContainer) Go(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/go.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("go", opts...)
}

func (c *languageContainer) Php(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/php.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("php", opts...)
}

func (c *languageContainer) R(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/r.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("r", opts...)
}

func (c *languageContainer) Ruby(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/ruby.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ruby", opts...)
}

func (c *languageContainer) Cpp(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/cpp.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("cpp", opts...)
}

func (c *languageContainer) Dart(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/programming/language/dart.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("dart", opts...)
}

func init() {
	const namespace = "programming.language"
	Register(namespace, "Bash", Language.Bash)
	Register(namespace, "Nodejs", Language.Nodejs)
	Register(namespace, "Matlab", Language.Matlab)
	Register(namespace, "Python", Language.Python)
	Register(namespace, "Rust", Language.Rust)
	Register(namespace, "Swift", Language.Swift)
	Register(namespace, "Typescript", Language.Typescript)
	Register(namespace, "Csharp", Language.Csharp)
	Register(namespace, "Java", Language.Java)
	Register(namespace, "Javascript", Language.Javascript)
	Register(namespace, "Kotlin", Language.Kotlin)
	Register(namespace, "C", Language.C)
	Register(namespace, "Go", Language.Go)
	Register(namespace, "Php", Language.Php)
	Register(namespace, "R", Language.R)
	Register(namespace, "Ruby", Language.Ruby)
	Register(namespace, "Cpp", Language.Cpp)
	Register(namespace, "Dart", Language.Dart)
}
