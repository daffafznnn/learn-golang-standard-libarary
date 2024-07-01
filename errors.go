package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError =  errors.New("validation error")
	NotFoundError   =  errors.New("not found error")
)

func getById(id string) error {
	if id == "1" {
		return ValidationError
	}

	if id != "daffa" {
		return NotFoundError
	}

	// sukses

	return nil
}

func main(){
	err := getById("daffa")
	if err != nil { 
		if errors.Is(err, ValidationError){
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
				fmt.Println("unknown error")
		}
		
	}
}