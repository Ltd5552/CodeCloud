package main

import (
	"fmt"
	"os/exec"
)

func RunCmd() {
	RunCmd := "docker run --name code -v /root/codecloud/workplace:/usr/src/codecloud -w /usr/src/codecloud python python *.py"
	err := exec.Command("/bin/sh", "-c", RunCmd).Start()
	if err != nil {
		return
	} else {
		fmt.Println("docker运行成功！")
	}
}
func ResultCmd() {
	DelCmd := "rm code.txt"
	cmd := exec.Command("/bin/sh", "-c", DelCmd)
	if err := cmd.Run(); err != nil {
		fmt.Println("RUN: ", err.Error())
		return
	} else {
		fmt.Println("删除成功！")
	}

	ResultCmd := "docker container logs code >>/root/codecloud/workplace/code.txt"
	for {
		cmd = exec.Command("/bin/sh", "-c", ResultCmd)
		err := cmd.Run()
		if err == nil {
			fmt.Println("保存成功！")
			break
		}
	}
}

func DeCmd() {
	DelCmd := "docker rm code"
	cmd := exec.Command("/bin/sh", "-c", DelCmd)
	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		return
	} else {
		fmt.Println("删除成功！")
	}
}
func main() {
	RunCmd()
	ResultCmd()
	ResultCmd()
	DeCmd()
}
