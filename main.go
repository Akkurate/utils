package main

import (
	"time"

	"github.com/Akkurate/utils/logging"
)

func main() {

	logging.SetLevel(logging.Everything)

	m := logging.Measure()

	time.Sleep(1 * time.Second)
	logging.Trace("Hello trace")
	logging.Debug("Hello Debug")
	logging.Info("Hello info")
	logging.Warn("Hello warn")

	logging.Error("Hello error")
	logging.Fatal("Hello fatal")

	logging.Level(logging.InfoLevel, "Hello (info)")
	logging.Level(logging.TraceLevel, "Hello (trace)")

	logging.Level(logging.DebugLevel, "Hello (debug1)")
	logging.Level(logging.DebugLevel-1, "Hello (debug2)")
	logging.Level(logging.DebugLevel-2, "Hello (debug2)")
	logging.Level(200, "Hello")

	m.Print("program", "It took")

	logging.Info(`
	Hello
	World  oi    moi
	`)
}
