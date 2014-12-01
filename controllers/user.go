package controllers

import (
  "net/http"
  "encoding/json"
  "fmt"
  "../models"
  "log"
  "time"
  "io/ioutil"
  "github.com/gorilla/mux"
  "github.com/coopernurse/gorp"
  "strconv"
)

func usersHandler(cb1 func(r *http.Request) *models.User, cb2 func(r *http.Request) *[]models.User) func(w http.ResponseWriter, r*http.Request) {

  return func(w http.ResponseWriter, r *http.Request) {
    var js []byte
    var dbError error

    if(cb1 != nil){
      js, dbError = json.Marshal( cb1(r) )
    } else{
      js, dbError = json.Marshal( cb2(r) )
    }

    if dbError != nil {
      http.Error(w, dbError.Error(), http.StatusInternalServerError)
      log.Fatal(dbError)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
  }
}

// TODO: add go routines, channels, locks

func GenericHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Whats up breh?")
}

func GetAllUsers(dbmap *gorp.DbMap) func(w http.ResponseWriter, r *http.Request) {

  return usersHandler(nil, func(r *http.Request) *[]models.User {
    var users []models.User
    _, dbError := dbmap.Select(&users, "select * from \"user\"")
    if dbError != nil {
      log.Fatal(dbError)
    }

    return &users
  })

}

func GetUser(dbmap *gorp.DbMap) func(w http.ResponseWriter, r *http.Request) {

  return usersHandler(func(r *http.Request) *models.User  {

    userId := mux.Vars(r)["userId"]

    var user models.User
    dbError := dbmap.SelectOne(&user, "select * from \"user\" where id=$1", userId)
    if dbError != nil {
      log.Fatal(dbError)
    }

    return &user
  }, nil)
}

func CreateUser(dbmap *gorp.DbMap) func(w http.ResponseWriter, r *http.Request) {

  return usersHandler(func(r *http.Request) *models.User {

    body, readError := ioutil.ReadAll(r.Body)
    if readError != nil {
      log.Fatal(readError)
    }

    // TODO: validations
    var user *models.User
    jsError := json.Unmarshal(body, &user)
    if jsError != nil {
      log.Fatal(jsError)
    }

    user.Created = time.Now().Unix()
    user.Updated = time.Now().Unix()

    dbError := dbmap.Insert(user)
    if dbError != nil {
      log.Fatal(dbError)
    }

    return user
  }, nil)
}

func UpdateUser(dbmap *gorp.DbMap) func(w http.ResponseWriter, r *http.Request) {

  return usersHandler(func(r *http.Request) *models.User {

    userId := mux.Vars(r)["userId"]
    body, readError := ioutil.ReadAll(r.Body)
    if readError != nil {
        log.Fatal(readError)
    }

    // TODO: validations
    var user *models.User
    jsError := json.Unmarshal(body, &user)
    if jsError != nil {
        log.Fatal(jsError)
    }

    user.Updated = time.Now().Unix()
    var strError error
    user.Id, strError = strconv.ParseInt(userId, 10, 64)
    if strError != nil {
      log.Fatal(strError)
    }

    // count, dbError
    _, dbError := dbmap.Update(user)
    if dbError != nil {
      log.Fatal(dbError)
    }

    return user
  }, nil)
}

func DeleteUser(dbmap *gorp.DbMap) func(w http.ResponseWriter, r *http.Request) {

  return usersHandler(func(r *http.Request) *models.User {

    userId := mux.Vars(r)["userId"]

    // count, dbError
    _, dbError := dbmap.Exec("delete from \"user\" where id=$1", userId)
    if dbError != nil {
      log.Fatal(dbError)
    }

    var user *models.User = nil

    return user
  }, nil)
}
