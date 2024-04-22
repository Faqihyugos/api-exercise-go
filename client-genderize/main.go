package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Person struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

func main() {
	// Dapatkan response dari swapi
	resp, err := http.Get("https://api.genderize.io/?name=Michael%20Jordan")
	if err != nil {
		log.Fatal(err)
	}
	// Dapatkan response.Body dalam bentuk []byte
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// tampilkan response Body dalam bentuk []byte sebagai string
	fmt.Println(string(body))
	// Decode json ke dalam interface
	var person Person
	json.Unmarshal(body, &person)
	fmt.Println("Name :", person.Name)
	fmt.Println("Gender :", person.Gender)
}
