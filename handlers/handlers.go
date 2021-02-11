package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

func serveContent(env *Env, writer http.ResponseWriter, request *http.Request) {
	log.Printf("/ [%v]\n", request.URL.Path[1:])
{{if .StaticAssets}}
	p := request.URL.Path
	if p == "/" {
		p = "./statics/index.html"
	} else {
		p = "./statics" + p
	}

	log.Printf("Serving content [%v]\n", p)
	http.ServeFile(writer, request, p)
{{- else}}
	n, err := fmt.Fprintf(writer, "Hello World, from %s! - [%s]", appName, request.URL.Path[1:])
	if err != nil {
		log.Printf("/ [%v] - Error processing [%v][%d]\n", request.URL.Path[1:], err, n)
	}
{{- end}}
}

func Application() {
	fmt.Println(appName + " ------------")
	fmt.Println(" Serving at on port {{.Port}}")
	fmt.Println(" ------------")

	env := &Env{
		Port: {{.Port}},
		Host: "localhost",
		// We might also have a custom log.Logger, our
		// template instance, and a config struct as fields
		// in our Env struct.
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", healthCheck)
	mux.HandleFunc("/", Handler{env, serveContent})

	server := &http.Server{
		Addr:    "0.0.0.0:{{.Port}}",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(appName + " -------- end")
}

func main() {
	Application()
}

func healthCheck(writer http.ResponseWriter, request *http.Request) {
	log.Printf("/healthcheck [%v]\n", request.URL.Path[1:])
	n, err := fmt.Fprintf(writer, "ok")
	if err != nil {
		log.Printf("/healthcheck - Error processing [%v][%d]\n", err, n)
	}
}


// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// A (simple) example of our application-wide configuration.
type Env struct {
	Port string
	Host string
}

// The Handler struct that takes a configured Env and a function matching
// our useful signature.
type Handler struct {
	*Env
	H func(e *Env, w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}