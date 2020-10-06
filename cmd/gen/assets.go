package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

const (
	nodePkg = "github.com/blushft/go-diagrams/node"
	attrPkg = "github.com/blushft/go-diagrams/attr"
)

func genAssetsCmd(c *cli.Context) error {
	providers, err := loadProviders("fixtures/assets", "assets")
	if err != nil {
		return errors.Wrap(err, "load profiles")
	}

	for _, p := range providers {
		if err := genProvider("nodes", p); err != nil {
			return errors.Wrap(err, fmt.Sprintf("generate %s", p.name))
		}
	}

	return nil
}

type provider struct {
	name   string
	root   string
	fp     string
	images map[string]map[string]string
}

func loadProviders(fp, root string) (map[string]*provider, error) {
	providers := map[string]*provider{}
	provDirs := map[string]string{}

	//var current *provider
	if err := filepath.Walk(fp, func(p string, info os.FileInfo, ferr error) error {
		if p == root {
			return nil
		}

		if info.IsDir() {
			dir, prov := filepath.Split(p)
			if strings.TrimRight(dir, "/") == fp {
				provDirs[prov] = filepath.Join(root, prov)
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	for pv, pr := range provDirs {
		pfp := filepath.Join(fp, pv)
		prov, err := loadProvider(pfp, pv, pr)
		if err != nil {
			return nil, err
		}

		providers[pv] = prov
	}

	return providers, nil
}

func loadProvider(fp, name, root string) (*provider, error) {
	np := &provider{
		name:   name,
		fp:     fp,
		root:   root,
		images: make(map[string]map[string]string),
	}

	current := ""
	if err := filepath.Walk(fp, func(p string, info os.FileInfo, ferr error) error {
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
				im[strings.Replace(img, filepath.Ext(img), "", -1)] = filepath.Join(root, current, img)
				np.images[current] = im
			}

		}

		return nil
	}); err != nil {
		return nil, err
	}

	return np, nil
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

	f.ImportName(nodePkg, "node")

	f.Type().Id(cname).
		Struct(
			jen.Id("path").String(),
			jen.Id("attrs").Index().Qual(attrPkg, "Attribute"),
		)

	f.Var().Id(expCName).Op("=").Op("&").Id(cname).Values(
		jen.Dict{
			jen.Id("path"): jen.Lit(imgpath),
		},
	)

	for img, imgPath := range imgs {
		method := strcase.ToCamel(img)
		f.Func().Params(jen.Id("c").Op("*").Id(cname)).
			Id(method).
			Params(
				jen.Id("opts").Op("...").Qual(attrPkg, "Attribute"),
			).
			Op("*").Qual(nodePkg, "Node").
			Block(
				jen.Return(jen.Qual(nodePkg, "New").Call(
					jen.Lit(img),
					jen.Qual(attrPkg, "AssetImage").Call(jen.Lit(imgPath)),
					jen.Qual(attrPkg, "Shape").Call(jen.Qual(attrPkg, "None")),
				)),
			).Line()
	}

	return f.Save(croot + ".go")
}
