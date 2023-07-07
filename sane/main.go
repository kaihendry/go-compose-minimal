package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	e := echo.New()

	e.Use(slogecho.New(logger))
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to sanity!")
	})

	err := e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))

	if err != nil {
		logger.Error("failed to listen and serve", err)
	}

}
