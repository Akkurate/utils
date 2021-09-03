package logging

import (
	"fmt"
	"os"
	"runtime"

	"gopkg.in/gookit/color.v1"
)

// Info Info
func Info(msg string, a ...interface{}) {
	var s string
	if len(a) == 0 {
		s = color.Sprintf(" <suc>Info: </>" + msg + "\n")
	} else {
		s = color.Sprintf(" <suc>Info: </>"+msg+"\n", a...)
	}
	os.Stdout.Write([]byte(s))
}

func Level(level int, msg string, a ...interface{}) {

	/*0-100 TRACE
	100-200 DEBUG
	200-300 INFO
	300-400 WARN
	400-500 ERROR
	*/
	if level < 100 {
		Info(msg+fmt.Sprintf(" level=%v", level), a...)
	} else if level < 200 {
		Info(msg+fmt.Sprintf(" level=%v", level), a...)
	} else if level < 300 {
		Info(msg+fmt.Sprintf(" level=%v", level), a...)
	} else if level < 400 {
		Warn(msg+fmt.Sprintf(" level=%v", level), a...)
	} else if level < 500 {
		Error(msg+fmt.Sprintf(" level=%v", level), a...)
	}
}

// Success Success
func Success(msg string, a ...interface{}) {
	var s string
	if len(a) == 0 {
		s = color.Sprintf("    <green>✔</>  " + msg + "\n")
	} else {
		s = color.Sprintf("    <green>✔</>  "+msg+"\n", a...)
	}
	os.Stdout.Write([]byte(s))
}

// Print Print
func Print(msg string, a ...interface{}) {
	var s string
	if len(a) == 0 {
		s = color.Sprintf(" <suc>Info: </>" + msg + "\n")
	} else {
		s = color.Sprintf(" <suc>Info: </>"+msg+"\n", a...)
	}
	os.Stdout.Write([]byte(s))
}

func Use(a ...interface{}) {
}

// Error ErrorF
func Error(msg string, a ...interface{}) {
	var s string

	_, fn, line, _ := runtime.Caller(1)
	if len(a) == 0 {
		s = color.Sprintf(" <err>Error:</><danger> "+msg+"</>\n", a)

	} else {
		s = color.Sprintf(" <err>Error:</><danger> "+msg+"</>\n", a...)
	}
	s += color.Sprintf("       <gray>%v:%v</>\n", fn, line)

	os.Stderr.Write([]byte(s))
}

// Warn WarnF
func Warn(msg string, a ...interface{}) {
	var s string
	if len(a) == 0 {
		s = color.Sprintf(" <warn>WARN:</> " + msg + "\n")
	} else {
		s = color.Sprintf(" <warn>WARN:</> "+msg+"\n", a...)
	}
	os.Stdout.Write([]byte(s))
}
