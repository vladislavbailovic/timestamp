package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const yearDiffCutoff = 1000

var GitCommitHash string = "development"
var BuildDate string = "right now"

func getVersion() string {
	return strings.Join([]string{
		fmt.Sprintf("Version:    %s", GitCommitHash),
		fmt.Sprintf("Build date: %s", BuildDate),
	}, "\n")
}

func main() {
	tsRaw, useTime, useTz, showHelp, showVersion := getParsedArgs(os.Args[1:])
	if showHelp {
		fmt.Println(getHelp())
		os.Exit(0)
	}

	if showVersion {
		fmt.Println(getVersion())
		os.Exit(0)
	}

	tsSrc, err := strconv.Atoi(tsRaw)
	if err != nil {
		fmt.Printf("ERROR: error parsing [%s]\n", tsRaw)
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

func getParsedArgs(args []string) (string, bool, bool, bool, bool) {
	showHelp := false
	showVersion := false
	useTime := false
	useTz := false
	tsRaw := fmt.Sprintf("%d", time.Now().Unix())

	for _, arg := range args {
		switch strings.Trim(arg, "-") {
		case "t":
			if useTime {
				useTz = true
			}
			useTime = true
		case "tt":
			useTime = true
			useTz = true
		case "v":
			showVersion = true
		case "h":
			showHelp = true
		default:
			tsRaw = arg
		}
	}

	return tsRaw, useTime, useTz, showHelp, showVersion
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
		fmt.Sprintf("%s-v: Show version", tab),
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
