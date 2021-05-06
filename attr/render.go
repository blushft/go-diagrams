package attr

import (
	"os"
	"path/filepath"

	"github.com/blushft/go-diagrams/assets"
)

func AssetLoader(dir string) RenderHook {
	return func(a Attribute) (Attribute, error) {
		if a.Name() != "assetimage" {
			return a, nil
		}

		if err := writeAssetImage(a.Value().String(), dir); err != nil {
			return nil, err
		}

		return Image(a.Value().String()), nil
	}
}

func writeAssetImage(img, out string) error {
	b, err := assets.ReadFile(img)
	if err != nil {
		return err
	}

	outdir := filepath.Join(out, filepath.Dir(img))
	if err := os.MkdirAll(outdir, os.ModePerm); err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	fp := filepath.Join(outdir, filepath.Base(img))
	if err := os.WriteFile(fp, b, os.ModePerm); err != nil {
		return err
	}

	return nil
}
