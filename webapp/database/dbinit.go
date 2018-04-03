package database

import (_ "github.com/go-sql-driver/mysql"
        "database/sql"
        "log")

var db *sql.DB
var err error

func InitialiseDB(){

//root user and password
  db, err := sql.Open("mysql","root:root@/")

  if err!= nil {
      log.Print(err.Error())
  }
   defer db.Close()

   _,err = db.Exec("CREATE DATABASE if not exists dbusers")
   if err != nil {
       panic(err)
   }
                                              //app user @ user ip
   _,err = db.Exec("GRANT ALL ON dbusers.* to 'niketa'@'127.0.0.1' identified by 'jain'")
   if err != nil {
       panic(err)
   }

   _,err = db.Exec("USE dbusers")
   if err != nil {
       panic(err)
   }

   _,err = db.Exec("CREATE TABLE if not exists users ( id int(5) not null auto_increment,email varchar(50) not null,password varchar(255) not null,primary key(id),unique(email) )")
   if err != nil {
       panic(err)
   }

}