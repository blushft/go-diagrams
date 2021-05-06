package diagram

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/blushft/go-diagrams/attr"

	graphviz "github.com/awalterschulze/gographviz"
)

type RenderOptions struct {
	Format     string
	Filename   string
	OutputPath string
	Hooks      []RenderHookFunc

	hooks []attr.RenderHook
}

type RenderOption func(*RenderOptions)
type RenderHookFunc func(RenderOptions) attr.RenderHook

func defaultRenderOptions(opts ...RenderOption) RenderOptions {
	options := RenderOptions{
		Format:     "dot",
		Filename:   "go-diagrams",
		OutputPath: "go-diagrams",
		Hooks: []RenderHookFunc{
			func(o RenderOptions) attr.RenderHook {
				return attr.AssetLoader(o.OutputPath)
			},
		},
	}

	for _, o := range opts {
		o(&options)
	}

	var hooks []attr.RenderHook
	for _, h := range options.Hooks {
		hooks = append(hooks, h(options))
	}

	options.hooks = hooks

	return options
}

func Format(f string) RenderOption {
	return func(o *RenderOptions) {
		o.Format = f
	}
}

func OutputPath(dir string) RenderOption {
	return func(o *RenderOptions) {
		o.OutputPath = dir
	}
}

func RegisterHook(fn RenderHookFunc) RenderOption {
	return func(o *RenderOptions) {
		o.Hooks = append(o.Hooks, fn)
	}
}

func Render(d Diagram, opts ...RenderOption) error {
	ro := defaultRenderOptions(opts...)
	g := graphviz.NewEscape()

	if err := g.SetName(d.ID()); err != nil {
		return err
	}

	gattr, err := d.Attributes().Render(ro.hooks...)
	if err != nil {
		return err
	}

	dir := false
	b, ok := gattr["dir"]
	if ok {
		b, err := strconv.ParseBool(b)
		if err == nil {
			dir = b
		}
	}

	if err := g.SetDir(dir); err != nil {
		return err
	}

	for k, v := range gattr {
		if err := g.AddAttr(d.ID(), k, v); err != nil {
			return err
		}
	}

	if err := renderGroups(g, d, ro); err != nil {
		return err
	}

	return renderOutput(g, ro)
}

func renderGroups(g *graphviz.Escape, d Diagram, ro RenderOptions) error {
	for _, n := range d.Nodes() {
		attr, err := n.Attributes().Render(ro.hooks...)
		if err != nil {
			return err
		}

		if err := g.AddNode(d.ID(), n.ID(), attr); err != nil {
			return err
		}
	}

	for _, e := range d.Edges() {
		attr, err := e.Attributes().Render(ro.hooks...)
		if err != nil {
			return err
		}

		if err := g.AddEdge(e.Start(), e.End(), true, attr); err != nil {
			return err
		}
	}

	for _, child := range d.Children() {
		attr, err := d.Attributes().Render(ro.hooks...)
		if err != nil {
			return err
		}

		if err := g.AddSubGraph(d.ID(), child.ID(), attr); err != nil {
			return err
		}

		if err := renderGroups(g, child, ro); err != nil {
			return err
		}
	}

	return nil
}

func renderOutput(g *graphviz.Escape, ro RenderOptions) error {
	outPath, err := filepath.Abs(ro.OutputPath)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(outPath, os.ModePerm); err != nil {
		return err
	}

	switch ro.Format {
	case "dot":
		return renderDot(g, ro.Filename, outPath)
	default:
		return errors.New("invalid output format")
	}
}

func renderDot(g *graphviz.Escape, fname, outPath string) error {
	outFile := filepath.Join(outPath, fname+".dot")

	return ioutil.WriteFile(outFile, []byte(g.String()), os.ModePerm)
}
