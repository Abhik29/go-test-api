package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	fs := http.FileServer(http.Dir("./static"))
  	http.Handle("/", fs)
    http.HandleFunc("/api/home", homePage)
    log.Fatal(http.ListenAndServe(":8001", nil))
}

func main() {
    handleRequests()
}