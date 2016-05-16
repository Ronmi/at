package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func help() {
	fmt.Println(`Accepted time format:
  You can specify absolute time in "2006-01-02 15:04:05" or "15:04:05" format, or
  any acceptable format in https://golang.org/pkg/time/#ParseDuration`)
	os.Exit(1)
}

func main() {
	tstr := strings.Join(os.Args[1:], " ")

	// test if it is absolute time
	n := time.Now()
	loc := n.Local().Location()

	// test full format
	if t, err := time.ParseInLocation("2006-01-02 15:04:05", tstr, loc); err == nil {
		absolute(t)
		return
	}

	// test short format
	if t, err := time.ParseInLocation("2006-01-02 15:04:05", n.Format("2006-01-02")+" "+tstr, loc); err == nil {
		absolute(t)
		return
	}

	// test if it is relative time
	if d, err := time.ParseDuration(tstr); err == nil {
		relative(d)
		return
	}

	help()
}

func absolute(t time.Time) {
	relative(t.Sub(time.Now()))
}

func relative(d time.Duration) {
	time.Sleep(d)

	c := exec.Command("bash")
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	c.Run() // discard error since stderr had been redirected
}
