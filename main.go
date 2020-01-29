package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	privilaged bool
	task       TaskFile
)

type TaskFile map[string]Task

type Task struct {
	Exec string
	Args string
	Envs map[string]string
}

func main() {
	if os.Getuid() == 0 {
		fmt.Println("This program is not ready to be ran as root")
		os.Exit(1)
	}
	str, err := ioutil.ReadFile("init.json")
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal([]byte(str), &task)

	for k := range task {
		fmt.Println(k+":", task[k].Exec, task[k].Args)
		runTask(task[k].Exec, task[k].Args, k)
	}
}
