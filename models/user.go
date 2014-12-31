package models

import (
  "golang.org/x/crypto/bcrypt"
  "log"
)

type User struct {
  // TODO: make the id a universally unique id
  Id int64
  Created int64
  Updated int64
  UserName string
  FirstName string
  LastName string
  Password string `db:"-"`
  PasswordHash []byte
  // TODO: hash password
}

func (u *User) HashPassword() {
  hashedPassword, hashError := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
  if hashError != nil {
      log.Fatal(hashError)
      panic(hashError) //this is a panic because bcrypt errors on invalid costs
  }

  u.PasswordHash = hashedPassword
}
