package diagram

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/blushft/go-diagrams/assets"

	graphviz "github.com/awalterschulze/gographviz"
)

type RenderOptions struct {
	Format     string
	Filename   string
	OutputPath string
}

type RenderOption func(*RenderOptions)

func defaultRenderOptions(opts ...RenderOption) RenderOptions {
	options := RenderOptions{
		Format:     "dot",
		Filename:   "go-diagrams",
		OutputPath: "go-diagrams",
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

func Format(f string) RenderOption {
	return func(o *RenderOptions) {
		o.Format = f
	}
}

type Identifiable interface {
	ID() string
}

type Attributed interface {
	Attributes() (map[string]string, error)
}

type Diagram interface {
	Identifiable
	Attributed
	Nodes() []Node
	Edges() []Edge
	Children() []Diagram
}

type Node interface {
	Identifiable
	Attributed
}

type Edge interface {
	Identifiable
	Attributed
	Start() string
	End() string
}

func Render(d Diagram, opts ...RenderOption) error {
	ro := defaultRenderOptions(opts...)
	g := graphviz.NewEscape()

	if err := g.SetName(d.ID()); err != nil {
		return err
	}

	gattr, err := d.Attributes()
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

	if err := renderGroups(g, d); err != nil {
		return err
	}

	return renderOutput(g, ro)
}

func renderGroups(g *graphviz.Escape, d Diagram) error {
	for _, n := range d.Nodes() {
		attr, err := n.Attributes()
		if err != nil {
			return err
		}

		if err := g.AddNode(d.ID(), n.ID(), attr); err != nil {
			return err
		}
	}

	for _, e := range d.Edges() {
		attr, err := e.Attributes()
		if err != nil {
			return err
		}

		if err := g.AddEdge(e.Start(), e.End(), true, attr); err != nil {
			return err
		}
	}

	for _, child := range d.Children() {
		attr, err := d.Attributes()
		if err != nil {
			return err
		}

		if err := g.AddSubGraph(d.ID(), child.ID(), attr); err != nil {
			return err
		}

		if err := renderGroups(g, child); err != nil {
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

	for _, n := range g.Nodes.Nodes {
		if img, ok := n.Attrs["image"]; ok {
			iData, err := assets.ReadFile(img)
			if err != nil {
				return err
			}

			iPath := filepath.Join(outPath, filepath.Dir(img))
			if err := os.MkdirAll(iPath, os.ModePerm); err != nil {
				return err
			}

			iFile := filepath.Join(iPath, filepath.Base(img))
			if err := ioutil.WriteFile(iFile, iData, os.ModePerm); err != nil {
				return err
			}
		}
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
