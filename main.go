package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const yearDiffCutoff = 1000

func main() {
	tsRaw, useTime, useTz, showHelp := getParsedArgs()
	if showHelp {
		fmt.Println(getHelp())
		os.Exit(0)
	}

	tsSrc, err := strconv.Atoi(tsRaw)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ts := getTimestamp(tsSrc)

	format := "2006-01-02"
	if useTime {
		format += "T15:04:05"
		if !useTz {
			ts = ts.UTC()
		} else {
			format += "ZZ07:00"
		}
	}
	fmt.Println(ts.Format(format))
}

func getParsedArgs() (string, bool, bool, bool) {
	showHelp := false
	useTime := false
	useTz := false
	tsRaw := fmt.Sprintf("%d", time.Now().Unix())

	for _, arg := range os.Args {
		switch strings.Trim(arg, "-") {
		case "t":
			if useTime {
				useTz = true
			}
			useTime = true
		case "tt":
			useTime = true
			useTz = true
		case "h":
			showHelp = true
		default:
			tsRaw = arg
		}
	}

	return tsRaw, useTime, useTz, showHelp
	// return "1653881338"    // s ts
	// return "1629484202017" // ms ts
}

func getHelp() string {
	tab := "\t"
	help := []string{
		"USAGE:",
		fmt.Sprintf("%s%s [-tth] TIMESTAMP", tab, os.Args[0]),
		"... where TIMESTAMP is a numeric string with UNIX timestamp in (milli)seconds",
		"OPTIONS:",
		fmt.Sprintf("%s-t: Show time (default: false).", tab),
		fmt.Sprintf("%s    If doubled up, also convert to local and show TZ (default: UTC)", tab),
		fmt.Sprintf("%s-h: Show help", tab),
	}

	return strings.Join(help, "\n")
}

func getTimestamp(tsRaw int) time.Time {
	tsS := time.Unix(int64(tsRaw), 0)
	tsMs := time.UnixMilli(int64(tsRaw))
	if tsS.Year()-tsMs.Year() > yearDiffCutoff {
		return tsMs
	}
	return tsS
}
