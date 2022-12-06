package utils

import (
	"encoding/json"
	"golang-test/models"
	"io/ioutil"
	"os"
)

func ReadJson(file *os.File) ([]byte, error) {
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte{}, err
	}
	return content, nil
}

func WriteJson(file *os.File, content []byte) error {
	err := os.WriteFile(file.Name(), content, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ArrayToJson(people []models.Person) ([]byte, error) {
	jsonData, err := json.Marshal(people)
	if err != nil {
		return []byte{}, err
	}
	return jsonData, nil
}

func JsonToArray(jsonData []byte) ([]models.Person, error) {
	var people []models.Person
	err := json.Unmarshal(jsonData, &people)
	if err != nil {
		return []models.Person{}, err
	}
	return people, nil
}

func OpenFile(filename string) (*os.File, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}
