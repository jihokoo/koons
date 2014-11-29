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
  r.Methods("GET").Path("/{userName}").HandlerFunc( controllers.GetUser(dbmap) )
  r.Methods("POST").Path("/{userName}").HandlerFunc( controllers.CreateUser(dbmap) )
  r.Methods("PUT").Path("/{userName}").HandlerFunc(controllers.UpdateUser(dbmap))
  r.Methods("DELETE").Path("/{userName}").HandlerFunc(controllers.DeleteUser(dbmap))
}
