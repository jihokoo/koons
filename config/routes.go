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
  r.Methods("QUERY").HandlerFunc(controllers.GetAllUsers(dbmap))
  r.Methods("GET").Path("/{userId}").HandlerFunc( controllers.GetUser(dbmap) )
  r.Methods("POST").HandlerFunc( controllers.CreateUser(dbmap) )
  r.Methods("PUT").Path("/{userId}").HandlerFunc(controllers.UpdateUser(dbmap))
  r.Methods("DELETE").Path("/{userId}").HandlerFunc(controllers.DeleteUser(dbmap))
}
