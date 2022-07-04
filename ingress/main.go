package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("coffee"))
	})
	http.ListenAndServe(":9090", nil)
}
