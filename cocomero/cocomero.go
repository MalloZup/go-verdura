package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"strings"
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
func get_failed_steps(output string) (numFailed int, nameFailed []string) {
	// rio_file">features/spacewalk_debug.feature:16</span><h3 id="scenario_442"><span class="keyword">Scenario:</span> <span class="val">Check that no scheduled events have failed on manager server</span></h3><ol><script>makeRed('scenario_442');</script><li id='features_spacewalk_debug_feature_17' class='
	failed_steps := regexp.MustCompile("step failed")
	indexes_fail := failed_steps.FindAllStringIndex(output, -1)

	for _, index := range indexes_fail {
		fmt.Println("\n=============================\n\n")
		fmt.Println(strings.Fields(strings.Split(output[index[0]-500:index[0]], "li id=")[1])[0])
	}
	numFailed = len(indexes_fail)
	return numFailed, nameFailed
}

func main() {
	cucumber.getHeadOutput()
	output := cucumber.readLastResults()
	get_failed_steps(output)
}
