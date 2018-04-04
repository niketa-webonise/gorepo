package model

import ("github.com/niketa/gorepo/webapp/database"
		"fmt"
		"log"
		)

var CurrentUserEmail = ""

type User struct{
	Email string
	FirstName string
	LastName string
}

func GetUserByEmail(email string) User {
	db := database.GetConnection()
	
	results, err := db.Query("SELECT firstname, lastname FROM users WHERE email='"+email+"'")
	if err != nil {
		panic(err.Error())
	}

	userdata := User{}
    userdata.Email = email

    var firstname string
    var lastname string

	for results.Next() {
		err := results.Scan(&firstname, &lastname)
		if err != nil {
			log.Fatal(err)
		}
		userdata.FirstName = firstname
		userdata.LastName = lastname
	}

	return userdata
}

func CreateUsers(email string,password string,firstname string,lastname string) {

	db := database.GetConnection()

	insert_users,err := db.Prepare("INSERT INTO users(email,password,firstname,lastname) values (?,password(?),?,?)")

    if err!=nil {
        panic(err.Error())
    }

    insert_users.Exec(email,password,firstname,lastname)

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
