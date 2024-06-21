package main

import (
	"fmt"
	"os"
)

type Pwd struct {
	home string
	pwd  string
}

func CreatePwd() Pwd {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Cannot get pwd")
		os.Exit(1)
	}
	home := os.Getenv("HOME")
	return Pwd{
		home: home,
		pwd:  pwd,
	}
}
