package cmd

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"

	"github.com/anvari1313/photopic-server/routes"
)

var serveCMD = &cobra.Command{
	Use:   "serve",
	Short: "serve API",
	Run: func(_ *cobra.Command, _ []string) {
		serve()
	},
}

func serve() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())

	e.GET("/list", routes.ListHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
