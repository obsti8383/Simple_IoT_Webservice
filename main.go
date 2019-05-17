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

/*

// Write a fish to the database
fish := Fish{}
if err := db.Write("fish", "onefish", fish); err != nil {
  fmt.Println("Error", err)
}

// Read a fish from the database (passing fish by reference)
onefish := Fish{}
if err := db.Read("fish", "onefish", &onefish); err != nil {
  fmt.Println("Error", err)
}

// Read all fish from the database, unmarshaling the response.
records, err := db.ReadAll("fish")
if err != nil {
  fmt.Println("Error", err)
}

fishies := []Fish{}
for _, f := range records {
  fishFound := Fish{}
  if err := json.Unmarshal([]byte(f), &fishFound); err != nil {
    fmt.Println("Error", err)
  }
  fishies = append(fishies, fishFound)
}

// Delete a fish from the database
if err := db.Delete("fish", "onefish"); err != nil {
  fmt.Println("Error", err)
}

// Delete all fish from the database
if err := db.Delete("fish", ""); err != nil {
  fmt.Println("Error", err)
}
*/
