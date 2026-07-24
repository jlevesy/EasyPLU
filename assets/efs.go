package assets

import (
	"embed"
	"encoding/json"
	"io/fs"
)

//go:embed "html" "static"
var files embed.FS

var (
	HTMLFiles   = sub(files, "html")
	StaticFiles = sub(files, "static")
)

func sub(f embed.FS, dir string) fs.FS {
	sub, err := fs.Sub(f, dir)
	if err != nil {
		panic(err)
	}
	return sub
}

type AssetEntry struct {
	File string `json:"file"`
}

type AssetManifest map[string]AssetEntry

var Manifest AssetManifest

func init() {
	f, err := StaticFiles.Open("assets-manifest.json")
	if err != nil {
		panic(err)
	}

	if err := json.NewDecoder(f).Decode(&Manifest); err != nil {
		panic(err)
	}
}
