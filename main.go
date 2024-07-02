package main

import (
	"fmt"
	"structures/lists"
)

func main() {
	var list lists.List[int]
	list = lists.LinkedList([]int{10, 20, 30, 40})
	size, err := list.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list, size)
}
