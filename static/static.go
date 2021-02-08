package static

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println(appName + " ------------")

	mux := http.NewServeMux()

	mux.HandleFunc("/healthcheck", healthCheck)

	mux.HandleFunc("/", serveContent)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(appName + " -------- end")

}

func healthCheck(writer http.ResponseWriter, request *http.Request) {
	log.Printf("/healthcheck [%v]\n", request.URL.Path[1:])
	n, err := fmt.Fprintf(writer, "ok")
	if err != nil {
		log.Printf("/healthcheck - Error processing [%v][%d]\n", err, n)
	}
}

func serveContent(writer http.ResponseWriter, request *http.Request) {
	log.Printf("[%v]\n", request.URL.Path)
	p := request.URL.Path
	if p == "/" {
		p = "./public/index.html"
	} else {
		p = "./public" + p
	}

	log.Printf("Serving content [%v]\n", p)
	http.ServeFile(writer, request, p)
}
