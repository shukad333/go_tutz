package main

import (
	"errors"
	"fmt"
	"log"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty Name")
	}

	message := fmt.Sprintf("Hi %v. Welcome!", name)
	return message, nil

}

func main() {

	msg, err := Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg, err)

}
