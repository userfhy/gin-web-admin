//go:build embed_frontend

package views

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var WebStaticFS embed.FS

func WebStaticRoot() (fs.FS, error) {
	return fs.Sub(WebStaticFS, "dist")
}
