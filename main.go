package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/config"
	"web/src/router"
	"web/src/utils"
)

func main() {

	config.LoadConfig()
	utils.LoadTemplates()
	r := router.GenerateRouter()

	fmt.Printf("Escutando na porta %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
