package main

import (
	"io/ioutil"
	"strings"

	"github.com/Bitnick2002/goa/log"
)

func main() {
	bytes, err := ioutil.ReadFile("website.txt")
	if err != nil {
		log.Error(err)
	}
	str := string(bytes)
	websites := strings.Split(str, "\n")
	filter := "[/S]*"
	for i, s := range websites {
		if s == "\r" || s == "" {
			continue
		}
		if i != 0 {
			filter += "|"
		}
		filter += "(" + s + ")"
	}
	filter += "[/S]*"
	log.Info(filter)
	ioutil.WriteFile("filter.txt", []byte(filter), 0x666)
}
