package main

import "fmt"
import "log"
import "os"

func main() {
	val, err := os.LookupEnv("JUMP_TO")

	if !err {
		fmt.Println("Please set JUMP_TO")
		log.Fatal(err)
	}
	fmt.Println("Jumping to", val)
}
