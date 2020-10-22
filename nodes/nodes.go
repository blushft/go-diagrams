package nodes

import (
	"github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
	"strings"
)

type RegistrationFunc func(opts ...attr.Attribute) *node.Node

var RegisteredNodes = map[string]RegistrationFunc{}

func Register(namespace, name string, registrationFunc RegistrationFunc){
	RegisteredNodes[strings.ToLower(namespace + "." + name)] = registrationFunc
}