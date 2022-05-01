package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_ParsedArgs_Defaults(t *testing.T) {
	tsActual, useTime, useTz, showHelp, showVersion := getParsedArgs([]string{})

	tsExpected := fmt.Sprintf("%d", time.Now().Unix())

	if tsActual != tsExpected {
		t.Fatalf("expected time to default to now (%s), got %s", tsExpected, tsActual)
	}

	if useTime {
		t.Fatal("use time should be off by default")
	}
	if useTz {
		t.Fatal("use tz should be off by default")
	}
	if showHelp {
		t.Fatal("show help should be off by default")
	}
	if showVersion {
		t.Fatal("show version should be off by default")
	}
}

func Test_ParsedArgs_ShowHelp(t *testing.T) {
	tsActual, useTime, useTz, showHelp, showVersion := getParsedArgs([]string{"-h"})

	tsExpected := fmt.Sprintf("%d", time.Now().Unix())

	if tsActual != tsExpected {
		t.Fatalf("expected time to default to now (%s), got %s", tsExpected, tsActual)
	}

	if useTime {
		t.Fatal("use time should be off by default")
	}
	if useTz {
		t.Fatal("use tz should be off by default")
	}
	if !showHelp {
		t.Fatal("show help should be turned on")
	}
	if showVersion {
		t.Fatal("show version should be off by default")
	}
}

func Test_ParsedArgs_ShowVersion(t *testing.T) {
	tsActual, useTime, useTz, showHelp, showVersion := getParsedArgs([]string{"-v"})

	tsExpected := fmt.Sprintf("%d", time.Now().Unix())

	if tsActual != tsExpected {
		t.Fatalf("expected time to default to now (%s), got %s", tsExpected, tsActual)
	}

	if useTime {
		t.Fatal("use time should be off by default")
	}
	if useTz {
		t.Fatal("use tz should be off by default")
	}
	if showHelp {
		t.Fatal("show help should be off by default")
	}
	if !showVersion {
		t.Fatal("show version should be turned on")
	}
}

func Test_ParsedArgs_UseTime(t *testing.T) {
	tsActual, useTime, useTz, showHelp, showVersion := getParsedArgs([]string{"-t"})

	tsExpected := fmt.Sprintf("%d", time.Now().Unix())

	if tsActual != tsExpected {
		t.Fatalf("expected time to default to now (%s), got %s", tsExpected, tsActual)
	}

	if !useTime {
		t.Fatal("use time should be turned on")
	}
	if useTz {
		t.Fatal("use tz should be off by default")
	}
	if showHelp {
		t.Fatal("show help should be off by default")
	}
	if showVersion {
		t.Fatal("show version should be off by default")
	}
}

func Test_ParsedArgs_ShowTzDoubleT(t *testing.T) {
	tsActual, useTime, useTz, showHelp, showVersion := getParsedArgs([]string{"-tt"})

	tsExpected := fmt.Sprintf("%d", time.Now().Unix())

	if tsActual != tsExpected {
		t.Fatalf("expected time to default to now (%s), got %s", tsExpected, tsActual)
	}

	if !useTime {
		t.Fatal("use time should be turned on")
	}
	if !useTz {
		t.Fatal("use tz should be turned on")
	}
	if showHelp {
		t.Fatal("show help should be off by default")
	}
	if showVersion {
		t.Fatal("show version should be off by default")
	}
}

func Test_ParsedArgs_ShowTzDoubleFlags(t *testing.T) {
	tsActual, useTime, useTz, showHelp, showVersion := getParsedArgs([]string{"-t", "-t"})

	tsExpected := fmt.Sprintf("%d", time.Now().Unix())

	if tsActual != tsExpected {
		t.Fatalf("expected time to default to now (%s), got %s", tsExpected, tsActual)
	}

	if !useTime {
		t.Fatal("use time should be turned on")
	}
	if !useTz {
		t.Fatal("use tz should be turned on")
	}
	if showHelp {
		t.Fatal("show help should be off by default")
	}
	if showVersion {
		t.Fatal("show version should be off by default")
	}
}

func Test_ParsedArgs_Timestamp(t *testing.T) {
	tsExpected := "1653881338"
	tsActual, useTime, useTz, showHelp, showVersion := getParsedArgs([]string{tsExpected})

	if tsActual != tsExpected {
		t.Fatalf("expected time to default to now (%s), got %s", tsExpected, tsActual)
	}

	if useTime {
		t.Fatal("use time should be off by default")
	}
	if useTz {
		t.Fatal("use tz should be off by default")
	}
	if showHelp {
		t.Fatal("show help should be off by default")
	}
	if showVersion {
		t.Fatal("show version should be off by default")
	}
}

func Test_GetTimestamp_ParsesUnixSeconds(t *testing.T) {
	tsRaw := 1653881338
	ts := getTimestamp(tsRaw)
	if ts.Year() != 2022 && ts.Month() != time.Month(5) && ts.Day() != 30 {
		t.Log(ts)
		t.Fatalf("expected 2022-05-30, got: %v", ts)
	}
}

func Test_GetTimestamp_ParsesUnixMilli(t *testing.T) {
	tsRaw := 1629484202017
	ts := getTimestamp(tsRaw)
	if ts.Year() != 2021 && ts.Month() != time.Month(8) && ts.Day() != 20 {
		t.Log(ts)
		t.Fatalf("expected 2021-08-20, got: %v", ts)
	}
}
