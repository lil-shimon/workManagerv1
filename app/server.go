package main

import (
    "net/http"
    "github.com/labstack/echo"
    "github.com/lil-shimon/workManagerV2/app/handler/type"
)

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello golang world")
    })

    // Routes
    e.POST("/store/type", type.CreateType)
    e.Logger.Fatal(e.Start(":1323"))
}