package main

import "fmt"

func main() {
	hello("Matt")
	hello("Teamcity build! :D")
	hello("webhook2")

}

func hello(s string) {
	fmt.Println("Hello", s)
}
