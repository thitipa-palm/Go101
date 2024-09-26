package main

import "fmt"

func main() {

	go doSomeThing1()
	doSomeThing2()
}

func doSomeThing1() {
	fmt.Println("do 1")
}

func doSomeThing2() {
	fmt.Println("do 2")
}
