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

type TaskFile struct {
	Name map[string]TaskFile
	Task struct {
		Exec string            `json:"exec"`
		Args string            `json:"args"`
		Env  map[string]string `json:"envs"`
	} `json:"taskname"`
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

	fmt.Println("EXEC:", task.Task.Exec)

}

//func runTask(taskName string) {
//	cmd := exec.Command(taskName)
//
//	cmd.Stdout = os.Stdout
//	err := cmd.Start()
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Printf("Just ran subprocess %d, exiting\n", cmd.Process.Pid)
//}
