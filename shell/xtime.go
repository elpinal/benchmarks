// +build ignore

package main

import (
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
	for i := 0; i < count; i++ {
		start := time.Now()
		n := 100000
		for i := 0; i < n; i++ {
			if time.Since(start) > 800*time.Millisecond {
				n = i
				break
			}
			cmd := exec.Command(prog)
			if err := cmd.Run(); err != nil {
				return err
			}
		}
		fmt.Printf("%v\t%v\n", n, time.Since(start)/time.Duration(n))
	}
	return nil
}
