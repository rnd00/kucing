package helpers

import (
	"fmt"
	"time"
)

type severity int

const (
	debug severity = iota
	info
	warn
	fatal
)

func (s severity) String() string {
	a := []string{"Debug", "Info", "Warn", "Fatal"}
	return a[s]
}

type log struct {
	time     time.Time
	severity severity
	message  string
}

func NewLogMessage(s severity, msg string) *log {
	return &log{
		time:     time.Now(),
		severity: s,
		message:  msg,
	}
}

func (l *log) String() string {
	return fmt.Sprintf("[%s] %s | %s", time.Now().Format(time.RFC3339), l.severity, l.message)
}

func (l *log) Print() {
	fmt.Println(l)
}

func RunLog() {
	// add later
}
