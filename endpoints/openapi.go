package endpoints

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func RouteOpenAPI(w http.ResponseWriter, r *http.Request) {

	openAPIConfig, _ := ioutil.ReadFile("./config/openapi.json")

	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, string(openAPIConfig))

	log.Println("GET /")
}
