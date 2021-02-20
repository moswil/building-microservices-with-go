package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Oops!", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		log.Println("Status OK")
		fmt.Fprintf(w, "Hello, %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye!")
	})

	http.ListenAndServe(":9090", nil)
}
