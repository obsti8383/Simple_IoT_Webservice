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

const dbDataName = "Devices"

func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	fmt.Fprintf(w, "<h1 style=\"font-family: Helvetica;\">Hello, welcome to IoT REST web service</h1>")
}

func GetDevices(w http.ResponseWriter, r *http.Request, _ mux.Params) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	records, err := IoTDB.ReadAll(dbDataName)
	HandleError(err)
	entries := []Device{}
	for _, e := range records {
		device := Device{}
		if err := json.Unmarshal([]byte(e), &device); err != nil {
			fmt.Println("Error", err)
		} else {
			entries = append(entries, device)
		}
	}

	if err := json.NewEncoder(w).Encode(entries); err != nil {
		panic(err)
	}
}

func GetDeviceId(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	HandleError(err)

	var device Device
	err = IoTDB.Read(dbDataName, strconv.Itoa(id), &device)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(device); err != nil {
			panic(err)
		}
	}
}

func PostDevice(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	// Read request body and close it
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	HandleError(err)
	defer r.Body.Close()

	// Convert JSON to Post struct
	var device Device
	if err := json.Unmarshal(body, &device); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// verify ID length
	if len(device.ID) > 50 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = IoTDB.Write(dbDataName, device.ID, device)
	HandleError(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(device); err != nil {
		panic(err)
	}

	//IoTDB.Read(dbDataName, strconv.Itoa(device.ID), &device)
	//fmt.Println("New Device: ", device)
}

func DeleteDeviceId(w http.ResponseWriter, r *http.Request, ps mux.Params) {
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
