package logging

import (
	"fmt"
	"os"
	"runtime"

	"gopkg.in/gookit/color.v1"
)

var LoggingLevel int = 300

const (
	TraceLevel = 600
	DebugLevel = 500
	InfoLevel  = 400
	WarnLevel  = 300
	ErrorLevel = 200
	FatalLevel = 100
	Everything = 700
)

// SetLevel set level
func SetLevel(l int) {
	LoggingLevel = l
}

// Info Info
func Info(msg string, a ...interface{}) {

	if LoggingLevel < InfoLevel {
		return
	}

	var s string
	if len(a) == 0 {
		s = color.Sprintf(" <suc>Info: </>" + msg + "\n")
	} else {
		s = color.Sprintf(" <suc>Info: </>"+msg+"\n", a...)
	}
	os.Stdout.Write([]byte(s))
}

// Debug Debug
func Debug(msg string, a ...interface{}) {
	if LoggingLevel < DebugLevel {
		return
	}

	var s string
	if len(a) == 0 {
		s = color.Sprintf(" <cyan>Debug: </>" + msg + "\n")
	} else {
		s = color.Sprintf(" <cyan>Debug: </>"+msg+"\n", a...)
	}
	os.Stdout.Write([]byte(s))

}

// Trace Trace
func Trace(msg string, a ...interface{}) {
	if LoggingLevel < TraceLevel {
		return
	}
	var s string
	if len(a) == 0 {
		s = color.Sprintf(" <magenta>Trace: </>" + msg + "\n")
	} else {
		s = color.Sprintf(" <magenta>Trace: </>"+msg+"\n", a...)
	}
	os.Stdout.Write([]byte(s))

}

// Level logs with predefined levels
// 0-100 TRACE
// 100-200 DEBUG
// 200-300 INFO
// 300-400 WARN
// 400-500 ERROR
func Level(level int, msg string, a ...interface{}) {
	//current 300
	if level >= LoggingLevel {
		return
	}

	if level < FatalLevel {
		Fatal(msg+fmt.Sprintf(" level=%v", level), a...)
	} else if level < ErrorLevel {
		Error(msg+fmt.Sprintf(" level=%v", level), a...)
	} else if level < WarnLevel {
		Warn(msg+fmt.Sprintf(" level=%v", level), a...)
	} else if level < InfoLevel {
		Info(msg+fmt.Sprintf(" level=%v", level), a...)
	} else if level < DebugLevel {
		Debug(msg+fmt.Sprintf(" level=%v", level), a...)
	} else if level < TraceLevel {
		Trace(msg+fmt.Sprintf(" level=%v", level), a...)
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

// Warn WarnF
func Warn(msg string, a ...interface{}) {

	if LoggingLevel < WarnLevel {
		return
	}
	var s string
	if len(a) == 0 {
		s = color.Sprintf(" <warn>WARN:</> " + msg + "\n")
	} else {
		s = color.Sprintf(" <warn>WARN:</> "+msg+"\n", a...)
	}
	os.Stdout.Write([]byte(s))
}

// Error ErrorF
func Error(msg string, a ...interface{}) {

	if LoggingLevel < ErrorLevel {
		return
	}
	var s string

	_, fn, line, _ := runtime.Caller(1)
	if len(a) == 0 {
		s = color.Sprintf(" <err>Error:</><danger> "+msg+"</>", a...)

	} else {
		s = color.Sprintf(" <err>Error:</><danger> "+msg+"</>", a...)
	}
	s += color.Sprintf("<gray> at: %v:%v</>\n", fn, line)

	os.Stderr.Write([]byte(s))
}

// Error ErrorF
func Fatal(msg string, a ...interface{}) {

	if LoggingLevel < FatalLevel {
		return
	}
	var s string

	_, fn, line, _ := runtime.Caller(1)
	if len(a) == 0 {
		s = color.Sprintf(" <err>Fatal:</><danger> "+msg+"</>", a...)

	} else {
		s = color.Sprintf(" <err>Fatal:</><danger> "+msg+"</>", a...)
	}
	s += color.Sprintf("<gray> at: %v:%v</>\n", fn, line)

	os.Stderr.Write([]byte(s))
}
