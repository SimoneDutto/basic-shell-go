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
	return Pwd{
		home: pwd,
		pwd:  pwd,
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
