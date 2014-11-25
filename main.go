package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
  "log"
  "database/sql"
  _ "github.com/lib/pq"
  "github.com/coopernurse/gorp"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Whats up breh?")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Whats up breh?")
}

type User struct {
  Id int64
  Created int64
  Updated int64
  UserName string
  FirstName string
  LastName string
}

func main() {
  r := mux.NewRouter().StrictSlash(true)
  // Subrouter will the url address where our API lives
  // s := r.Host("www.domain.com").Subrouter()

  // mapping url routes to handlers

  // '/' home route
  r.HandleFunc("/", HomeHandler).
    Methods("GET")

  // postgresql server running at 5432

  // '/:userName' route should handle all requests related to the user
  user := r.PathPrefix("/user").Subrouter()
  user.Methods("QUERY").HandlerFunc(UserHandler)

  user.Methods("GET").Path("/{userName}").HandlerFunc(UserHandler)
  user.Methods("POST").HandlerFunc(UserHandler)
  user.Methods("PUT").HandlerFunc(UserHandler)
  user.Methods("DELETE").HandlerFunc(UserHandler)

  http.Handle("/", r)

  fmt.Println("Starting server on :3000")
  http.ListenAndServe(":3000", nil)

  db, err := sql.Open("postgres", "dbname=koons user=jihokoo port=5432")
  if err != nil {
    log.Fatal(err)
  }
  // construct a gorp DbMap
  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
  fmt.Println(dbmap)
}


