package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

func myAtoi(s string) int {
	statusMachine := map[string][]string{
		"init":    {"init", "start", "parsing", "end"},
		"start":   {"end", "end", "parsing", "end"},
		"parsing": {"end", "end", "parsing", "end"},
	}
	var status = "init"
	var idx int
	var result string
	for _, i := range s {
		switch {
		case i == ' ':
			idx = 0
		case i == '-' || i == '+':
			idx = 1
		case unicode.IsDigit(i):
			idx = 2
		default:
			idx = 3
		}
		status = statusMachine[status][idx]
		if status == "start" || status == "parsing" {
			result = result + string(i)
		} else if status == "end" {
			break
		}
	}
	r, _ := strconv.Atoi(result)
	if r < math.MinInt32 {
		return math.MinInt32
	} else if r > math.MaxInt32 {
		return math.MaxInt32
	}
	return r

}
func main() {
	fmt.Println(myAtoi(" -42"))
}
