package main

import (
    // "net/http"
    "github.com/labstack/echo"
    "github.com/lil-shimon/workManager/handler"
)

func main() {
    e := echo.New()
    // e.GET("/", func(c echo.Context) error {
        // return c.String(http.StatusOK, "Hello golang world")
    // })

    // Routes
    e.POST("/", handler.CreateType)
    e.Logger.Fatal(e.Start(":1323"))
}