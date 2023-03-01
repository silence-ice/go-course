package main

import (
	"fmt"
	"os/exec"
	"sync"
)

type result struct {
	output []byte
	err    error
}

var (
	cmd        *exec.Cmd
	resultChan chan result
	data       result
	wg         sync.WaitGroup
)

func main() {
	resultChan = make(chan result)

	wg.Add(1)
	go func() {
		defer wg.Done()

		cmd = exec.Command("ls", "-l", "-a")
		out, err := cmd.CombinedOutput()
		//fmt.Println(string(out),err)

		resultChan <- result{
			output: out,
			err:    err,
		}
	}()

	//time.Sleep(time.Second)
	data = <-resultChan
	wg.Wait()

	fmt.Println(string(data.output))
}
