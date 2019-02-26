package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Firstname string `json:"firsname"`
	Lastname  string `json:"lastname"`
}

func main() {
	bytes, err := json.Marshal(Person{
		Firstname: "Jhon",
		Lastname:  "Dow",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
