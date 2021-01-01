package cmd

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"

	"github.com/anvari1313/photopic-server/config"
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
	_, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())

	e.GET("/list", routes.ListHandler)
	e.Static("/thumbnail", config.C.BasePath)
	e.Static("/static", config.C.BasePath)

	if err := e.Start(config.C.Address); err != nil {
		log.Fatal(err)
	}
}
