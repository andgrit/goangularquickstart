package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "log"
  "fmt"
)

const PORT_STRING = ":8080"
func main() {
  r := mux.NewRouter()
  r.PathPrefix("/").Handler(http.FileServer(http.Dir("../angular"))) // assume running the go program in goserver directory
  fmt.Printf("Try http://localhost%s\n", PORT_STRING)
  log.Fatal(http.ListenAndServe(PORT_STRING, r))
}
