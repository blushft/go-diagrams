package diagram

import (
	"github.com/blushft/go-diagrams/nodes"
	"github.com/sirupsen/logrus"
	"os"
	"regexp"
	"strings"

	"github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
	_ "github.com/blushft/go-diagrams/nodes"
	"github.com/google/uuid"
)

type GroupType string

const (
	RootGraph GroupType = "G"
	Subgraph  GroupType = "S"
	Cluster   GroupType = "C"
)

func New(id string, attrs ...attr.Attribute) *Group {
	return new(id, RootGraph, defaultRootAttrs(attrs...))
}

type Group struct {
	t        GroupType
	id       string
	attrs    attr.Attributes
	parent   *Group
	children map[string]*Group

	nodes map[string]INode
	edges map[string]IEdge
}

func NewGroup(id string, attrs ...attr.Attribute) *Group {
	return new(id, Subgraph, defaultGroupAttrs(attrs...))
}

func NewCluster(id string, attrs ...attr.Attribute) *Group {
	return new("cluster"+id, Cluster, defaultClusterAttrs(attrs...))
}

func new(id string, t GroupType, attrs attr.Attributes) *Group {
	return &Group{
		t:        t,
		id:       id,
		attrs:    attrs,
		children: make(map[string]*Group),
		nodes:    make(map[string]INode),
		edges:    make(map[string]IEdge),
	}
}

func (g *Group) ID() string {
	return g.id
}

func (g *Group) Attributes() (map[string]string, error) {
	return g.attrs.Render()
}

func (g *Group) Nodes() []INode {
	ns := make([]INode, 0, len(g.nodes))

	for _, n := range g.nodes {
		ns = append(ns, n)
	}

	return ns
}

func (g *Group) Edges() []IEdge {
	es := make([]IEdge, 0, len(g.edges))

	for _, e := range g.edges {
		es = append(es, e)
	}

	return es
}

func (g *Group) Children() []IDiagram {
	cs := make([]IDiagram, 0, len(g.children))

	for _, c := range g.children {
		cs = append(cs, c)
	}

	return cs
}

func (g *Group) AddChildren(group *Group) {
	g.children[group.id] = group
}

func (g *Group) AddNode(n INode) *Group {
	g.nodes[n.ID()] = n
	return g
}

func (g *Group) AddNodes(ns ...INode) *Group {
	for _, n := range ns {
		g.AddNode(n)
	}

	return g
}

func (g *Group) Connect(start, end INode) *Group {
	return g.ConnectWithAttrs(start, end)
}

func (g *Group) ConnectWithAttrs(start, end INode, attributes ...attr.Attribute) *Group {
	g.AddNodes(start, end)
	attributes = append(attributes, attr.Arrowhead(attr.Normal))
	attributes = append(attributes, attr.Direction(attr.Forward))
	eid := strings.SplitAfter(strings.Replace(uuid.New().String(), "-", "", 1), "-")[0]
	g.edges[eid] = NewEdge(eid, start, end, attributes...)

	return g
}


func (g *Group) Parse(file *os.File) error {
	return g.parse(file)
}

func (g *Group) interpretYML(yml *DiagramData) error {
	var attrs []attr.Attribute

	if yml.Label != "" {
		attrs = append(attrs, attr.Label(yml.Label))
	}

	if yml.Direction != "" {
		attrs = append(attrs, attr.Direction(yml.Direction))
	}
	
	if yml.Direction != attr.EdgeNone {
		attrs = append(attrs, attr.Directed(true))
	}

	g.attrs.Set(attrs...)

	for _, v := range yml.Groups {
		parsedGroup := g.parseGroup(v)
		g.AddChildren(&parsedGroup)
	}


	for _, v := range yml.Groups {
		g.parseGroupConnection(v)
	}

	return nil
}

func (g *Group) parseNode(v NodeData) INode {
	attrs := parseAttributes(v.Meta)

	if v.Name == "" {
		v.Name = uuid.New().String()
	}

	// Get n init
	r, found := nodes.RegisteredNodes[strings.ToLower(v.Type)]
	if !found {
		logrus.Warnf("type %v not found", v.Type)
		return node.New(v.Name, attrs...)
	}
	
	n := r(attrs...)
	n.SetID(v.Name)
	return n
}

func parseAttributes(v Meta) []attr.Attribute {
	var attrs []attr.Attribute
	if v.Label != "" {
		attrs = append(attrs, attr.Label(v.Label))
	}

	if v.BackgroundColor != "" {
		attrs = append(attrs, attr.BGColor(v.BackgroundColor))
	}
	
	if v.LabelLoc != "" {
		attrs = append(attrs, attr.LabelLocation(v.LabelLoc))
	} else {
		attrs = append(attrs, attr.LabelLocation(attr.Bottom))
	}
	
	attrs = append(attrs, attr.ImageScale("true"))
	attrs = append(attrs, attr.NoJustify(true))

	return node.DefaultAttributes(attrs...)
}

func connectionLabel(s string) (string, string) {
	re := regexp.MustCompile(`(.*)\[(.*?)]`)
	if !re.MatchString(s){
		return s, ""
	}
	
	matches := re.FindStringSubmatch(s)
	
	return matches[1], matches[2]
}

func (g *Group) parseGroup(v GroupData) Group {
	attrs := parseAttributes(v.Meta)
	
	if v.Name == "" {
		v.Name = uuid.New().String()
	}

	group := NewCluster(v.Name, attrs...)

	for _, child := range v.Nodes {
		parsedNode := g.parseNode(child)
		group.AddNode(parsedNode)
	}
	
	for _, groupData := range v.Group {
		parsedGroup := g.parseGroup(groupData)
		group.AddChildren(&parsedGroup)
	}

	return *group
}

func (g *Group) findByID(to string) INode {
	// Get to the root
	var root = g
	for ;; {
		if root.parent == nil {
			break
		} else {
			root = root.parent
		}
	}
	
	// Search
	return innerFindById(root, to)
}

func (g *Group) parseGroupConnection(v GroupData) {
	for _, srcNode := range v.Nodes {
		srcRef := g.findByID(srcNode.Name)
		if srcRef == nil {
			continue
		}
		for _, dest := range v.ConnectAllTo {
			nodeId, label := connectionLabel(dest)
			destRef := g.findByID(nodeId)
			if destRef != nil {
				if label != "" {
					g.ConnectWithAttrs(srcRef, destRef, attr.Label(label),
						attr.HeadLabel(label))
				} else {
					g.Connect(srcRef, destRef)
				}
			}
		}

		for _, dest := range srcNode.ConnectTo {
			nodeId, label := connectionLabel(dest)
			destRef := g.findByID(nodeId)
			if destRef != nil {
				if label != "" {
					g.ConnectWithAttrs(srcRef, destRef,
						attr.HeadLabel(label))
				} else {
					g.Connect(srcRef, destRef)
				}
			}
		}
	}
	
	for _, ig := range v.Group {
		g.parseGroupConnection(ig)
	}
}

func innerFindById(root *Group, to string) INode {
	if root == nil {
		return nil
	}
	
	if root.ID() == to || root.ID() == "cluster" + to {
		return root
	}

	// Search children
	for _, v := range root.children {
		found := innerFindById(v, to)
		if found != nil {
			return found
		}
	}


	for _, v := range root.nodes {
		if v.ID() == to || v.ID() == "cluster" + to {
			return v
		}
	}
	
	return nil
}
