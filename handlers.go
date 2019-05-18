package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	mux "github.com/julienschmidt/httprouter"
)

const dbDataName = "IoT_Datasets"

func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	fmt.Fprintf(w, "<h1 style=\"font-family: Helvetica;\">Hello, welcome to IoT REST web service</h1>")
}

func GetDatasets(w http.ResponseWriter, r *http.Request, _ mux.Params) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	records, err := IoTDB.ReadAll(dbDataName)
	HandleError(err)
	entries := []IoTDataset{}
	for _, e := range records {
		dataset := IoTDataset{}
		if err := json.Unmarshal([]byte(e), &dataset); err != nil {
			fmt.Println("Error", err)
		} else {
			entries = append(entries, dataset)
		}
	}

	if err := json.NewEncoder(w).Encode(entries); err != nil {
		panic(err)
	}
}

func GetDatasetId(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	HandleError(err)

	var dataset IoTDataset
	err = IoTDB.Read(dbDataName, strconv.Itoa(id), &dataset)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(dataset); err != nil {
			panic(err)
		}
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

	err = IoTDB.Write(dbDataName, strconv.Itoa(dataset.ID), dataset)
	HandleError(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dataset); err != nil {
		panic(err)
	}

	IoTDB.Read(dbDataName, strconv.Itoa(dataset.ID), &dataset)
	//fmt.Println("New Dataset: ", dataset)
}

func DeleteDatasetId(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	HandleError(err)

	// Delete a fish from the database
	if err := IoTDB.Delete(dbDataName, strconv.Itoa(id)); err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
