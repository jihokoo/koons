package config

import (
  "log"
  "database/sql"
  _ "github.com/lib/pq"
  "github.com/coopernurse/gorp"
  "../models"
)

func DataBaseStart() (*gorp.DbMap) {
  db, dbError := sql.Open("postgres", "dbname=koons user=jihokoo sslmode=disable port=5432")
  if dbError != nil {
    log.Fatal(dbError)
  }

  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

  usersTable := dbmap.AddTable(models.User{}).SetKeys(true, "Id")
  usersTable.ColMap("Id").SetUnique(true)
  usersTable.ColMap("UserName").SetUnique(true)

  dbError = dbmap.CreateTablesIfNotExists()
  if dbError != nil {
    log.Fatal(dbError)
  }
  return dbmap
}
