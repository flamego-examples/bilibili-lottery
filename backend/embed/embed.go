package embed

import (
	"embed"
	"io/fs"
	"path"
)

//go:embed build/*
var assets embed.FS

type webAssets struct {
	embed.FS
}

func (w *webAssets) Open(name string) (fs.File, error) {
	return w.FS.Open(path.Join("build", name))
}

func WebAssets() fs.FS {
	return &webAssets{assets}
}
