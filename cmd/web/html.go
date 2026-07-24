package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"path"
	"time"

	"github.com/jlevesy/easyplu/assets"
)

type htmlRenderer struct {
	templateFS      fs.FS
	sharedTemplates *template.Template
	logger          *slog.Logger
}

func newHTMLRenderer(templateFS fs.FS, logger *slog.Logger, sharedTemplateFiles ...string) (*htmlRenderer, error) {
	funcs := template.FuncMap{
		"now":   time.Now,
		"asset": assetPath,
	}

	sharedTemplates, err := template.New("").Funcs(funcs).ParseFS(templateFS, sharedTemplateFiles...)
	if err != nil {
		return nil, err
	}

	return &htmlRenderer{
		templateFS:      templateFS,
		sharedTemplates: sharedTemplates,
		logger:          logger,
	}, nil
}

func (r *htmlRenderer) render(w http.ResponseWriter, status int, data any, templateName string, additionalTemplateFiles ...string) {
	if err := r.renderWithError(w, status, data, templateName, additionalTemplateFiles...); err != nil {
		r.logger.Error("Could not render template", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (r *htmlRenderer) renderWithError(w http.ResponseWriter, status int, data any, templateName string, additionalTemplateFiles ...string) error {
	ts, err := r.sharedTemplates.Clone()
	if err != nil {
		return err
	}

	if len(additionalTemplateFiles) > 0 {
		ts, err = ts.ParseFS(r.templateFS, additionalTemplateFiles...)
		if err != nil {
			return err
		}
	}

	var buf bytes.Buffer

	if err := ts.ExecuteTemplate(&buf, templateName, data); err != nil {
		return err
	}

	w.Header().Set("Vary", "HX-Request")
	w.WriteHeader(status)
	if _, err := buf.WriteTo(w); err != nil {
		return err
	}

	return nil
}

func assetPath(in string) string {
	ent, ok := assets.Manifest[in]
	if !ok {
		panic(fmt.Sprintf("Asset %s not found", in))
	}

	return path.Join("/static", ent.File)
}
