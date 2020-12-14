package nodes

import attr "github.com/blushft/go-diagrams/attr"

type firebaseContainer struct {
	path  string
	attrs []attr.Attribute
}

var Firebase = &firebaseContainer{path: "assets/firebase/firebase"}
