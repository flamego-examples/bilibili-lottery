package embed

import (
	"embed"
	"io/fs"
)

//go:embed *
var webAssets embed.FS

func WebAssets() fs.FS {
	fs, err := fs.Sub(webAssets, "build")
	if err != nil {
		panic(err)
	}
	return fs
}
