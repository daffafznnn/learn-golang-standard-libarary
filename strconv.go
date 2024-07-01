package main

import (
	"fmt"
	"strconv"
)

func main(){
	// string to boolean
	resultBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	 } else {
			 fmt.Println("Result:", resultBool)
	 }

	//  string to integer
		resultInt, err := strconv.Atoi("500")
	if err != nil {
		fmt.Println("Error", err.Error())
	 } else {
			 fmt.Println("Result:", resultInt)
	 }
	 
}