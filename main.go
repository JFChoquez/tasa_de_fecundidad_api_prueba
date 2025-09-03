package main

import (
	"encoding/json"
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

	w.WriteHeader(http.StatusOK)
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

func Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Api hecha con los datos extraidos de: https://datos.bancomundial.org/indicador/SP.DYN.TFRT.IN?end=2023&start=1960&view=chart`))
}

func main() {
	http.HandleFunc("/list", ListCountrys)
	http.HandleFunc("/get/", Get)
	http.HandleFunc("/", Root)

	http.ListenAndServe(":10000", nil)
}
