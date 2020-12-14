package tools

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Verbose struct {
	Verb bool
	W    io.Writer
}

// if verbose is set printf.
func (v *Verbose) Printf(format string, a ...interface{}) {
	if v.W == nil {
		v.W = os.Stdout
	}
	if v.Verb {
		fmt.Fprintf(v.W, format, a...)
	}
}

// if verbose is sert println.
func (v *Verbose) Println(a ...interface{}) {
	if v.W == nil {
		v.W = os.Stdout
	}
	if v.Verb {
		fmt.Fprintln(v.W, a...)
	}
}

// Error handle Output to any io.writer
func ErrorHandle(err error, message string, w io.Writer, exit bool) {
	if err != nil {
		fmt.Printf("ERROR: %v, %v\n")
	}
	if exit {
		os.Exit(1)
	}
}

// Error handle output to stderr
func ErrorHandleErr(err error, message string, exit bool) {
	ErrorHandle(err, message, os.Stderr, exit)
}

// Error handle output to stdout
func ErrorHandleOut(err error, message string, exit bool) {
	ErrorHandle(err, message, os.Stdout, exit)
}

// Error handle HTTP Send 500 response.
func ErrorHandle500(err error, message string, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Fprintf(w, "ERROR:  %v", err)
	}
}
