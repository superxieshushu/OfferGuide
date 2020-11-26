package main

import (
	"container/list"
	"fmt"
)

//使用堆栈
func PrintViaStack(l *list.List) {
	result := list.New()

	for v := l.Front(); v != nil; v = v.Next() {
		result.PushBack(v.Value)
	}

	for v := result.Front(); v != nil; v = v.Next() {
		fmt.Printf("%v", v.Value)
	}
}

//使用递归
func PrintViaRecursion(e *list.Element) {
	if e != nil {
		fmt.Printf("%v", e.Value)
		PrintViaRecursion(e.Next())
	}

}
