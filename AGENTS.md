# AGENTS.md

## Project: EasyPlu

A lightweight Go web application scaffold with Vite-powered frontend assets and htmx for dynamic DOM updates.

## Directory Structure

```
├── cmd/web/                  # Go application entry point
│   ├── main.go               # Server setup, routing, HTTP handlers, listens on :5051
│   ├── handlers.go           # Route handlers (home, test/gopher)
│   └── html.go               # HTML template rendering engine with Vite manifest support
├── assets/                   # Frontend assets and embedded static files
│   ├── efs.go                # Go embed filesystem for html/ and static/ directories
│   ├── html/                 # Go HTML templates (embedded at compile time)
│   │   ├── base.tmpl         # Base HTML5 layout shell with block slots
│   │   ├── pages/
│   │   │   └── home.tmpl     # Home page template
│   │   └── partials/
│   │       ├── images.tmpl   # Gopher image partial
│   │       └── skeleton.tmpl # Placeholder partial
│   ├── src/                  # Uncompiled source assets
│   │   ├── css/style.scss    # SCSS styles (empty)
│   │   ├── img/gopher.png    # Go gopher mascot
│   │   └── js/layout.js      # htmx import (window.htmx)
│   └── static/               # Vite-compiled output (hashed filenames + manifest)
│       └── assets-manifest.json
├── go.mod                    # Go module: github.com/jlevesy/easyplu (Go 1.26.4, zero external deps)
├── package.json              # NPM scripts: serve, build, build-watch, lint, fmt
└── vite.config.js            # Vite config: root=assets/, output=static/
```

## Tech Stack

| Layer | Technology |
|---|---|
| Backend | Go 1.26.4 (standard library `net/http`, `slog`) |
| Frontend build | Vite 8 |
| Frontend runtime | htmx.org 2 |
| Styles | Sass (sass-embedded) |
| Templating | Go `html/template` (embedded via `//go:embed`) |

## Key Concepts

**Static embedding.** `assets/efs.go` uses `//go:embed` to embed `html/` and `static/` into the Go binary. All templates and static files ship inside the compiled binary.

**Asset resolution.** `assets/html/base.tmpl` defines a custom `asset` template function backed by Vite's `assets-manifest.json` (loaded at init time in `efs.go`). This provides hash-busted URLs for CSS/JS/images.

**htmx integration.** The home page renders a button with `hx-get="/gopher"` that triggers a partial request. The `/gopher` handler returns a rendered image partial. The renderer sets `Vary: HX-Request` for htmx compatibility.

## Running

```bash
npm run serve        # Build assets + start Go server on :5051
npm run build        # Build frontend assets to assets/static/
npm run build-watch  # Build and watch for changes
```

The Go server runs on **port 5051** with routes:

| Route | Handler | Description |
|---|---|---|
| `GET /` | `home` | Home page |
| `GET /gopher` | `test` | Gopher image partial (htmx demo) |
| `GET /static/*` | file server | Serves embedded static assets |

## Development Notes

- No linting config exists on disk despite `package.json` scripts referencing `eslint` and `prettier` for a `./frontend` directory that does not exist.
- No tests exist.
- No Makefile, Dockerfile, or CI/CD configuration.
- No README.md.
- Go has zero external dependencies.
- The `./frontend` directory referenced in `package.json` scripts does not exist; assets live under `assets/` instead.
