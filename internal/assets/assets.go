package assets

import (
	"embed"
	"io/fs"
)

//go:embed *
var assetFs embed.FS

func Open(name string) (fs.File, error) {
	return assetFs.Open(name)
}

func ReadDir(name string) ([]fs.DirEntry, error) {
	return assetFs.ReadDir(name)
}

func ReadFile(name string) ([]byte, error) {
	return assetFs.ReadFile(name)
}
