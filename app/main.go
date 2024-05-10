package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})
	e.POST("/redis", postRedis)
	e.GET("/redis", getRedis)
	e.GET("/routes", getRoutes(e))

	e.Logger.Info(e.Routes())
	e.Logger.Fatal(e.Start(":1323"))
}

func postRedis(c echo.Context) error {
	ctx := context.Background()
	key := c.QueryParam("key")
	val := c.QueryParam("val")
	if err := rdb.Set(ctx, key, val, 0).Err(); err != nil {
		return c.String(http.StatusInternalServerError, "ng")
	}
	return c.String(http.StatusOK, "ok")
}

func getRedis(c echo.Context) error {
	ctx := context.Background()
	key := c.QueryParam("key")
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return c.String(http.StatusInternalServerError, "ng")
	}
	return c.String(http.StatusOK, val)
}

func getRoutes(e *echo.Echo) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, e.Routes())
	}
}
