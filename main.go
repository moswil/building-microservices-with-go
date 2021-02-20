package main

import (
	"log"
	"net/http"
	"os"

	"github.com/moswil/building-microservices-with-go/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.HandleFunc("/", hh.ServeHTTP)

	http.ListenAndServe(":9090", sm)
}
