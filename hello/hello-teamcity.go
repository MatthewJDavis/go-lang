package main

import "fmt"

func main() {
	hello("Matt")
	hello("Teamcity build! :D")
	hello("webhook")
	hello("test me")

}

func hello(s string) {
	fmt.Println("Hello", s)
}
