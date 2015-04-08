package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

var precisionFlag string

func init() {
	flag.StringVar(&precisionFlag, "p", "ms", "Precision of the output. Options: h, m, s, ms, us, ns.")
}

func main() {
	flag.Parse()
	p := precision(precisionFlag)
	s := bufio.NewScanner(os.Stdin)

	data := make([]time.Duration, 0)
	count := 0

	for s.Scan() {
		if err := s.Err(); err != nil {
			log.Fatalf("error reading input:", err)
		}
		count++
		tok := s.Text()

		if d, err := time.ParseDuration(strings.TrimSpace(tok)); err != nil {
			log.Fatalf("not a duration string: %s", tok)
		} else {
			data = append(data, d)
		}
	}

	if len(data) == 0 {
		log.Println("no data passed")
		return
	}

	var sum time.Duration = 0
	var min time.Duration = math.MaxInt64
	var max time.Duration = math.MinInt64

	for _, v := range data {
		sum += v
		if v <= min {
			min = v
		}
		if v >= max {
			max = v
		}
	}
	avg := time.Duration(int64(sum) / int64(len(data)))

	fmt.Println()
	fmt.Printf("pts\t%v\n", len(data))
	fmt.Printf("avg\t%s\n", roundDown(avg, p))
	fmt.Printf("min\t%s\n", roundDown(min, p))
	fmt.Printf("max\t%s\n", roundDown(max, p))
}

func roundDown(d, precision time.Duration) time.Duration {
	return time.Duration((int64(d) / int64(precision)) * int64(precision))
}

func precision(flag string) time.Duration {
	switch flag {
	case "h":
		return time.Hour
	case "m":
		return time.Minute
	case "s":
		return time.Second
	case "ms":
		return time.Millisecond
	case "us":
		return time.Microsecond
	case "ns":
		return time.Nanosecond
	default:
		log.Fatalf("unknown precision: %s", flag)
		return 0
	}
}
