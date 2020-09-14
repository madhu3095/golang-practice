package main

import (
	"fmt"
	"math/rand"
	"time"
)

type linkedlist struct {
	value int
	pointer *linkedlist
}

func main() {

	rand.Seed(time.Now().UnixNano())
	list := createList(10, nil)

	//for loop for traversal
	for ; list !=nil ; list = list.pointer {
		fmt.Printf("%v\n",list.value)
	}

}

func createList(num int ,temp *linkedlist) *linkedlist {
	if temp == nil {
		temp =&linkedlist {rand.Intn(100) , nil}
		num--
	}
	tempList := temp

	for i:=0 ; i< num ; i++ {
		t :=&linkedlist {rand.Intn(100), nil}
		tempList.pointer = t
		tempList =tempList.pointer
	}
	return temp
}