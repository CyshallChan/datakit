// Copyright © 2016 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package jwalterweatherman

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	TRACE    *log.Logger
	DEBUG    *log.Logger
	INFO     *log.Logger
	WARN     *log.Logger
	ERROR    *log.Logger
	CRITICAL *log.Logger
	FATAL    *log.Logger

	LOG      *log.Logger
	FEEDBACK *Feedback

	defaultNotepad *Notepad
)

func reloadDefaultNotepad() {
	TRACE = defaultNotepad.TRACE
	DEBUG = defaultNotepad.DEBUG
	INFO = defaultNotepad.INFO
	WARN = defaultNotepad.WARN
	ERROR = defaultNotepad.ERROR
	CRITICAL = defaultNotepad.CRITICAL
	FATAL = defaultNotepad.FATAL

	LOG = defaultNotepad.LOG
	FEEDBACK = defaultNotepad.FEEDBACK
}

func init() {
	defaultNotepad = NewNotepad(LevelError, LevelWarn, os.Stdout, ioutil.Discard, "", log.Ldate|log.Ltime)
	reloadDefaultNotepad()
}

// SetLogThreshold set the log threshold for the default notepad. Trace by default.
func SetLogThreshold(threshold Threshold) {
	defaultNotepad.SetLogThreshold(threshold)
	reloadDefaultNotepad()
}

// SetLogOutput set the log output for the default notepad. Discarded by default.
func SetLogOutput(handle io.Writer) {
	defaultNotepad.SetLogOutput(handle)
	reloadDefaultNotepad()
}

// SetStdoutThreshold set the standard output threshold for the default notepad.
// Info by default.
func SetStdoutThreshold(threshold Threshold) {
	defaultNotepad.SetStdoutThreshold(threshold)
	reloadDefaultNotepad()
}

// SetPrefix set the prefix for the default logger. Empty by default.
func SetPrefix(prefix string) {
	defaultNotepad.SetPrefix(prefix)
	reloadDefaultNotepad()
}

// SetFlags set the flags for the default logger. "log.Ldate | log.Ltime" by default.
func SetFlags(flags int) {
	defaultNotepad.SetFlags(flags)
	reloadDefaultNotepad()
}

// Level returns the current global log threshold.
func LogThreshold() Threshold {
	return defaultNotepad.logThreshold
}

// Level returns the current global output threshold.
func StdoutThreshold() Threshold {
	return defaultNotepad.stdoutThreshold
}

// GetStdoutThreshold returns the defined Treshold for the log logger.
func GetLogThreshold() Threshold {
	return defaultNotepad.GetLogThreshold()
}

// GetStdoutThreshold returns the Treshold for the stdout logger.
func GetStdoutThreshold() Threshold {
	return defaultNotepad.GetStdoutThreshold()
}

// LogCountForLevel returns the number of log invocations for a given threshold.
func LogCountForLevel(l Threshold) uint64 {
	return defaultNotepad.LogCountForLevel(l)
}

// LogCountForLevelsGreaterThanorEqualTo returns the number of log invocations
// greater than or equal to a given threshold.
func LogCountForLevelsGreaterThanorEqualTo(threshold Threshold) uint64 {
	return defaultNotepad.LogCountForLevelsGreaterThanorEqualTo(threshold)
}

// ResetLogCounters resets the invocation counters for all levels.
func ResetLogCounters() {
	defaultNotepad.ResetLogCounters()
}
