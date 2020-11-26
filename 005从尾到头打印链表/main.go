package main

import "container/list"

func main() {
	var data = list.New()
	data.PushFront("1")
	data.PushFront("2")
	data.PushFront("3")
	data.PushFront("4")
	data.PushFront("5")
	data.PushFront("6")
	//PrintViaStack(data)
	PrintViaRecursion(data.Front())
}
