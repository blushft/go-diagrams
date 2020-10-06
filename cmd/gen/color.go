package main

import (
	"bufio"
	"image/color"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/urfave/cli/v2"
)

func genColorsCmd(c *cli.Context) error {
	gc, err := loadColors()
	if err != nil {
		return err
	}

	return genColors(gc)
}

type genColor struct {
	Color color.RGBA
	Name  string
}

func loadColors() ([]genColor, error) {
	cFile, err := os.Open("./fixtures/rgb.txt")
	if err != nil {
		return nil, err
	}
	defer cFile.Close()

	colors := []genColor{}

	scanner := bufio.NewScanner(cFile)
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			continue
		}

		name := strings.TrimRight(string(row[13:]), " ")
		if strings.Contains(name, " ") {
			continue
		}

		r, err := toUint(string(row[:3]))
		if err != nil {
			return nil, err
		}

		g, err := toUint(string(row[4:7]))
		if err != nil {
			return nil, err
		}

		b, err := toUint(string(row[8:11]))
		if err != nil {
			return nil, err
		}

		tc := color.RGBA{
			R: uint8(r),
			B: uint8(g),
			G: uint8(b),
			A: 255,
		}

		c := genColor{
			Color: tc,
			Name:  strcase.ToCamel(name),
		}

		colors = append(colors, c)
	}

	return colors, nil
}

func toUint(s string) (uint8, error) {
	s = strings.Trim(s, " ")
	i, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return 0, err
	}

	return uint8(i), nil
}

func genColors(gc []genColor) error {
	outpkg := "attr/color"
	f := jen.NewFilePathName("attr/color", "color")

	for _, c := range gc {
		f.Func().Id(c.Name).Call().Id("Color").BlockFunc(func(g *jen.Group) {
			g.ReturnFunc(func(rg *jen.Group) {
				rg.Id("newColorAttr").Call(jen.Qual("image/color", "RGBA").Values(
					jen.Dict{
						jen.Id("R"): jen.Lit(c.Color.R),
						jen.Id("B"): jen.Lit(c.Color.B),
						jen.Id("G"): jen.Lit(c.Color.G),
						jen.Id("A"): jen.Lit(c.Color.A),
					},
				))
			})
		}).Line()
	}

	outpath := filepath.Join(outpkg, "colors.go")

	return f.Save(outpath)
}
