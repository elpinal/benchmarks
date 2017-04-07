// +build ignore

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	count := flag.Int("count", 1, "Run benchmark n times")
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "usage: go run xtime.go [-count n] prog_name")
		os.Exit(2)
	}
	if err := bench(*count, flag.Arg(0)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func bench(count int, prog string) error {
	start := time.Now()
	for i := 0; i < count; i++ {
		cmd := exec.Command(prog)
		if err := cmd.Run(); err != nil {
			return err
		}
		if time.Since(start) > 10*time.Minute {
			return errors.New("timeout: 10 minutes elapsed")
		}
	}
	fmt.Println(time.Since(start) / time.Duration(count))
	return nil
}
