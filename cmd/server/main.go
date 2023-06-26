package main

import "fmt"

func Run() error {
	fmt.Println("stating up our application!..")
	return nil
}

func main() {
	fmt.Println("Go Rest API Learning!..")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
