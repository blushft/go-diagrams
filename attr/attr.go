package attr

import graphviz "github.com/awalterschulze/gographviz"

type AttributableElement int

func (ae AttributableElement) String() string {
	switch ae {
	case 0:
		return "any"
	case 1:
		return "edges"
	case 2:
		return "nodes"
	case 3:
		return "graph"
	case 4:
		return "subgraphs"
	case 5:
		return "clusters"
	default:
		return "invalid"
	}
}

const (
	Any AttributableElement = iota
	Edges
	Nodes
	Graph
	Subs
	Clusters
)

type Attributes map[string]Attribute
type AttributeApplier func(Attributes)
type AttributeFormatter func(AttributeApplier) AttributeApplier

type Attribute interface {
	Name() string
	Value() *Value
	Render(map[string]string) error
	ApplyTo(AttributableElement) bool
	Apply(Attributes)
}

func NewAttributes(attrs ...Attribute) Attributes {
	a := make(map[string]Attribute)

	for _, attr := range attrs {
		attr.Apply(a)
	}

	return a
}

func (a Attributes) Render() (map[string]string, error) {
	m := make(map[string]string)

	for _, attr := range a {
		if err := attr.Render(m); err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (a Attributes) Set(attrs ...Attribute) {
	for _, attr := range attrs {
		attr.Apply(a)
	}
}

type attribute struct {
	name    graphviz.Attr
	value   *Value
	applyTo []AttributableElement
}

func newAttr(n graphviz.Attr, v interface{}, applyTo ...AttributableElement) *attribute {
	return &attribute{
		name:    n,
		value:   NewValue(v),
		applyTo: applyTo,
	}

}

func (a *attribute) Name() string {
	var name = (*string)(&a.name)
	return *name
}

func (a *attribute) Value() *Value {
	return a.value
}

func (a *attribute) Render(m map[string]string) error {
	m[a.Name()] = a.value.String()
	return nil
}

func (a *attribute) ApplyTo(ae AttributableElement) bool {
	for _, typ := range a.applyTo {
		if ae == typ {
			return true
		}
	}

	return false
}

func (a *attribute) Apply(m Attributes) {
	m[a.Name()] = a
}

type attributeSet struct {
	attrs []Attribute
}

func newAttrSet(attrs ...Attribute) *attributeSet {
	return &attributeSet{
		attrs: attrs,
	}
}

func (a *attributeSet) Name() string {
	return "attr_set"
}

func (a *attributeSet) Value() *Value {
	return nil
}

func (a *attributeSet) Render(map[string]string) error {
	return nil
}

func (a *attributeSet) ApplyTo(AttributableElement) bool {
	return false
}

func (a *attributeSet) Apply(m Attributes) {
	for _, attr := range a.attrs {
		attr.Apply(m)
	}
}
