package server

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
)


func OpenDB(dbName string) *sql.DB {


  config := DBConfig{"127.0.0.1","3306","root",
    "ig3Axi8faV4ooth3Chu5Loh3Ohmohlah"}

  conStr := fmt.Sprintf("%[1]s:%[2]s@tcp(%[3]s:%[4]s)/%[5]s", config.Username,
    config.Password, config.DBAddr, config.DBPort, dbName)

  fmt.Println(conStr)

  db, err := sql.Open("mysql", conStr) 

  if err != nil {
    fmt.Println(err)
  }

  return db

}


type DBConfig struct {
  DBAddr string
  DBPort string
  Username string
  Password string
}
