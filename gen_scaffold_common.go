package main

import (
	"fmt"
)

func generateScaffoldCommon(structureType projectStructureType) int {
	logVerbose(fmt.Sprintf("Generating scaffold common content for project %s in %s", structureType.Name, structureType.Path))

	writeFile("./"+structureType.Path+"/handlers/go.mod", goModTmpl)
	writeFile("./"+structureType.Path+"/handlers/handlers.go", handlerGoTmpl)

	//t, _ := template.ParseFiles("handlers/constants.go")
	//t.Execute(w, structureType.Name)
	//writeFile("./"+structureType.Path+"/handlers/constants.go", constantsGoTmpl)
	writeTemplate("handlers/constants.go", "./"+structureType.Path+"/handlers/constants.go", structureType.Name)

	return 0
}

var goModTmpl = `module brunasti.it/handlers

go 1.16
`

//var constantsGoTmpl = `package main
//
//const appName = "{{ . }}"
//`

var handlerGoTmpl = `package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlers(writer http.ResponseWriter, request *http.Request) {
	log.Printf("/ [%v]\n", request.URL.Path[1:])
	n, err := fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
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
