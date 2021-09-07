package logging

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"time"

	"gopkg.in/gookit/color.v1"
)

var LoggingLevel int = 300

const (
	TraceLevel = 600 // 500 - 599
	DebugLevel = 500 // 400 - 499
	InfoLevel  = 400 // 300 - 399
	WarnLevel  = 300 // 200 - 299
	ErrorLevel = 200 // 100 - 199
	FatalLevel = 100 // 0   - 99
	Everything = 700
)

func stripText(text string) string {

	//text = strings.ReplaceAll(text, "\n", "\\n")
	space := regexp.MustCompile(`\s+`)
	text = space.ReplaceAllString(text, " ")

	return text
}

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
		s = color.Sprintf(" <suc>Info: </>" + stripText(msg) + "\n")
	} else {
		s = color.Sprintf(" <suc>Info: </>"+stripText(msg)+"\n", a...)
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
		s = color.Sprintf(" <cyan>Debug: </>" + stripText(msg) + "\n")
	} else {
		s = color.Sprintf(" <cyan>Debug: </>"+stripText(msg)+"\n", a...)
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
		s = color.Sprintf(" <magenta>Trace: </>" + stripText(msg) + "\n")
	} else {
		s = color.Sprintf(" <magenta>Trace: </>"+stripText(msg)+"\n", a...)
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
		Fatal(msg+fmt.Sprintf(" <gray>level=%v</>", level), a...)
	} else if level < ErrorLevel {
		Error(msg+fmt.Sprintf(" <gray>level=%v</>", level), a...)
	} else if level < WarnLevel {
		Warn(msg+fmt.Sprintf(" <gray>level=%v</>", level), a...)
	} else if level < InfoLevel {
		Info(msg+fmt.Sprintf(" <gray>level=%v</>", level), a...)
	} else if level < DebugLevel {
		Debug(msg+fmt.Sprintf(" <gray>level=%v</>", level), a...)
	} else if level < TraceLevel {
		Trace(msg+fmt.Sprintf(" <gray>level=%v</>", level), a...)
	}

}

// Success Success
func Success(msg string, a ...interface{}) {

	var s string
	if len(a) == 0 {
		s = color.Sprintf("    <green>✔</>  " + stripText(msg) + "\n")
	} else {
		s = color.Sprintf("    <green>✔</>  "+stripText(msg)+"\n", a...)
	}
	os.Stdout.Write([]byte(s))
}

// Print Print
func Print(msg string, a ...interface{}) {
	var s string
	if len(a) == 0 {
		s = color.Sprintf(" <suc>Info: </>" + stripText(msg) + "\n")
	} else {
		s = color.Sprintf(" <suc>Info: </>"+stripText(msg)+"\n", a...)
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
		s = color.Sprintf(" <warn>WARN:</> " + stripText(msg) + "\n")
	} else {
		s = color.Sprintf(" <warn>WARN:</> "+stripText(msg)+"\n", a...)
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
		s = color.Sprintf(" <err>Error:</><danger> "+stripText(msg)+"</>", a...)

	} else {
		s = color.Sprintf(" <err>Error:</><danger> "+stripText(msg)+"</>", a...)
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
		s = color.Sprintf(" <err>Fatal:</><danger> "+stripText(msg)+"</>", a...)

	} else {
		s = color.Sprintf(" <err>Fatal:</><danger> "+stripText(msg)+"</>", a...)
	}
	s += color.Sprintf("<gray> at: %v:%v</>\n", fn, line)

	os.Stderr.Write([]byte(s))
}

// TimeMeasure GoPython
type TimeMeasure struct {
	start time.Time
}

// Print print
func (t *TimeMeasure) Print(tag string, msg string) {
	duration := time.Since(t.start)
	ms := duration.Milliseconds()
	Info("%v <yellow>%vms</> <gray>tag=%v value=%v level=5000</>", stripText(msg), ms, tag, ms)
}

// GetMilliseconds GetMilliseconds
func (t *TimeMeasure) GetMilliseconds() int64 {
	duration := time.Since(t.start)
	return int64(duration.Milliseconds())
}

func Measure() *TimeMeasure {
	x := &TimeMeasure{
		start: time.Now(),
	}
	return x
}
