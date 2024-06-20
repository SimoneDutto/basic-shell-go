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
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			slog.Info("Error during reading command", err)
		}
		input, _ := strings.CutSuffix(command, "\n")
		commands := strings.Fields(input)
		command, args := commands[0], commands[1:]
		_ = args
		switch command {
		case "ciao":
			fmt.Printf("bau\n")
		case "exit":
			os.Exit(0)
		default:
			fmt.Printf("%s: command not found\n", command)
		}
	}

}
