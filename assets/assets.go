package assets

import (
	"embed"
	"io/fs"
	"strings"
)

//go:embed *
var assetFS embed.FS

func cleanPath(p string) string {
	p = strings.TrimPrefix(p, "assets")
	return strings.TrimPrefix(p, "/")
}

func ReadFile(p string) ([]byte, error) {
	return assetFS.ReadFile(cleanPath(p))
}

func ReadDir(p string) ([]fs.DirEntry, error) {
	return assetFS.ReadDir(cleanPath(p))
}
