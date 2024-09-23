package main

import "log"

func main() {
	 var hello string = echolot("Hello, world!")
	 log.Println(hello)
}

func echolot(s string) string {
	return s;
}