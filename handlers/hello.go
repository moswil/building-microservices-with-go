package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello implements the DefaultServeMux
type Hello struct {
	l *log.Logger
}

// NewHello for logging purposes
func NewHello(l *log.Logger) *Hello {
	return &Hello{
		l,
	}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oops!", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	h.l.Println("Status OK")
	fmt.Fprintf(w, "Hello, %s\n", d)
}
