package model

import ("github.com/niketa/gorepo/webapp/database"
		"fmt"
		"log"
		)

var CurrentUserEmail = " "

func CreateUsers(email string,password string){

	db := database.GetConnection()

	insert_users,err := db.Prepare("INSERT INTO users(email,password) values (?,password(?))")


       if err!=nil {
                panic(err.Error())
              }

                insert_users.Exec(email,password)

                defer db.Close()
                }

func CheckUsers(email string,password string) bool {

	db := database.GetConnection()
	
	
	
	results, err := db.Query("SELECT * FROM users WHERE email='"+email+"' AND `password`=password('"+password+"')")
	if err != nil {
		panic(err.Error())
	}

	if(!results.Next()){
		//no records
		fmt.Println("user/password invalid")
		return false
	} 

	return true
}



func DisplayEmail(email string) string{

	var emailid,psw string


	db := database.GetConnection()

	results,err := db.Query("SELECT email,password FROM users WHERE email=?",email)
	if err != nil {
		panic(err.Error()) 
	}

	defer results.Close()

		for results.Next() {
			err := results.Scan(&emailid,&psw)
				if err != nil {
				log.Fatal(err)
			}
	
}	
return  emailid 	
}
