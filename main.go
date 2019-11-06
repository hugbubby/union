package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	flag "github.com/spf13/pflag"
)

var quiet bool

func clog(msg string) {
	log.Println("map: ", msg)
}

func main() {
	var delimiter, newDelimiter string

	flag.StringVarP(&delimiter, "delimiter", "d", "\\s+", "The delimiter for your sets. Regexp permitted.")
	flag.StringVarP(&newDelimiter, "new-delimiter", "n", "\n", "The delimiter for the outputted set. Defaults to newline.")
	flag.Parse()

	rg, err := regexp.Compile(delimiter)
	if err != nil {
		fmt.Println("invalid regular expression for delimiter: ", err)
	}
	output := make([]string, 0)
	for _, v := range flag.Args() {
		output = append(output, rg.Split(v, -1)...)
	}
	if len(output) > 0 {
		sort.Strings(output)
		last := output[0]
		i := 1
		for i < len(output) {
			this := output[i]
			if strings.EqualFold(this, last) {
                output = append(output[:i], output[i+1:]...)
			} else {
				i++
			}
            last = this
		}
        for _, v := range output[:len(output)-1] {
            fmt.Print(v + newDelimiter)
        }
        fmt.Println(output[len(output)-1])
	}
}
