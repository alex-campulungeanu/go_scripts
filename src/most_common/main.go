package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func readFromStdin() chan string {
	buffer := make([]byte, 1024)
	out := make(chan string)
	go func() {
		for {
			count, err := os.Stdin.Read(buffer)
			if err == io.EOF{
				break
			}
			out <- string(buffer[:count])
		}
	}()
	return out
}

func main() {
	stdin := readFromStdin()
	ticker := time.NewTicker(time.Second * 2)
	counts := make(map[string]int)

	for {
		select {
		case s := <- stdin:
			parts := strings.SplitN(s, ":", 2)
			msg := parts[1]
			counts[msg] ++
		case  <- ticker.C:
			max := 0
			maxMsg := ""
			for k, v := range counts {
				if v > max {
					max = v
					maxMsg = k
				}
			}
			fmt.Printf("Most common mesage : %s \n", maxMsg)	
			counts = make(map[string]int)
		}
	}
}