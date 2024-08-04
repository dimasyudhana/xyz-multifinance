package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dimasyudhana/xyz-multifinance/app/config"
	"github.com/dimasyudhana/xyz-multifinance/app/middlewares"
	"github.com/dimasyudhana/xyz-multifinance/app/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	certFile        = "ssl/server.crt"
	keyFile         = "ssl/server.pem"
	maxHeaderBytes  = 1 << 20
	stackSize       = 1 << 10
	csrfTokenHeader = "X-CSRF-Token"
	bodyLimit       = "2M"
)

type Server struct {
	db   *sql.DB
	echo *echo.Echo
	cfg  *config.AppConfig
}

func NewServer(db *sql.DB, echo *echo.Echo, cfg *config.AppConfig) *Server {
	return &Server{
		db:   db,
		echo: echo,
		cfg:  cfg,
	}
}
func (s *Server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := log.New(os.Stderr, "INFO: ", log.LstdFlags)
	app.Println("Starting XYZ-Multifinance App")
	app.Printf("AppVersion: %s", s.cfg.HttpServer.AppVersion)

	e := s.echo
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID, csrfTokenHeader},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         stackSize,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.BodyLimit(bodyLimit))
	e.Pre(middleware.RemoveTrailingSlash())
	e.HTTPErrorHandler = middlewares.ProblemDetailsHandler

	router.InitRouter(s.db, s.echo)

	go func() {
		app.Printf("Metrics server is running on port: %s", s.cfg.Metrics.Port)
		if err := http.ListenAndServe(s.cfg.Metrics.Port, nil); err != nil {
			log.Printf("Metrics server error: %v", err)
			cancel()
		}
	}()

	go func() {
		app.Printf("Server is listening on PORT: %s", s.cfg.HttpServer.Port)
		e.Server.ReadTimeout = time.Second * time.Duration(s.cfg.HttpServer.ReadTimeout)
		e.Server.WriteTimeout = time.Second * time.Duration(s.cfg.HttpServer.WriteTimeout)
		e.Server.MaxHeaderBytes = maxHeaderBytes

		if err := e.StartTLS(s.cfg.HttpServer.Port, certFile, keyFile); err != nil {
			log.Fatalf("Error starting TLS Server: %v", err)
		}
	}()

	go func() {
		app.Printf("Starting Debug Server on PORT: %s", s.cfg.HttpServer.PprofPort)
		if err := http.ListenAndServe(s.cfg.HttpServer.PprofPort, nil); err != nil {
			log.Printf("Error starting PPROF server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Printf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Printf("ctx.Done: %v", done)
	}

	if err := s.echo.Server.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("Server Exited Properly")
	return nil
}
