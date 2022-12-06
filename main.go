package main

import (
	"fmt"
	"golang-test/models"
	"golang-test/utils"
)

func main() {
	file, err := utils.OpenFile("person.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	jsonContent, err := utils.ReadJson(file)
	if err != nil {
		panic(err)
	}
	peopleContent, err := utils.JsonToArray(jsonContent)
	if err != nil {
		panic(err)
	}
	var (
		firstName string
		lastName  string
	)
	fmt.Print("First Name > ")
	fmt.Scanln(&firstName)
	fmt.Print("Last Name > ")
	fmt.Scanln(&lastName)
	newPerson := models.Person{
		FirstName: firstName,
		LastName:  lastName,
		Online:    false,
	}
	peopleContent = append(peopleContent, newPerson)
	jsonData, err := utils.ArrayToJson(peopleContent)
	if err != nil {
		panic(err)
	}
	err = utils.WriteJson(file, jsonData)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}
