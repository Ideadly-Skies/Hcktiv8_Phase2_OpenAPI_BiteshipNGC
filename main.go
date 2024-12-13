package main

import (
	"w3/ngc/config"
	"w3/ngc/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){
	cfg := config.LoadConfig()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/couriers", func (c echo.Context) error {
		return handler.GetCouriers(c, cfg)	
	})

	e.POST("/shipping-cost", func (c echo.Context) error {
		return handler.CalculateShippingCost(c, cfg)	
	})

	e.Logger.Fatal(e.Start(":8080"))
}