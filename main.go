package main

import (
	"fmt"
	"log"

	"github.com/fazriegi/fga_assignment2/config"
	"github.com/fazriegi/fga_assignment2/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	viper := config.NewViper()
	db := config.NewDatabase(viper)

	app := gin.Default()
	routes.NewRouter(db, app)
	port := viper.GetString("web.port")

	log.Fatal(app.Run(fmt.Sprintf(":%s", port)))
}
