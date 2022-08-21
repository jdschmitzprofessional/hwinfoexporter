package main

import (
	"encoding/json"
	"fmt"
	"hwstatexporter/collector"
	"hwstatexporter/data"
	"net/http"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getCurrentStats(exportData *data.ExportData) string {
	inst, err := json.Marshal(exportData)
	checkError(err)
	return string(inst)
}

func main() {
	hostname, err := os.Hostname()
	checkError(err)
	exprt := new(data.ExportData)
	exprt.Hostname = hostname
	go collector.Collector(exprt)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		responseString := getCurrentStats(exprt)
		fmt.Fprint(writer, responseString)
	})

	server := http.Server{
		Addr: ":8082",
	}
	fmt.Println("Starting server")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
