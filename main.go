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

// func init() {
// 	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(hashKey)
// 	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(blockKey)
// }

func main() {

	fmt.Println()
	config.LoadConfig()
	cookies.ConfigureCookie()
	utils.LoadTemplates()
	r := router.GenerateRouter()

	fmt.Printf("Escutando na porta %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
