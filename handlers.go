package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	mux "github.com/julienschmidt/httprouter"
	//	scribble "github.com/nanobox-io/golang-scribble"
)

func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	fmt.Fprintf(w, "<h1 style=\"font-family: Helvetica;\">Hello, welcome to IoT REST web service</h1>")
}

func GetDatasets(w http.ResponseWriter, r *http.Request, _ mux.Params) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var test = IoTDataset{12828292, 9292, "ordered", "", "15.05.2019"}

	if err := json.NewEncoder(w).Encode(test); err != nil {
		panic(err)
	}
}

func GetDatasetId(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println(r.RequestURI)
	//fmt.Println(r.Header)
	//fmt.Println(r.Context())

	HandleError(err)

	var test = IoTDataset{id, 9191, "delivered", "", "16.05.2019"}

	if err := json.NewEncoder(w).Encode(test); err != nil {
		panic(err)
	}
}

func PostDataset(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	// Read request body and close it
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	HandleError(err)
	defer r.Body.Close()

	// Save JSON to Post struct
	var dataset IoTDataset
	if err := json.Unmarshal(body, &dataset); err != nil {
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// write to file db
	//Database, err := scribble.New("IoT_DB", nil)
	//HandleError(err)

	err = IoTDB.Write("fish", strconv.Itoa(dataset.ID), dataset)
	HandleError(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dataset); err != nil {
		panic(err)
	}

	fmt.Println("ID: ", dataset.ID)
	fmt.Println("Location: ", dataset.Location)
}
