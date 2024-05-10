package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
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
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})
	e.POST("/redis", func(c echo.Context) error {
		ctx := context.Background()
		key := c.QueryParam("key")
		val := c.QueryParam("val")
		if err := rdb.Set(ctx, key, val, 0).Err(); err != nil {
			return c.String(http.StatusInternalServerError, "ng")
		}
		return c.String(http.StatusOK, "ok")
	})
	e.GET("/redis", func(c echo.Context) error {
		ctx := context.Background()
		key := c.QueryParam("key")
		val, err := rdb.Get(ctx, key).Result()
		if err != nil {
			return c.String(http.StatusInternalServerError, "ng")
		}
		return c.String(http.StatusOK, val)
	})
	e.GET("/routes", func(c echo.Context) error {
		return c.JSON(http.StatusOK, e.Routes())
	})

	e.Logger.Info(e.Routes())
	e.Logger.Fatal(e.Start(":1323"))
}
