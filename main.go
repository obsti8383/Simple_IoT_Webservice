package main

import (
	"log"
	"net/http"

	scribble "github.com/nanobox-io/golang-scribble"
)

var IoTDB *scribble.Driver

func main() {
	// a new scribble driver, providing the directory where it will be writing to,
	// and a qualified logger if desired
	var err error
	IoTDB, err = scribble.New("IoT_DB", nil)
	HandleError(err)

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8888", router))
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
