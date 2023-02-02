package main

import (
	application "exercise4/application"
)

func main() {
	studentdb := application.StudentDB{}
	studentdb.Initialize("oneplus", "oneplus@123K", "student") //user,password,database
	studentdb.Run(":8080")                                     //localhost port
}
