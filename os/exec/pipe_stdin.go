package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, _ := cmd.Output()
	log.Printf("out:%s\n", out)

	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("finished:%s\n", out)
}