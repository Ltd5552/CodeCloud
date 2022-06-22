package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	RunCmd := "docker run --name code -v /root/codecloud/workplace:/usr/src/codecloud -w /usr/src/codecloud python python *.py"
	cmd := exec.Command("/bin/sh", "-c", RunCmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: ", err.Error())
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		return
	}
	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println("ReadAll stderr: ", err.Error())
		return
	}
	if len(bytesErr) != 0 {
		fmt.Printf("stderr is not nil: %s", bytesErr)
		return
	}
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: ", err.Error())
		return
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait: ", err.Error())
		return
	}
	fmt.Printf("stdout: %s", bytes)
}
