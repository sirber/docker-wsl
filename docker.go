package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Parse arguments for Docker
	args := []string{"docker"}
	args = append(args, os.Args[1:]...)

	firstArg := os.Args[1]

	switch firstArg {
	case "wslstart":
		exec.Command("wsl", "sudo", "service", "docker", "start")
		break

	default:
		// Command
		cmd := exec.Command("wsl", args...)

		// Buffers
		var out bytes.Buffer
		var outError bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &outError

		// Execute
		cmd.Run()

		// Console oupout
		fmt.Print(outError.String())
		fmt.Print(out.String())
		break
	}
}
