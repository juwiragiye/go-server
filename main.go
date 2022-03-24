package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		 http.Error(w, "http method not supported", http.StatusMethodNotAllowed)
		 return
	}

	fmt.Fprintf(w, "Hello!")
}

type Person struct {
	Name string `json:"name"`
	Address string `json:"address"`
}


func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		 http.Error(w, "404 not found", http.StatusNotFound)
		 return

	}

	if r.Method != "POST" {
		 http.Error(w, "http method not supported", http.StatusMethodNotAllowed)
		 return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "the form has an error, %v", err)
		return
	}

	person := Person{r.Form.Get("name"), r.Form.Get("address")}
	res, err := json.Marshal(person)

	if err != nil {
		 http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w,  string(res))
	fmt.Println(r.Form)
	return



}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
		
	}
}