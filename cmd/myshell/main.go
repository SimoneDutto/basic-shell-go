package main

import (
	"bufio"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"
)

var Commands map[string]Command
var pwd Pwd

func init() {
	pwd = CreatePwd()
	Commands = make(map[string]Command, 10)
	Commands["echo"] = Command{
		Type:    INTERNAL,
		Command: "echo",
		F: func(args []string) error {
			fmt.Println(strings.Join(args, " "))
			return nil
		},
	}
	Commands["exit"] = Command{
		Type:    INTERNAL,
		Command: "exit",
		F: func(args []string) error {
			os.Exit(0)
			return nil
		},
	}
	Commands["cd"] = Command{
		Type:    INTERNAL,
		Command: "cd",
		F: func(args []string) error {
			if len(args) != 1 {
				return &WrongArgumentsError{msg: "wrong arguments"}
			}
			cdPath := args[0]
			if strings.HasPrefix(cdPath, "~") {
				cdPath = strings.Replace(cdPath, "~", pwd.home, -1)
			}
			if !strings.HasPrefix(cdPath, "/") {
				cdPath = path.Join(pwd.pwd, cdPath)
			}
			s, err := os.Stat(cdPath)
			if err != nil || !s.IsDir() {
				fmt.Printf("cd: %s: No such file or directory\n", cdPath)
			}
			pwd.pwd = cdPath
			return nil
		},
	}
	Commands["pwd"] = Command{
		Type:    INTERNAL,
		Command: "pwd",
		F: func(args []string) error {
			fmt.Println(pwd.pwd)
			return nil
		},
	}
	Commands["type"] = Command{
		Type:    INTERNAL,
		Command: "type",
		Man: `type: missing arguments
usage: type <command>`,
		F: func(args []string) error {
			if len(args) != 1 {
				return &WrongArgumentsError{msg: "wrong arguments"}
			}
			c, ok := getCommand(args[0])
			if !ok {
				fmt.Printf("%s: not found\n", args[0])
				return nil
			}
			if c.Type == SYSTEM {
				fmt.Printf("%s is %s\n", c.Command, c.Path)
			} else if c.Type == INTERNAL {
				fmt.Printf("%s is a shell builtin\n", c.Command)
			}
			return nil
		},
	}
}

func getCommand(command string) (Command, bool) {
	c, ok := Commands[command]
	if ok {
		return c, ok
	}
	paths := os.Getenv("PATH")
	if paths == "" {
		return Command{}, false
	}
	for _, p := range strings.Split(paths, ":") {
		commandPath := path.Join(p, command)
		_, err := os.Stat(commandPath)
		if err != nil {
			continue
		}
		return CreateSystemCommand(command, commandPath), true
	}
	return Command{}, false
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		inputReader, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			slog.Error("Error during reading command", err)
		}
		input, _ := strings.CutSuffix(inputReader, "\n")
		fields := strings.Fields(input)
		var command string
		var args []string
		if len(fields) > 0 {
			command = fields[0]
		}
		if len(fields) > 1 {
			args = fields[1:]
		}
		c, ok := getCommand(command)
		if !ok {
			fmt.Printf("%s: command not found\n", command)
			continue
		}
		err = c.F(args)
		if err != nil {
			var wrongargError *WrongArgumentsError
			if errors.As(err, &wrongargError) {
				fmt.Println(c.Man)
			} else {
				fmt.Println(err)
			}
		}
	}

}
