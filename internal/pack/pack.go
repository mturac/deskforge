package pack

import (
	"embed"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:embed all:tree
var tree embed.FS

// File is one embedded path written by init.
type File struct {
	Path string
	Data []byte
}

// Files returns every template file under embed/tree.
func Files() ([]File, error) {
	var out []File
	err := fs.WalkDir(tree, "tree", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		data, err := tree.ReadFile(path)
		if err != nil {
			return err
		}
		rel := strings.TrimPrefix(path, "tree/")
		rel = filepath.FromSlash(rel)
		out = append(out, File{Path: rel, Data: data})
		return nil
	})
	return out, err
}
