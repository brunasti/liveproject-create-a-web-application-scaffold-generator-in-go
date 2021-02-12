package main

import (
	"fmt"
	"log"
	"net/http"
)

func serveContent(writer http.ResponseWriter, request *http.Request) {
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

	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", healthCheck)
	mux.HandleFunc("/", serveContent)

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
