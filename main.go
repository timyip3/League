package main

import (
	"league/main/controller"
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		/echo:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
//		/invert:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/invert"
//		/flatten:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/flatten"
//		/sum:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/sum"
//		/multiply:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/multiply"

func main() {
	http.HandleFunc("/echo", controller.EchoHandler)
	http.HandleFunc("/invert", controller.InvertHandler)
	http.HandleFunc("/flatten", controller.FlattenHandler)
	http.HandleFunc("/sum", controller.SumHandler)
	http.HandleFunc("/multiply", controller.MultiplyHandler)

	http.ListenAndServe(":8080", nil)
}
