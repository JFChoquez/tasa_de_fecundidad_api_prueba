package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var Countrys []map[string]any

func init() {
	f, _ := os.ReadFile("./data.json")

	json.Unmarshal(f, &Countrys)
}

func ListCountrys(w http.ResponseWriter, r *http.Request) {
	list := make([]string, 0, len(Countrys))

	for _, country := range Countrys {
		switch v := country["Country Name"].(type) {
		case string:
			if v != "" {
				list = append(list, v)
			}
		}

	}

	w.Header().Set("Content-Type", "application/json")

	jsonData, _ := json.Marshal(list)
	w.Write(jsonData)

}

func Get(w http.ResponseWriter, r *http.Request) {
	countryName := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]

	w.Header().Set("Content-Type", "application/json")

	for _, el := range Countrys {

		if el["Country Name"] == countryName {

			w.WriteHeader(http.StatusOK)

			jsonData, _ := json.Marshal(el)
			w.Write(jsonData)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{msg: "no data"}`))

}

func main() {
	http.HandleFunc("/list", ListCountrys)
	http.HandleFunc("/get/", Get)

	fmt.Println("Runing on port: ")
	http.ListenAndServe(":8080", nil)
}
