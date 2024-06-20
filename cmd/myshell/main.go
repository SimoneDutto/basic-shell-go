package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		slog.Info("Error during reading command", err)
	}
	command, _ = strings.CutSuffix(command, "\n")
	switch command {
	case "ciao":
		fmt.Printf("bau")
	default:
		fmt.Printf("%s: command not found", command)
	}
}
