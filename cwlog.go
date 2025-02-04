package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

var (
	logData logDataStruct
)

type logDataStruct struct {
	logDesc  *os.File
	logName  string
	logReady bool

	logBuf      []string
	logBufLines int
	logBufLock  sync.Mutex
}

/*
 * Log this, can use printf arguments
 * Write to buffer, async write
 */
func doLog(withTrace, debug bool, format string, args ...interface{}) {

	if wasmMode {
		return
	}

	if !*debugMode && debug {
		return
	}

	var buf string

	if withTrace {
		// Get current time
		ctime := time.Now()
		// Get calling function and line
		_, filename, line, _ := runtime.Caller(1)
		// printf conversion
		text := fmt.Sprintf(format, args...)
		// Add current date
		date := fmt.Sprintf("%2v:%2v.%2v", ctime.Hour(), ctime.Minute(), ctime.Second())
		// Date, go file, go file line, text
		buf = fmt.Sprintf("%v: %15v:%5v: %v\n", date, filepath.Base(filename), line, text)
	} else {
		// Get current time
		ctime := time.Now()
		// printf conversion
		text := fmt.Sprintf(format, args...)
		// Add current date
		date := fmt.Sprintf("%2v:%2v.%2v", ctime.Hour(), ctime.Minute(), ctime.Second())
		// Date, go file, go file line, text
		buf = fmt.Sprintf("%v: %v\n", date, text)
	}

	if !logData.logReady || logData.logDesc == nil {
		fmt.Print(buf)
		return
	}

	// Add to buffer
	logData.logBufLock.Lock()
	logData.logBuf = append(logData.logBuf, buf)
	logData.logBufLines++
	logData.logBufLock.Unlock()
}

func logDaemon() {

	if wasmMode {
		return
	}

	go func() {
		for {
			logData.logBufLock.Lock()

			// Are there lines to write?
			if logData.logBufLines == 0 {
				logData.logBufLock.Unlock()
				time.Sleep(time.Millisecond * 100)
				continue
			}

			// Write line
			_, err := logData.logDesc.WriteString(logData.logBuf[0])
			if err != nil {
				fmt.Println("DoLog: WriteString failure")
				logData.logDesc.Close()
				logData.logDesc = nil
			}
			fmt.Print(logData.logBuf[0])

			// Remove line from buffer
			logData.logBuf = logData.logBuf[1:]
			logData.logBufLines--

			logData.logBufLock.Unlock()
		}
	}()
}

// Prep logger
func startLog() {
	if wasmMode {
		return
	}

	t := time.Now()

	// Create our log file names
	logData.logName = fmt.Sprintf("log/auth-%v-%v-%v.log", t.Day(), t.Month(), t.Year())

	// Make log directory
	errr := os.MkdirAll("log", os.ModePerm)
	if errr != nil {
		fmt.Print(errr.Error())
		return
	}

	// Open log files
	bdesc, errb := os.OpenFile(logData.logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	// Handle file errors
	if errb != nil {
		doLog(true, false, "An error occurred when attempting to create the log. Details: %s", errb)
		return
	}

	// Save descriptors, open/closed elsewhere
	logData.logDesc = bdesc
	logData.logReady = true

}
