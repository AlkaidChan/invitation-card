package assets

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
)

// nolint:typecheck
//
//go:embed styles bundle.js index.html
var FrontendAssets embed.FS

var (
	allowCompressExtentions = map[string]bool{
		".js":  true,
		".css": true,
	}
)

// EmbedServer is the server that serves the embedded frontend assets.
type EmbedServer struct {
	dist     fs.FS
	tpl      *template.Template
	fsServer http.Handler
}

// NewEmbedServer creates a new EmbedServer.
func NewEmbedServer() *EmbedServer {
	dist, err := fs.Sub(FrontendAssets, "ui/dist")
	if err != nil {
		panic(err)
	}
	root := http.FS(dist)


	s := &EmbedServer{
		dist:     dist,
		fsServer: http.FileServer(root),
	}
	return s
}

// FaviconHandler favicon处理
func (e *EmbedServer) FaviconHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = "/favicon.ico"

	// 添加缓存
	w.Header().Set("Content-Type", "image/x-icon")
	w.Header().Set("Cache-Control", "max-age=86400, public")

	e.fsServer.ServeHTTP(w, r)
}

// StaticFileHandler 静态文件处理
func (e *EmbedServer) StaticFileHandler(prefix string) http.Handler {
	return e.fsServer
}
