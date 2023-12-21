package main

import (
	"fmt"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var acumulated int

func getRandom() (int, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return 0, fmt.Errorf("unabel to get the caller program")
	}
	nodePath := path.Dir(filename)

	prg := "node"
	arg1 := filepath.Join(nodePath, "index.js")
	arg2 := "33"
	arg3 := "44"
	cmd := exec.Command(prg, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		return 0, fmt.Errorf("unable to run command")
	}
	formattedOutput := strings.Split(string(stdout), "\n")
	res, err2 := strconv.Atoi(formattedOutput[0])
	if err2 != nil {
		return 0, fmt.Errorf("unable to run command")
	}
	return res, nil
}

func main() {
	ban_count := 0
	sum := 0
	// startTime := time.Now()
	for j := 0; j < 5; j++ {
		for i := 0; i < 4; i++ {
			// wg.Add(1)
			output, err := getRandom()
			if err != nil {
				fmt.Println("unable to run random functi", err)
			}
			sum += output
			fmt.Printf("lucy number %v: %v\n", i, output)
		}
		// executionTime := time.Since(startTime)
		fmt.Printf("sum is %v\n", sum)
		if sum %2 == 0 {
			ban_count += 1
		} else {
			ban_count += 0
		}
	}
	fmt.Printf("ban_count %v\n", ban_count)
	// fmt.Println("program execution time", executionTime)
}
