package main

import (
	"context"
	"fmt"
	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
	"github.com/magmel48/social-network/internal/api"
	"github.com/magmel48/social-network/internal/config"
	"github.com/magmel48/social-network/internal/db"
	"github.com/magmel48/social-network/internal/repositories/users"
	"github.com/magmel48/social-network/internal/server"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type HealthcheckResponse struct {
	MySQL int `json:"mysql"`
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGSTOP, syscall.SIGTERM)
	defer cancel()

	// fetch environment variables
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	// register logger
	logLevel, err := zap.ParseAtomicLevel(cfg.LogLevel)
	if err != nil {
		panic(err)
	}

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stdout,
		logLevel,
	))
	defer func() {
		err = logger.Sync()
		if err != nil {
			log.Println(err)
		}
	}()

	// open db connection and configure store
	database, err := db.Open(
		fmt.Sprintf(
			"%s:%s@tcp(db:%d)/%s", cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.Port, cfg.MySQL.Database))
	if err != nil {
		panic(err)
	}

	queries := db.New(database)
	repository := users.New(queries, database)

	// register healthcheck
	r := gin.Default()
	r.GET("/hc", func(c *gin.Context) {
		response := HealthcheckResponse{}

		err = database.PingContext(c)
		if err != nil {
			response.MySQL = 1
		}

		c.JSON(http.StatusOK, response)
	})

	// register endpoints
	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}

	r.Use(middleware.OapiRequestValidator(swagger))
	api.RegisterHandlers(r, server.New(repository, logger))

	// start the server
	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
	}

	logger.Fatal(s.ListenAndServe().Error())
}
