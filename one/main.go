package main

import (
	"log"
	"net/http"
	"strings"
)

func do(writer http.ResponseWriter, request *http.Request) {
	header := request.Header
	for key, value := range header {
		writer.Header().Set(key, strings.Join(value, ""))
	}

	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write([]byte("200"))
	if err != nil {
		log.Println("write response err ", err)
	}
}

func HandleFunc(writer http.ResponseWriter, request *http.Request) {
	do(writer, request)
}

type handler struct{}

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	do(writer, request)
}

func main() {
	// 1
	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		do(writer, request)
	})

	if err := http.ListenAndServe(":1024", nil); err != nil {
		log.Fatal(err)
	}

	// 2
	if err := http.ListenAndServe(":1024", http.HandlerFunc(HandleFunc)); err != nil {
		log.Fatal(err)
	}

	// 3
	if err := http.ListenAndServe(":1024", &handler{}); err != nil {
		log.Fatal(err)
	}

	// 4
	server := http.NewServeMux()
	server.HandleFunc("/healthz", do)
	if err := http.ListenAndServe(":1024", server); err != nil {
		log.Fatal(err)
	}
}
