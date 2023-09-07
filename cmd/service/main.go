package main

import (
	"context"
	"education-project/database"
	"education-project/internal/config"
	"education-project/internal/http/controllers"
	"education-project/internal/pkg/logger/handlers"
	"education-project/internal/pkg/logger/sl"
	"education-project/internal/repositories"
	"education-project/internal/router"
	"education-project/internal/services"
	"flag"
	"golang.org/x/exp/slog"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// config
	cfg := config.MustLoad(getConfigPath())

	// logger
	log := setUpLogger(cfg.Env)
	log.Info("Starting... ENV: " + cfg.Env)

	// database
	storage := setUpStorage(log)

	// repositories
	userRepo := repositories.NewUserRepository(storage)
	authTokenRepo := repositories.NewAuthTokenRepository(storage)

	// services
	userService := services.NewUserServiceImpl(
		userRepo,
		authTokenRepo,
	)

	// controllers
	authController := controllers.NewAuthController(userService)

	// router
	r := router.NewRouter(
		log,
		authController,
	)

	srv := setUpServer(cfg, r)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Warn(err.Error())
		}
	}()
	log.Info("Server started on " + cfg.Host + ":" + cfg.Port)

	<-done

	storage.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Failed to stop srv", sl.Err(err))

		return
	}

	log.Info("Server stopped")
}

func setUpStorage(log *slog.Logger) *database.Storage {
	_, err := fs.Sub(database.StaticFiles, "migrations")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	dbCon, err := database.NewSqliteConnection(log)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	return dbCon
}

func getConfigPath() string {
	var cfgPath = flag.String("config", "../../config/config.yaml", "test")
	flag.Parse()

	return *cfgPath
}

func setUpLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = setupPrettySlog()
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return logger
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

func setUpServer(cfg *config.Config, r http.Handler) *http.Server {
	srv := &http.Server{
		Addr:         cfg.Host + ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	return srv
}
