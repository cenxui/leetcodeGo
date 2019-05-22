package main

import "fmt"

func main() {
	b := make(chan interface{})

	fmt.Println("Hello main thread start")

	fmt.Println("Hello main thread running")

	go func() {
		fmt.Println("Hello new thread start")

		fmt.Println("Hello new thread running")

		fmt.Println("Hello new thread stop")
		<-b

	}()
	fmt.Println("Hello main thread stop")
	<-b
}

