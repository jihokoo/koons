package models

import (
  "golang.org/x/crypto/bcrypt"
  "labix.org/v2/mgo/bson"
  "log"
)

type User struct {
  Id bson.ObjectId `bson:"_id,omitempty"`
  Created string
  Updated string
  UserName string
  FirstName string
  LastName string
  Password string `db:"-"`
  PasswordHash []byte
}

func (u *User) HashPassword() {
  hashedPassword, hashError := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
  if hashError != nil {
      log.Fatal(hashError)
      panic(hashError) //this is a panic because bcrypt errors on invalid costs
  }

  u.PasswordHash = hashedPassword
}
