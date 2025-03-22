package diagram

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func (d *Diagram) parseYML(inputFile *os.File) (*DiagramData, error) {
	var decodedYML *DiagramData

	decoder := yaml.NewDecoder(inputFile)
	decoder.KnownFields(true)
	err := decoder.Decode(&decodedYML)

	if err != nil {
		return nil, err
	}

	return decodedYML, nil
}

func (d *Diagram) parse(inputFile *os.File) error {
	if inputFile == nil {
		return fmt.Errorf("input file not specified")
	}

	yml, err := d.parseYML(inputFile)
	if err != nil {
		return err
	}

	return d.interpretYML(yml)
}
