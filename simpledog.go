/*
A simple process watchdog
Usage: simpledog <process to start with arguments>
*/
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func usage() {
	selfname := filepath.Base(os.Args[0])
	fmt.Printf("Usage: simpledog <process to start>\n", selfname)
}

func main() {
	if (len(os.Args) == 1) {
		usage()
		os.Exit(1)
	}

	argsWithoutProg := os.Args[1:]
	log.Printf("simpledog starting: %s\n", argsWithoutProg)

	subName := argsWithoutProg[0]
	subArgs := argsWithoutProg[1:]
	cmd := exec.Command(subName, subArgs...)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)

	// TODO(kangas) does not pass through stdout/stderr from cmd
	// TODO(kangas) does correctly return error code
}

/*
References
https://gobyexample.com/spawning-processes
http://www.darrencoxall.com/golang/executing-commands-in-go/
	- has example of killing
*/