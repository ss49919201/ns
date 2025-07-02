package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println(time.Now().UnixNano())
		return
	}

	input := args[0]

	if nanoseconds, err := strconv.ParseInt(input, 10, 64); err == nil {
		t := time.Unix(0, nanoseconds)
		fmt.Println(t.Format(time.RFC3339))
		return
	}

	if t, err := time.Parse(time.RFC3339, input); err == nil {
		fmt.Println(t.UnixNano())
		return
	}

	if t, err := time.Parse("2006-01-02 15:04:05", input); err == nil {
		fmt.Println(t.UnixNano())
		return
	}

	if t, err := time.Parse("2006-01-02", input); err == nil {
		fmt.Println(t.UnixNano())
		return
	}

	fmt.Fprintf(os.Stderr, "Error: Unable to parse input '%s' as timestamp or date\n", input)
	os.Exit(1)
}