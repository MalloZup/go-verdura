package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

type Urls struct {
	head         string
	devel30      string
	head_file    string
	devel30_file string
}

var cucumber = Urls{"http://m226.mgr.suse.de/workspace/manager-Head-sumaform-cucumber/last.html",
	"http://m226.mgr.suse.de/workspace/manager-3.0-sumaform30/last.html",
	"last.html",
	"last30.html"}

// this function get the latest cucumber result
func (url Urls) getHeadOutput() {
	cmd := exec.Command("wget", url.head, "-O", url.head_file)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}

func (url Urls) readLastResults() string {
	data, err := ioutil.ReadFile(url.head_file)
	if err != nil {
		panic(err)
	}
	str := string(data)
	return str
}

// get name of failed steps
func get_failed_steps(output string) {
}

func main() {
	cucumber.getHeadOutput()
	output := cucumber.readLastResults()
	get_failed_steps(output)
}
