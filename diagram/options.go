package diagram

import (
	"os"
	"strconv"
)

type Options struct {
	Name       string
	FileName   string
	InputFile  *os.File
	OutFormat  string
	Direction  string
	CurveStyle string
	Show       bool
	Label      string
	Pad        float64
	Splines    string
	NodeSep    float64
	RankSep    float64
	Font       Font
	Attributes map[string]string
}

func (o Options) attrs() map[string]string {
	m := map[string]string{
		"pad":       strconv.FormatFloat(o.Pad, 'f', -1, 64),
		"label":     o.Label,
		"splines":   o.Splines,
		"nodesep":   strconv.FormatFloat(o.NodeSep, 'f', -1, 64),
		"rankdir":   o.Direction,
		"ranksep":   strconv.FormatFloat(o.RankSep, 'f', -1, 64),
		"fontname":  o.Font.Name,
		"fontsize":  strconv.FormatInt(int64(o.Font.Size), 10),
		"fontcolor": o.Font.Color,
	}

	for k, v := range o.Attributes {
		m[k] = v
	}

	return m
}

type Option func(*Options)

func DefaultOptions(opts ...Option) Options {
	options := Options{
		Name:       "go-diagrams",
		FileName:   "go-diagram",
		OutFormat:  "dot",
		Label:      "",
		Direction:  string(LeftToRight),
		CurveStyle: "ortho",
		Show:       true,
		Pad:        2.0,
		Splines:    "ortho",
		NodeSep:    0.60,
		RankSep:    0.75,
		Font:       defaultFont(),
		Attributes: make(map[string]string),
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

func Filename(f string) Option {
	return func(o *Options) {
		o.FileName = f
	}
}

func DirName(d string) Option {
	return func(o *Options) {
		o.Name = d
	}
}

func Label(l string) Option {
	return func(o *Options) {
		o.Label = l
	}
}

func Direction(d string) Option {
	return func(o *Options) {
		o.Direction = d
	}
}

func WithAttribute(name, value string) Option {
	return func(o *Options) {
		o.Attributes[name] = value
	}
}

func WithAttributes(attrs map[string]string) Option {
	return func(o *Options) {
		for k, v := range attrs {
			o.Attributes[k] = v
		}
	}
}

func PenColor(c string) Option {
	return func(o *Options) {
		o.Attributes["pencolor"] = c
	}
}

func Shape(s string) Option {
	return func(o *Options) {
		o.Attributes["shape"] = s
	}
}

func Style(s string) Option {
	return func(o *Options) {
		o.Attributes["style"] = s
	}
}

func LabelJustify(j string) Option {
	return func(o *Options) {
		o.Attributes["labeljust"] = j
	}
}
