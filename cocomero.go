package main

import (
       "fmt"
       "os/exec"
       "log"
)


func getHeadOutput() {
        cmd := exec.Command("uptime")
        stdoutStderr, err := cmd.CombinedOutput()
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("%s\n", stdoutStderr)

}

func main() {
  getHeadOutput()
}
