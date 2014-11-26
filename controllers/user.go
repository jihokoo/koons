package controllers

import (
  "net/http"
  "fmt"
  "../models"
  "log"
  "time"
  "github.com/coopernurse/gorp"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Whats up breh?")
}

func CreateUser(dbmap *gorp.DbMap) func(w http.ResponseWriter, r *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    user := &models.User{
      Created: time.Now().Unix(),
      Updated: time.Now().Unix(),
      UserName: "jihokoo",
      FirstName: "Ji Ho",
      LastName: "Koo",
    }

    err := dbmap.Insert(user)
    if err != nil {
      log.Fatal(err)
    }

    fmt.Fprintf(w, "Whats up breh?")
  }
}
