package main

import (
	// Go Internal Packages
	"context"
	"os"
	"os/signal"
	"syscall"
	// Local Packages
	config "learn-go/config"
	xhttp "learn-go/http"
	handlers "learn-go/http/handlers"
	xhmodels "learn-go/models/xhandlers"
	mongodb "learn-go/repositories/mongodb"
	health "learn-go/services/health"
	students "learn-go/services/students"

	// External Packages
	"github.com/alecthomas/kingpin/v2"
	_ "github.com/jsternberg/zap-logfmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/rawbytes"
	"go.uber.org/zap"
)

// InitializeServer sets up an HTTP server with defined handlers.
// Repositories are initialized, creates the services, and subsequently constructs
// handlers for the services
func InitializeServer(ctx context.Context, k config.Config, logger *zap.Logger) (*xhttp.Server, error) {
	// Mongo Connection
	client, err := mongodb.Connect(ctx, k.Mongo.MetaURI)
	if err != nil {
		return nil, err
	}

	// creating cache with default expiration, individual expiry can be set to keys later on.
	// appCache := cache.New(5*time.Minute, 10*time.Minute)

	// Init repos, services && handlers
	studentsRepo := mongodb.NewStudentsRepository(client)

	healthSvc := health.NewService(logger, client)
	studentsSvc := students.NewService(studentsRepo)

	studentsHandler := handlers.NewSegmentsHandler(studentsSvc)

	xHandlers := xhmodels.XHandlers{
		StudentsHandlers: studentsHandler,
	}

	server := xhttp.NewServer(k.Prefix, logger, &xHandlers, healthSvc)
	return server, nil
}

// LoadConfig loads the default configuration and overrides it with the config file
// specified by the path defined in the config flag
func LoadConfig() *koanf.Koanf {

	confifPathMsg := "Path to the application config file"
	configPath := kingpin.Flag("config", confifPathMsg).Short('c').Default("config.yml").String()

	kingpin.Parse()

	k := koanf.New(".")
	_ = k.Load(rawbytes.Provider(config.DefaultConfig), yaml.Parser())
	if *configPath != "" {
		_ = k.Load(file.Provider(*configPath), yaml.Parser())
	}

	return k
}

func main() {
	k := LoadConfig()
	appKonf := config.Config{}
	k.Unmarshal("", &appKonf)

	if !appKonf.IsProdMode {
		k.Print()
	}

	// FIX ME: Rewrite this logger config section
	cfg := zap.NewProductionConfig()
	cfg.Encoding = "logfmt"
	_ = cfg.Level.UnmarshalText([]byte(k.String("logger.level")))
	cfg.InitialFields = make(map[string]any)
	cfg.InitialFields["host"], _ = os.Hostname()
	cfg.InitialFields["service"] = "learn-go"
	cfg.OutputPaths = []string{"stdout"}
	logger, _ := cfg.Build()
	defer func() {
		_ = logger.Sync()
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv, err := InitializeServer(ctx, appKonf, logger)
	if err != nil {
		logger.Fatal("Cannot initialize server", zap.Error(err))
	}
	if err := srv.Listen(ctx, k.String("listen")); err != nil {
		logger.Fatal("Cannot listen", zap.Error(err))
	}
}
