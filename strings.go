package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Contains("Daffa Fauzan", "Daffa"))
	fmt.Println(strings.Split("Daffa Fauzan", " "))
	fmt.Println(strings.ToLower("Daffa Fauzan"))
	fmt.Println(strings.ToUpper("Daffa Fauzan"))
	fmt.Println(strings.ReplaceAll("Daffa Fauzan Muhammad Fauzan", "Daffa", "Agus"))
}