package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)
type Danchoi struct {
	Id int
	Name string
	Age int8
	Class []int
}

type DSDanchoi []Danchoi

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/cc", CcPage)
	http.HandleFunc("/api/music",DataApi)
	http.HandleFunc("/api/danchoi",DataDanchoi)
	http.HandleFunc("/api/danchois",DataListDanchoi)
	log.Fatal(http.ListenAndServe(":8080",nil))
}


func HomePage(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w,"This is home page")
}
func CcPage(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w,"This is cc page")
}
func DataApi(w http.ResponseWriter,r *http.Request) {
	var data = map[string]interface{} {
		"name": "Doraemon Remix sap xinh",
		"casi": "Hieu Hoa Hong",
	}
	json.NewEncoder(w).Encode(data)
}

func DataDanchoi(w http.ResponseWriter,r *http.Request) {
	var danchoi = Danchoi{1,"Hieu ngu", 22,[]int{1,2,3,4}}
	json.NewEncoder(w).Encode(danchoi)
}

func DataListDanchoi(w http.ResponseWriter,r *http.Request) {
	var danchois = []Danchoi{
		{1,"Hieu ngu", 22,[]int{1,2,3,4}},
		{2,"Hieu ko ngu", 22,[]int{1,2,3,4}},
		{3,"Hieu ngu vl", 22,[]int{1,2,3,4}},
	}
	json.NewEncoder(w).Encode(danchois)
}