package main

import (
	"container/list"
	"fmt"
)

func main() {
		var data *list.List = list.New()

		data.PushBack("Muhammad")
		data.PushBack("Daffa")
		data.PushBack("Fauzan")

		var head *list.Element = data.Front()
    fmt.Println(head.Value) // Muhammad

		next := head.Next() // Daffa
		fmt.Println(next.Value)

		next = next.Next() // Fauzan
		fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}