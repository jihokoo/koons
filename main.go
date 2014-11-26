package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
  "./config"
)

func main() {
  r := mux.NewRouter().StrictSlash(true)
  // Subrouter will the url address where our API lives
  // s := r.Host("www.domain.com").Subrouter()

  dbmap := config.DataBaseStart() // create database map

  // Register routes and handlers to http
  config.RegisterIndexRoutes(r)

  user := r.PathPrefix("/user").Subrouter() // create subroute for user endpoints
  config.RegisterUserRoutes(user, dbmap)

  http.Handle("/", r)
  fmt.Println("Starting server on :9000")
  http.ListenAndServe(":9000", nil)
}
