package diagram

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func (g *Group) parseYML(inputFile *os.File) (*DiagramData, error) {
	var decodedYML *DiagramData

	decoder := yaml.NewDecoder(inputFile)
	decoder.KnownFields(true)
	err := decoder.Decode(&decodedYML)

	if err != nil {
		return nil, err
	}

	return decodedYML, nil
}

func (g *Group) parse(inputFile *os.File) error {
	if inputFile == nil {
		return fmt.Errorf("input file not specified")
	}

	yml, err := g.parseYML(inputFile)
	if err != nil {
		return err
	}

	return g.interpretYML(yml)
}
