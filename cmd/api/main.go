package main

import (
	"context"
	"fmt"
	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
	"github.com/magmel48/social-network/internal/api"
	"github.com/magmel48/social-network/internal/config"
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

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGSTOP, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

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

	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(middleware.OapiRequestValidator(swagger))
	api.RegisterHandlers(r, server.New())

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
	}

	logger.Fatal(s.ListenAndServe().Error())
}
