package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	resp, err := http.Get("http://localhost:8080/users")

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Status code:", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)

	var response = []Users{}

	err = json.Unmarshal(body, &response)

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
