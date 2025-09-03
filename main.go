package main

import (
	"encoding/json"
	"net/http"
	"os"
	"slices"
	"strings"
)

var Countries []map[string]any
var ListCountries []byte
var JSONFile, _ = os.ReadFile("./data.json")

func init() {
	// Unmarshall json file
	json.Unmarshal(JSONFile, &Countries)

	slices.SortFunc(Countries, func(a, b map[string]any) int {
		if a["Country Name"].(string) > b["Country Name"].(string) {
			return 1
		}
		return -1
	})

	// List all countries
	list := make([]string, 0, len(Countries))
	for _, country := range Countries {
		list = append(list, country["Country Name"].(string))
	}

	ListCountries, _ = json.Marshal(list)
}

func GetListCountries(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ListCountries)

}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	countryName := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]

	if countryName == "all" {
		w.WriteHeader(http.StatusOK)

		w.Write(JSONFile)
		return
	}

	index, found := slices.BinarySearchFunc[[]map[string]any, map[string]any, string](Countries, countryName, func(m map[string]any, s string) int {
		if m["Country Name"].(string) == s {
			return 0
		}
		if m["Country Name"].(string) > s {
			return 1
		}
		return -1
	})

	if !found {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"msg": "no data"}`))
		return
	}

	w.WriteHeader(http.StatusOK)

	jsonData, _ := json.Marshal(Countries[index])
	w.Write(jsonData)

}

func Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Api hecha con los datos extraidos de: https://datos.bancomundial.org/indicador/SP.DYN.TFRT.IN?end=2023&start=1960&view=chart`))
}

func main() {
	http.HandleFunc("/list", GetListCountries)
	http.HandleFunc("/get/", Get)
	http.HandleFunc("/", Root)

	http.ListenAndServe(":10000", nil)
}
