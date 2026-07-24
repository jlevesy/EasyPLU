package main

import (
	"log/slog"
	"net/http"
	"os"
	"path"

	"github.com/jlevesy/easyplu/assets"
)

type application struct {
	logger *slog.Logger
	html   *htmlRenderer
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	htmlRenderer, err := newHTMLRenderer(assets.HTMLFiles, logger, "base.tmpl", "partials/*.tmpl")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger: logger,
		html:   htmlRenderer,
	}

	fileserver := http.FileServerFS(assets.StaticFiles)

	mux := http.NewServeMux()

	mux.Handle("GET /static/", noAssetsManifets(http.StripPrefix("/static", fileserver)))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("POST /chat", app.chat)

	logger.Info("starting server", "port", 5051)

	err = http.ListenAndServe(":5051", mux)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func noAssetsManifets(h http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if path.Base(r.URL.Path) == "assets-manifest.json" {
			http.NotFound(rw, r)
			return
		}

		h.ServeHTTP(rw, r)
	}
}
