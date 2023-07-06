package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/config"
	"web/src/cookies"
	"web/src/router"
	"web/src/utils"
)

func main() {

	fmt.Println()
	config.LoadConfig()
	cookies.ConfigureCookie()
	utils.LoadTemplates()
	r := router.GenerateRouter()

	fmt.Printf("Escutando na porta %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
