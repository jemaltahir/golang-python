package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Person struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type r []Result

type Result struct {
	ID      string `json:"id" `
	Result1 string `json:"result1"`
	Result2 string `json:"result2"`
	Result3 string `json:"result3"`
	Result4 string `json:"result4"`
	Result5 string `json:"result5"`
}

func main() {
	person := Person{
		ID:        "123",
		FirstName: "John",
		LastName:  "Doe",
	}

	personJSON, _ := json.Marshal(person)
	cmd := exec.Command("python3", "devel.py", string(personJSON))
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var results r
	json.Unmarshal(output, &results)

	fmt.Println(results)
}
