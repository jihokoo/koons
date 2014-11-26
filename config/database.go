package config

import (
  "log"
  "database/sql"
  _ "github.com/lib/pq"
  "github.com/coopernurse/gorp"
  "../models"
)

func DataBaseStart() (*gorp.DbMap) {
  db, err := sql.Open("postgres", "dbname=koons user=jihokoo sslmode=disable port=5432")
  if err != nil {
    log.Fatal(err)
  }

  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

  dbmap.AddTable(models.User{}).SetKeys(true, "Id")
  return dbmap
}
