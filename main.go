package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bketelsen/teachweb/actions"
	"github.com/gobuffalo/envy"
)

func main() {
	port := envy.Get("PORT", "3000")
	log.Printf("Starting teachweb on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), actions.App()))
}
