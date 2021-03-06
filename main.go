package main

import "fmt"
import "os"
import "os/exec"

func main() {
	val, err := os.LookupEnv("JUMP_TO")

	if !err {
		fmt.Println("Please set JUMP_TO")
		return
	}
	cmd := exec.Command("cd")
	cmd.Dir = val
	cmd.Run()
	fmt.Println("Jumped to", val)
}
