package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"

	"aabbcc/internal/database"
	"aabbcc/internal/smtp"
	"aabbcc/internal/version"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	err := run(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

type config struct {
	baseURL  string
	httpPort int
	db       struct {
		dsn         string
		automigrate bool
	}
	jwt struct {
		secretKey string
	}
	notifications struct {
		email string
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		from     string
	}
}

type application struct {
	config config
	db     *database.DB
	logger *slog.Logger
	mailer *smtp.Mailer
	wg     sync.WaitGroup
}

func run(logger *slog.Logger) error {
	var cfg config

	flag.StringVar(&cfg.baseURL, "base-url", "http://localhost:4444", "base URL for the application")
	flag.IntVar(&cfg.httpPort, "http-port", 4444, "port to listen on for HTTP requests")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "user:pass@tcp(localhost:3306)/example?parseTime=true", "mySQL DSN")
	flag.BoolVar(&cfg.db.automigrate, "db-automigrate", true, "run migrations on startup")
	flag.StringVar(&cfg.jwt.secretKey, "jwt-secret-key", "kxg4srlrvpuwgathq3rwu5k5hlcv2uum", "secret key for JWT authentication")
	flag.StringVar(&cfg.notifications.email, "notifications-email", "", "contact email address for error notifications")
	flag.StringVar(&cfg.smtp.host, "smtp-host", "example.smtp.host", "smtp host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 25, "smtp port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "example_username", "smtp username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "pa55word", "smtp password")
	flag.StringVar(&cfg.smtp.from, "smtp-from", "Example Name <no-reply@example.org>", "smtp sender")

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	db, err := database.New(cfg.db.dsn, cfg.db.automigrate)
	if err != nil {
		return err
	}
	defer db.Close()

	mailer, err := smtp.NewMailer(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.from)
	if err != nil {
		return err
	}

	app := &application{
		config: cfg,
		db:     db,
		logger: logger,
		mailer: mailer,
	}

	return app.serveHTTP()
}
