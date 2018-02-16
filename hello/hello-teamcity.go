package main

import "fmt"

func main() {
	hello("Matt")
	hello("Teamcity build! :D")
	hello("webhook")

}

func hello(s string) {
	fmt.Println("Hello", s)
}
