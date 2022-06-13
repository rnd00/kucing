package helpers

import (
	"fmt"
	"time"
)

type severity int

const (
	debug severity = iota
	info
	fatal
)

func (s severity) String() string {
	a := []string{"Debug", "Info", "Fatal"}
	return a[s]
}

type Log struct {
	time     time.Time
	severity severity
	message  string
}

func NewLogMessage(s severity, msg string) *Log {
	return &Log{
		time:     time.Now(),
		severity: s,
		message:  msg,
	}
}

func NewLogDebug(msg string) *Log {
	return &Log{
		time:     time.Now(),
		severity: debug,
		message:  msg,
	}
}

func NewLogInfo(msg string) *Log {
	return &Log{
		time:     time.Now(),
		severity: info,
		message:  msg,
	}
}

func NewLogError(e err) *Log {
	return &Log{
		time:     time.Now(),
		severity: fatal,
		message:  e.String(),
	}
}

func (l *Log) String() string {
	return fmt.Sprintf("[%s] %s | %s", time.Now().Format(time.RFC3339), l.severity, l.message)
}

func (l *Log) Print() {
	fmt.Println(l)
}
