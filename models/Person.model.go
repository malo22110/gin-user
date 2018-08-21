package models

type Person struct {
	Name   		string `form:"name"`
	Email  		string `form:"email"`
	Secret 		string `form:"secret"`
	BirthDate 	string `form:"secret"`
}