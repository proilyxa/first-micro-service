package router

import (
	"education-project/internal/http/controllers"
	"education-project/internal/http/middlewares"
	"education-project/internal/pkg/logger/sl"
	"education-project/react"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"golang.org/x/exp/slog"
	"io/fs"
	"log"
	"net/http"
	"strings"
	"time"
)

func NewRouter(
	log *slog.Logger,
	authController *controllers.AuthController,
) *chi.Mux {
	r := chi.NewRouter()
	initMiddlewares(r, log)
	initStaticFileServer(r)
	initApiRoutes(
		r,
		authController,
	)

	return r
}

func initMiddlewares(r *chi.Mux, log *slog.Logger) {
	r.Use(middlewares.HeaderInjection)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(sl.ContextWithLogger(log))
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           68100,
	}))
}

func initStaticFileServer(r *chi.Mux) {
	staticFs, err := fs.Sub(react.StaticFiles, "dist")
	if err != nil {
		log.Fatal(err)
	}
	httpFS := http.FileServer(http.FS(staticFs))
	r.Handle("/assets/*", httpFS)

	// index page
	r.HandleFunc("/*", indexHandler)
}

func indexHandler(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(resp, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	if strings.HasPrefix(req.URL.Path, "/api") {
		http.NotFound(resp, req)
		return
	}

	if req.URL.Path == "/favicon.svg" {
		rawFile, _ := react.StaticFiles.ReadFile("dist/vite.svg")
		resp.Write(rawFile)
		return
	}

	rawFile, _ := react.StaticFiles.ReadFile("dist/index.html")
	resp.Write(rawFile)
}
