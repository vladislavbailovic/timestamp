package main

import (
	"log"
	"strconv"
	"time"
)

const YEAR_DIFF_CUTOFF = 1000

func main() {
	useTime := false
	tsRaw, err := strconv.Atoi(getTime())
	if err != nil {
		log.Fatal(err)
	}
	ts := getTimestamp(tsRaw)

	format := "2006-01-02"
	if useTime {
		format += "15:04:05"
	}
	log.Println(ts.UTC().Format(format))
}

func getTime() string {
	return "1653881338"    // s ts
	return "1629484202017" // ms ts
}

func getTimestamp(tsRaw int) time.Time {
	tsS := time.Unix(int64(tsRaw), 0)
	tsMs := time.UnixMilli(int64(tsRaw))
	log.Println(tsS.Year(), tsMs.Year(), tsS.Year()-tsMs.Year())
	if tsS.Year()-tsMs.Year() > YEAR_DIFF_CUTOFF {
		return tsMs
	}
	return tsS
}
