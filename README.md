Ascii Art Web
=============

Small Go HTTP server that renders ASCII art from text using the `standard`, `shadow`, or `thinkertoy` banners.

Requirements
------------
- Go 1.23+
- No external Go modules (standard library only; see `go.mod`)
- Banner files in repo root: `standard.txt`, `shadow.txt`, `thinkertoy.txt`
- Template: `templates/index.html`

Run locally
-----------
```
go run .
# or
go build -o ascii-art-web
./ascii-art-web
```
The server listens on `$PORT` (default `:8080`).

Docker
------
- Build: `docker image build -f Dockerfile -t ascii-art-web-docker .`
- Run: `docker container run -p 8080:8080 --name dockerize ascii-art-web-docker`
- Inspect: `docker images` then `docker ps -a`, and `docker exec -it dockerize /bin/sh` to see `/app` contents (expected: `ascii-art-web`, `templates/`, `standard.txt`, `shadow.txt`, `thinkertoy.txt`, `style.css`).
- Metadata: the Dockerfile applies OCI labels and runs as a non-root `asciiart` user.
- Note: Docker commands could not be executed in this audit environment (no access to the Docker daemon); run the above locally to verify image/container creation.

How to use
----------
1) Open `http://localhost:8080`.  
2) Enter text and choose a banner.  
3) Click Generate; the result is rendered on the same page.

Static files
------------
- `style.css` is served at `/style.css` by the server.
- Banners are loaded once and cached in memory after first use.

Scripts
-------
- `scripts/docker-build-run.sh` builds the image and runs the container (defaults: image `ascii-art-web-docker`, container `dockerize`, port `8080`).

Tests
-----
- Run `go test ./...` to check banner loading, rendering, and HTTP endpoints.

Audit notes
-----------
- Allowed packages: only Go standard library is used.
- Dockerfile: present with labels, multi-stage build, non-root runtime.
- Build helper script: provided at `scripts/docker-build-run.sh`.
- Tests: basic coverage for banner loading, rendering, and HTTP endpoints via `go test ./...`.
- API: simple HTML form endpoints (`/` and `/ascii-art`), not a separate JSON API.
- Known gap: Docker build/run still needs to be executed locally (daemon unavailable in this audit environment).

License
-------
MIT (`LICENSE`).

Authors
-------
Alioune Sall  
Emilia Chedot  
Thiago Vargues
