package main

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	IP string `json:"ip"`

	UserAgent string `json:"user_agent"`

	Host string `json:"host"`
}

func handler(c echo.Context) error {
	ip := c.RealIP()
	userAgent := c.Request().UserAgent()
	host := c.Request().Host

	return c.JSON(200, Response{
		IP:        ip,
		UserAgent: userAgent,
		Host:      host,
	})
}

func main() {
	e := echo.New()

	e.IPExtractor = echo.ExtractIPDirect()

	e.GET("/", handler)
	e.Logger.Fatal(e.Start(":1323"))
}
