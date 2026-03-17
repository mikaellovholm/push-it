# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Run

```bash
go build -o push-it .        # build binary
./push-it                     # run locally on :8080
```

No external dependencies — pure standard library.

## Deploy

Uses [ko](https://ko.build/) to build and deploy to Google Cloud Run (no Dockerfile needed):

```bash
KO_DOCKER_REPO=europe-north1-docker.pkg.dev/claude-dev-ml-01/claude-dev-services/push-it \
  ko build . --bare --tags=latest

gcloud run deploy push-it \
  --image=europe-north1-docker.pkg.dev/claude-dev-ml-01/claude-dev-services/push-it:latest \
  --region=europe-north1 --port=8080 --allow-unauthenticated --project=claude-dev-ml-01
```

Live URL: https://push-it-98638225257.europe-north1.run.app

## Architecture

Single-file Go HTTP server (`main.go`) that displays a random name from a hardcoded list. The HTML template (`templates/index.html`) is embedded into the binary via `//go:embed` — this is required because ko uses a distroless base image with no filesystem. Any new templates or static assets must also be embedded.
