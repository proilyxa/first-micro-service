package router

import (
	"education-project/internal/http/controllers"
	"education-project/internal/http/middlewares"
	"education-project/internal/pkg/logger/sl"
	"education-project/web"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

}

func initStaticFileServer(r *chi.Mux) {
	staticFs, err := fs.Sub(web.StaticFiles, "dist")
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
		rawFile, _ := web.StaticFiles.ReadFile("dist/favicon.svg")
		resp.Write(rawFile)
		return
	}

	rawFile, _ := web.StaticFiles.ReadFile("dist/index.html")
	resp.Write(rawFile)
}
