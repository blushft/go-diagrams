package node

import (
	"strconv"
	"strings"

	graphviz "github.com/awalterschulze/gographviz"
	"github.com/blushft/go-diagrams/font"
	"github.com/blushft/go-diagrams/pkg/rand"
)

type Node struct {
	id      string
	Options Options
}

func New(opts ...Option) *Node {
	options := DefaultOptions(opts...)

	return &Node{
		id:      rand.String(8),
		Options: options,
	}
}

func (n *Node) ID() string {
	return n.id
}

func (n *Node) Graph(name string, graph *graphviz.Escape) error {
	return graph.AddNode(name, n.id, n.attrs())
}

func (n *Node) attrs() map[string]string {

	if n.Options.Icon != "" {
		labelH := float64(len(strings.Split(n.Options.Label, "/n")))
		n.Options.Height = n.Options.Height + (0.4 * labelH)
	}

	return map[string]string{
		"label":      n.Options.Label,
		"image":      n.Options.Icon,
		"height":     strconv.FormatFloat(n.Options.Height, 'f', -1, 64),
		"width":      strconv.FormatFloat(n.Options.Width, 'f', -1, 64),
		"shape":      n.Options.Shape,
		"style":      n.Options.Style,
		"labelloc":   n.Options.LabelLocation,
		"imagescale": "true",
		"fixedsize":  "true",
		"fontsize":   "13",
		"fontname":   "Sans-Serif",
		"fontcolor":  "#2D3436",
	}
}

func Graph(name string, graph *graphviz.Escape, nodes ...*Node) error {
	for _, n := range nodes {
		err := n.Graph(name, graph)
		if err != nil {
			return err
		}
	}

	return nil
}

type Options struct {
	Name          string
	Label         string
	Provider      string
	Icon          string
	IconDirection string
	Shape         string
	Style         string
	FixedSize     bool
	Width         float64
	Height        float64
	LabelLocation string
	ImageScale    bool
	Font          font.Options
}

type Option func(*Options)

type OptionSet []Option

func DefaultOptions(opts ...Option) Options {
	nopts := Options{
		Name:          "node",
		Label:         "node",
		Shape:         "box",
		Style:         "rounded",
		FixedSize:     true,
		Width:         1.4,
		Height:        1.4,
		LabelLocation: "b",
		ImageScale:    true,
		Font:          font.DefaultOptions(),
	}

	for _, o := range opts {
		o(&nopts)
	}

	return nopts
}

func MergeOptionSets(sets ...OptionSet) OptionSet {
	ns := OptionSet{}

	for _, set := range sets {
		for _, opt := range set {
			ns = append(ns, opt)
		}
	}

	return ns
}

func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

func Label(l string) Option {
	return func(o *Options) {
		o.Label = l
	}
}

func Provider(p string) Option {
	return func(o *Options) {
		o.Provider = p
	}
}

func Icon(i string) Option {
	return func(o *Options) {
		o.Icon = i
	}
}

func IconDirection(d string) Option {
	return func(o *Options) {
		o.IconDirection = d
	}
}

func Shape(s string) Option {
	return func(o *Options) {
		o.Shape = s
	}
}

func Style(s string) Option {
	return func(o *Options) {
		o.Style = s
	}
}

func FixedSize(b bool) Option {
	return func(o *Options) {
		o.FixedSize = b
	}
}

func Width(w float64) Option {
	return func(o *Options) {
		o.Width = w
	}
}

func Height(h float64) Option {
	return func(o *Options) {
		o.Height = h
	}
}

func LabelLocation(l string) Option {
	return func(o *Options) {
		o.LabelLocation = l
	}
}

func ImageScale(b bool) Option {
	return func(o *Options) {
		o.ImageScale = b
	}
}

func SetFontOptions(f font.Options) Option {
	return func(o *Options) {
		o.Font = f
	}
}
