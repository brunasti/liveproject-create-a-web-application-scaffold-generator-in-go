package main

import (
	"fmt"
)

func generateScaffoldCommon(structureType projectStructureType) int {
	logVerbose(fmt.Sprintf("Generating scaffold common content for project %s in %s", structureType.Name, structureType.Path))

	//writeFile("./"+structureType.Path+"/handlers/go.mod", goModTmpl)
	//writeFile("./"+structureType.Path+"/handlers/handlers.go", handlerGoTmpl)

	writeTemplate("handlers/handlers.go", "./"+structureType.Path+"/handlers/handlers.go", structureType)
	writeTemplate("handlers/constants.go", "./"+structureType.Path+"/handlers/constants.go", structureType)
	writeTemplate("handlers/go.mod", "./"+structureType.Path+"/handlers/go.mod", structureType)

	if structureType.StaticAssets {
		copyFile("handlers/statics/index.html", "./"+structureType.Path+"/handlers/statics/index.html")
	}
	return 0
}

//var goModTmpl = `module brunasti.it/handlers
//
//go 1.16
//`

var handlerGoTmpl = `package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlers(writer http.ResponseWriter, request *http.Request) {
	log.Printf("/ [%v]\n", request.URL.Path[1:])
	n, err := fmt.Fprintf(writer, "Hello World, from %s! - [%s]", appName, request.URL.Path[1:])
	if err != nil {
		log.Printf("/ [%v] - Error processing [%v][%d]\n", request.URL.Path[1:], err, n)
	}
}

func healthCheck(writer http.ResponseWriter, request *http.Request) {
	log.Printf("/healthcheck [%v]\n", request.URL.Path[1:])
	n, err := fmt.Fprintf(writer, "ok")
	if err != nil {
		log.Printf("/healthcheck - Error processing [%v][%d]\n", err, n)
	}
}

func Application() {
	fmt.Println(appName + " ------------")
	http.HandleFunc("/", handlers)
	http.HandleFunc("/healthcheck", healthCheck)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(appName + " -------- end")
}

func main() {
	Application()
}
`
