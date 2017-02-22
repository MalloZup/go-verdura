package main

import (
	"fmt"
	"log"
	"os/exec"
)

// this function get the latest cucumber result
func getHeadOutput(url string) {
	cmd := exec.Command("wget", url, "-O", "last.html")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}

// get name of failed steps
func get_failed_steps() {
	fmt.Printf("\n NOT yet implemented.")
}

func main() {
	// we can assume that this is constant.
	const url_head = "http://10.162.210.226/workspace/manager-Head-sumaform-cucumber/last.html"
	getHeadOutput(url_head)
	get_failed_steps()
}
