//go:build !embed_frontend

package views

import (
	"embed"
	"io/fs"
	"os"
)

//go:embed fallback_dist
var WebStaticFallbackFS embed.FS

func WebStaticRoot() (fs.FS, error) {
	localDist := os.DirFS("views/dist")
	if _, err := fs.Stat(localDist, "index.html"); err == nil {
		return localDist, nil
	}

	return fs.Sub(WebStaticFallbackFS, "fallback_dist")
}
