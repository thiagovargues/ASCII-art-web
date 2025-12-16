package main

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type PageData struct {
	Text   string
	Banner string
	Result string
	Error  string
}

const templatesDir = "templates"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/ascii-art", handleAsciiArt)

	addr := listenAddr()
	log.Printf("Starting ASCII Art server on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r) // 404
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405 (not required by subject, but correct)
		return
	}

	renderPage(w, http.StatusOK, PageData{Banner: "standard"})
}

func handleAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		return
	}

	if err := r.ParseForm(); err != nil {
		renderPage(w, http.StatusBadRequest, PageData{
			Banner: "standard",
			Error:  "Bad request: cannot parse form",
		})
		return
	}

	text := r.FormValue("text")
	bannerName := r.FormValue("banner")
	if bannerName == "" {
		bannerName = "standard"
	}

	// 400 rules
	if text == "" {
		renderPage(w, http.StatusBadRequest, PageData{
			Text:   text,
			Banner: bannerName,
			Error:  "Bad request: text is required",
		})
		return
	}
	if bannerName != "standard" && bannerName != "shadow" && bannerName != "thinkertoy" {
		renderPage(w, http.StatusBadRequest, PageData{
			Text:   text,
			Banner: bannerName,
			Error:  "Bad request: unknown banner",
		})
		return
	}

	bannerFile := bannerName + ".txt"
	banner, err := LoadBanner(bannerFile)
	if err != nil {
		// subject wants 404 if banners not found
		if errors.Is(err, os.ErrNotExist) {
			renderPage(w, http.StatusNotFound, PageData{
				Text:   text,
				Banner: bannerName,
				Error:  "Not found: banner file missing",
			})
			return
		}
		// other banner errors
		renderPage(w, http.StatusInternalServerError, PageData{
			Text:   text,
			Banner: bannerName,
			Error:  "Internal error: cannot load banner",
		})
		return
	}

	result := RenderASCII(text, banner)

	renderPage(w, http.StatusOK, PageData{
		Text:   text,
		Banner: bannerName,
		Result: result,
	})
}

func renderPage(w http.ResponseWriter, status int, data PageData) {
	tplPath := filepath.Join(templatesDir, "index.html")

	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		// subject wants 404 if templates not found
		if errors.Is(err, os.ErrNotExist) {
			http.Error(w, "Not found: template missing", http.StatusNotFound)
			return
		}
		renderFallback(w, http.StatusInternalServerError, "template parse failed")
		return
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		renderFallback(w, status, "template render failed")
		return
	}

	w.WriteHeader(status)
	_, _ = w.Write(buf.Bytes())
}

// renderFallback emits a basic error response when the template cannot render.
func renderFallback(w http.ResponseWriter, status int, message string) {
	// Default to 500 if the provided status is not an error code.
	if status < 400 || status > 599 {
		status = http.StatusInternalServerError
	}

	switch status {
	case http.StatusBadRequest:
		message = "Bad request: " + message
	case http.StatusInternalServerError:
		message = "Internal error: " + message
	}

	http.Error(w, message, status)
}

// listenAddr returns an address from $PORT or falls back to :8080.
func listenAddr() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	return port
}
