package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Parse arguments for Docker
	args := []string{"docker"}
	args = append(args, os.Args[1:]...)

	// Execute
	cmd := exec.Command("wsl", args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	// Results
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(out.String())
}
