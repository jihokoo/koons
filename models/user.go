package models

type User struct {
  // TODO: make the id a universally unique id
  Id int64
  Created int64
  Updated int64
  UserName string
  FirstName string
  LastName string
  // TODO: hash password
}
