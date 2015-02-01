package controllers

import (
  "net/http"
  "encoding/json"
  "../models"
  "log"
  "time"
  "io/ioutil"
  "github.com/gorilla/mux"
  "github.com/coopernurse/gorp"
  "gopkg.in/mgo.v2/bson"
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
      log.Print(dbError)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
  }
}

// TODO: add go routines, channels, locks

func GetAllUsers(dbmap *gorp.DbMap) func(w http.ResponseWriter, r *http.Request) {

  return usersHandler(nil, func(r *http.Request) *[]models.User {
    var users []models.User
    _, dbError := dbmap.Select(&users, "select * from \"user\"")
    if dbError != nil {
      log.Print(dbError)
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
      log.Print(dbError)
    }

    return &user
  }, nil)
}

func CreateUser(dbmap *gorp.DbMap) func(w http.ResponseWriter, r *http.Request) {

  return usersHandler(func(r *http.Request) *models.User {

    body, readError := ioutil.ReadAll(r.Body)
    if readError != nil {
      log.Print(readError)
    }

    // TODO: validations (validate using user struct)
    var user *models.User
    jsError := json.Unmarshal(body, &user)
    if jsError != nil {
      log.Print(jsError)
    }

    var currentTime string = time.Now().Format("2006-01-02T15:04:05.999999999Z0700")
    user.Created = currentTime
    user.Updated = currentTime

    user.HashPassword()
    user.Id = bson.NewObjectId()

    dbError := dbmap.Insert(user)
    if dbError != nil {
      log.Print(dbError)
    }

    return user
  }, nil)
}

func UpdateUser(dbmap *gorp.DbMap) func(w http.ResponseWriter, r *http.Request) {

  return usersHandler(func(r *http.Request) *models.User {

    userId := mux.Vars(r)["userId"]
    body, readError := ioutil.ReadAll(r.Body)
    if readError != nil {
        log.Print(readError)
    }

    // TODO: validations
    var user *models.User
    jsError := json.Unmarshal(body, &user)
    if jsError != nil {
        log.Print(jsError)
    }

    user.Updated = time.Now().Format("2006-01-02T15:04:05.999999999Z0700")

    if isObjectId := bson.IsObjectIdHex(userId); isObjectId {
      user.Id = bson.ObjectIdHex(userId)
    } else {
      log.Print("User Id is incorrect.")
    }

    // count, dbError
    _, dbError := dbmap.Update(user)
    if dbError != nil {
      log.Print(dbError)
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
      log.Print(dbError)
    }

    var user *models.User = nil

    return user
  }, nil)
}
