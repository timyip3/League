package controller

import (
	"encoding/csv"
	"fmt"
	"league/main/matrix"
	"net/http"
	"strings"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	records, hasError := readFile(r, w)
	if hasError {
		return
	}

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	fmt.Fprint(w, response)
}

func InvertHandler(w http.ResponseWriter, r *http.Request) {
	records, hasError := readFile(r, w)
	if hasError {
		return
	}

	invertedMatrix, err := matrix.InvertMatrix(records)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	var response string
	for _, row := range invertedMatrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	fmt.Fprint(w, response)
}

func FlattenHandler(w http.ResponseWriter, r *http.Request) {
	records, hasError := readFile(r, w)
	if hasError {
		return
	}

	flattenedMatrix, err := matrix.FlattenMatrix(records)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	fmt.Fprint(w, flattenedMatrix)
}

func SumHandler(w http.ResponseWriter, r *http.Request) {
	records, hasError := readFile(r, w)
	if hasError {
		return
	}

	result, err := matrix.SumMatrix(records)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	fmt.Fprint(w, result, "\n")
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	records, hasError := readFile(r, w)
	if hasError {
		return
	}

	result, err := matrix.MultiplyMatrix(records)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}

	fmt.Fprint(w, result, "\n")
}

func readFile(r *http.Request, w http.ResponseWriter) ([][]string, bool) {
	file, _, err := r.FormFile("file")

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil, true
	}

	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil, true
	}
	return records, false
}
