package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
  "./config"
  "os"
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

  port := os.Getenv("PORT")
  fmt.Println("Starting server on :" + port)
  http.ListenAndServe(":" + port, nil)
}
