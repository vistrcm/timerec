package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		start := time.Now()
		fmt.Print(format(start), " ")

		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		used := duration(time.Since(start))
		fmt.Println(used)
	}
}

func duration(used time.Duration) string {
	return fmt.Sprintf("%d", int(math.Round(used.Minutes())))
}

func format(t time.Time) string {
	return fmt.Sprintf("%02d%02d", t.Hour(), t.Minute())
}
