package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Model string `json: "model"`
	Color string `json: "color"`
}

func main() {
	car := Car{
		Model: "subaru",
		Color: "red",
	}
	jsonData , err := json.Marshal(car)
	if err != nil {
		fmt.Println("error encoding Json:",err)
		return
	}
	fmt.Println(string(jsonData))
}
