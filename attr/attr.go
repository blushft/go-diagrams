package attr

type Attributes map[string]string

type Attribute func(Attributes)

func New() Attributes {
	return make(map[string]string)
}

func (a Attributes) Set(attrs ...Attribute) {
	for _, at := range attrs {
		at(a)
	}
}

func Label(s string) Attribute {
	return func(a Attributes) {
		a["label"] = s
	}
}

func Shape(s string) Attribute {
	return func(a Attributes) {
		a["shape"] = s
	}
}

func Image(s string) Attribute {
	return func(a Attributes) {
		a["image"] = s
	}
}

func ImagePosition(s string) Attribute {
	return func(a Attributes) {
		a["imagepos"] = s
	}
}

func ImageScale(s string) Attribute {
	return func(a Attributes) {
		a["imagescale"] = s
	}
}
