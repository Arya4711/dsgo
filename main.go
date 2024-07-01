package main

import (
	"fmt"
	"structures/lists"
)

func main() {
	var list lists.List[int]
	list = lists.LinkedList([]int{10, 20, 30})
	err := list.Add(0, 40)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list)
}
