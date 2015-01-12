/*
A simple process watchdog.
Spawns arguments as a subprocess.
Terminates subprocess when orphaned by parent.

Implemented with Go 1.4

Author: Matt Kangas <kangas@gmail.com>
Date: January 2015
*/
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
)

const appname string = "simpledog"

func usage() {
	fmt.Printf("Usage: %s <process to start>\n", appname)
}

func killIfOrphaned(cmd *exec.Cmd) {
	buf := make([]byte, 1024)
	for {
		_, err := os.Stdin.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println(appname, "detected EOF on stdin")
				break
			}
			fmt.Println(err)
		}
	}

	fmt.Println(appname, "killing cmd", cmd.Path)
	err := cmd.Process.Kill()
	if err != nil {
		log.Println(appname, "failed to kill", err)
	}
}

func main() {
	if len(os.Args) == 1 {
		usage()
		os.Exit(1)
	}

	args := os.Args[1:]
	log.Printf("%s starting: %s\n", appname, args)

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go killIfOrphaned(cmd)

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0
			// There is no platform independent way to retrieve
			// the exit code, but the following will work on Unix
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				os.Exit(status.ExitStatus())
			} else {
				log.Fatal(err)
			}
		}
	}
}

/*
References
https://gobyexample.com/spawning-processes
http://www.darrencoxall.com/golang/executing-commands-in-go/
	- has example of killing
*/
