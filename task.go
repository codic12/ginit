package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func runTask(taskname string, taskarg string, task string) {

	logpath := "log/"
	ensureDir(logpath)
	output, err := os.OpenFile(logpath+task+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0750)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer output.Close()

	errors, err := os.OpenFile("log/"+task+".err", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0750)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer errors.Close()

	logOut := io.Writer(output)
	logErr := io.Writer(errors)

	logTime := time.Now()
	r := strings.NewReader(logTime.Format("\n##LOG: 2006-01-02 15:04:05 \n"))
	io.Copy(logOut, r)
	io.Copy(logErr, r)

	args := strings.Fields(taskarg)
	cmd := exec.Command(taskname, args...)
	cmd.Stdout = logOut
	cmd.Stderr = logErr
	fmt.Println("Running", taskname, "With args:", taskarg)
	cmd.Start()
}

func ensureDir(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, 0750)
	}
	if err != nil {
		return err
	}
	return nil
}
