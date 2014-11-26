package config

import (
  "github.com/gorilla/mux"
  "github.com/coopernurse/gorp"
  "../controllers"
)

func RegisterIndexRoutes(r *mux.Router) {
  // '/' home route
  r.HandleFunc("/", controllers.HomeHandler).
    Methods("GET")
}

func RegisterUserRoutes(r *mux.Router, dbmap *gorp.DbMap) {
  r.Methods("QUERY").HandlerFunc(controllers.UserHandler)
  r.Methods("GET").Path("/{userName}").HandlerFunc(controllers.UserHandler)
  r.Methods("POST").Path("/{userName}").HandlerFunc( controllers.CreateUser(dbmap) )
  r.Methods("PUT").Path("/{userName}").HandlerFunc(controllers.UserHandler)
  r.Methods("DELETE").Path("/{userName}").HandlerFunc(controllers.UserHandler)
}
