package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

const (
	diaPkg = "github.com/blushft/go-diagrams/diagram"
)

type provider struct {
	name   string
	root   string
	images map[string]map[string]string
}

func loadProvider(name, root string) (*provider, error) {
	np := &provider{
		name:   name,
		root:   root,
		images: make(map[string]map[string]string),
	}

	current := ""
	if err := filepath.Walk(root, func(p string, info os.FileInfo, ferr error) error {
		if p == root {
			return nil
		}

		if info.IsDir() {
			_, cont := filepath.Split(p)
			np.images[cont] = make(map[string]string)
			current = cont

			return nil
		}

		if filepath.Ext(p) == ".png" {
			dir, img := filepath.Split(p)
			if strings.Contains(dir, current) {
				im, ok := np.images[current]
				if !ok {
					return nil
				}
				im[strings.Replace(img, filepath.Ext(img), "", -1)] = p
				np.images[current] = im
			}

		}

		return nil
	}); err != nil {
		return nil, err
	}

	return np, nil
}

func loadProviders(root string) (map[string]*provider, error) {
	providers := map[string]*provider{}
	provDirs := map[string]string{}

	//var current *provider
	if err := filepath.Walk(root, func(p string, info os.FileInfo, ferr error) error {
		if info.IsDir() {
			dir, prov := filepath.Split(p)
			if strings.ReplaceAll(dir, "/", "") == root {
				provDirs[prov] = p
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	for pv, pr := range provDirs {
		prov, err := loadProvider(pv, pr)
		if err != nil {
			return nil, err
		}

		providers[pv] = prov
	}

	return providers, nil
}

func main() {
	providers, err := loadProviders("assets")
	if err != nil {
		log.Fatalf("load profiles: %v", err)
	}

	for _, p := range providers {
		if err := genProvider("nodes", p); err != nil {
			log.Fatalf("generate %s: %v", p.name, err)
		}
	}
}

func genProvider(outroot string, p *provider) error {
	rootPath := filepath.Join(outroot, p.name)
	if err := os.MkdirAll(rootPath, os.ModePerm); err != nil {
		return err
	}

	for c, imgs := range p.images {
		if err := genContainer(rootPath, p.name, p.root, c, imgs); err != nil {
			return err
		}
	}

	return nil
}

func genContainer(root, pName, pPath, name string, imgs map[string]string) error {
	f := jen.NewFile(pName)

	cname := name + "Container"
	expCName := strcase.ToCamel(name)
	croot := filepath.Join(root, name)
	imgpath := filepath.Join(pPath, name)

	f.ImportName(diaPkg, "diagram")

	f.Type().Id(cname).Struct(jen.Id("path").String(), jen.Id("opts").Index().Qual(diaPkg, "NodeOption"))

	f.Var().Id(expCName).Op("=").Op("&").Id(cname).Values(
		jen.Dict{
			jen.Id("path"): jen.Lit(imgpath),
			jen.Id("opts"): jen.Qual(diaPkg, "OptionSet").Values(
				jen.Qual(diaPkg, "Provider").Call(jen.Lit(pName)),
				jen.Qual(diaPkg, "NodeShape").Call(jen.Lit("none")),
			),
		},
	)

	for img, imgPath := range imgs {
		method := strcase.ToCamel(img)
		f.Func().Params(jen.Id("c").Op("*").Id(cname)).
			Id(method).
			Params(
				jen.Id("opts").Op("...").Qual(diaPkg, "NodeOption"),
			).
			Op("*").Qual(diaPkg, "Node").
			Block(
				jen.Id("nopts").Op(":=").Qual(diaPkg, "MergeOptionSets").Call(
					jen.Qual(diaPkg, "OptionSet").Values(
						jen.Qual(diaPkg, "Icon").Call(jen.Lit(imgPath)),
					),
					jen.Id("c").Dot("opts"),
					jen.Id("opts"),
				),
				jen.Return(jen.Qual(diaPkg, "NewNode").Call(
					jen.Id("nopts").Op("..."),
				)),
			).Line()
	}

	return f.Save(croot + ".go")
}
