package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func getRandom() (ret int, err error) {
	prg := "node"
	arg1 := "index.js"
	arg2 := "33"
	arg3 := "44"
	cmd := exec.Command(prg, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		return 0, fmt.Errorf("unable to run cmd")
	}
	formattedOutput := strings.Split(string(stdout), "\n")
	res, err2 := strconv.Atoi(formattedOutput[0])
	if err2 != nil {
		return 0, fmt.Errorf("unabel to convert output to int", err2)
	}
	return res, nil
}

func main() {
	ret, err := getRandom()
	if err != nil {
		fmt.Println("unable to run random functi", err)
	}
	fmt.Println(ret)
}