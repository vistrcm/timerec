package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		start := time.Now()

		task, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		used := duration(time.Since(start))
		fmt.Printf("\033[1A")
		ct := cleanTask(task)
		fmt.Printf("%s %s %s\n", ct, format(start), used)
	}
}

func cleanTask(task string) string {
	return strings.Trim(task, "\n")
}

func duration(used time.Duration) string {
	return fmt.Sprintf("%d", int(math.Round(used.Minutes())))
}

func format(t time.Time) string {
	return fmt.Sprintf("%02d%02d", t.Hour(), t.Minute())
}
