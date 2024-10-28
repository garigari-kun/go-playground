package main

import "fmt"

type LinkedList struct {
	Value any
	Next  *LinkedList
}

func (ll *LinkedList) Insert(pos int, val any) *LinkedList {
	if ll == nil || pos == 0 {
		return &LinkedList{
			Value: val,
			Next:  ll,
		}
	}

	ll.Next = ll.Next.Insert(pos-1, val)
	return ll
}

func main() {
	var i any
	i = 20
	fmt.Println(i)
	i = "Hello"
	fmt.Println(i)
	i = struct {
		FirstName string
	}{"Hoge"}
	fmt.Println(i)
}
