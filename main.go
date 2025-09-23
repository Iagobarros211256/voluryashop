package voluryashop

import (
	"log"

	"github.com/Iagobarros211256/voluryashop/repository"
	"github.com/Iagobarros211256/voluryashop/routes"
	"github.com/go-delve/delve/pkg/config"
)

func main() {
	config.LoadConfig()
	repository.ConnDB()

	r := routes.SetupRouter()
	log.Fatal(r.Run(":" + config.GetEnv("PORT")))
}
