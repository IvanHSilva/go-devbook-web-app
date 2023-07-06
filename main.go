package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/router"
)

func main() {

	fmt.Println("Rodando apliação web...escutando na porta 3000")

	r := router.GenerateRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}
