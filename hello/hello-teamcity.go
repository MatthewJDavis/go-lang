package main

import "fmt"

func main() {
	hello("Matt")
	hello("Teamcity build! :D !!!")

}

func hello(s string) {
	fmt.Println("Hello", s)
}
