package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/apex/gateway/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	if _, ok := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME"); ok {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(slogecho.New(logger))
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to sanit!")
	})

	e.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	if _, ok := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME"); ok {
		gateway.ListenAndServe("", e)
	} else {
		slog.Info("local development", "port", os.Getenv("PORT"))
		e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
	}

}
