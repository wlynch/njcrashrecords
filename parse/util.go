package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func parseBool(s string) bool {
	return strings.ToUpper(strings.TrimSpace(s)) == "Y"
}

func parseInt(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Printf("parseInt: %v", err)
	}
	return i
}

func parseString(s string) string {
	return strings.TrimSpace(s)
}

func parseTime(d, t string) time.Time {
	val, err := time.Parse("01/02/2006 1504 MST", fmt.Sprintf("%s %s EST", d, t))
	if err != nil {
		log.Printf("parseTime: %v", err)
	}
	return val
}
