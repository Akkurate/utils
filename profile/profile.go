package profile

import (
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/Akkurate/utils/logging"
)

// Executes pprof cpu profiler and saves result to cpu.prof in project root.
// Profiling is stopped by calling ProfileCPUStop() . Example:
//  func ProfileMe() {
//  	p := log.ProfileCPU()
//  	defer ProfileCPUstop(p)
//  	... regular code starts here
//  }
func ProfileCPUStart() *os.File {
	f, err := os.Create("cpu.prof")
	if err != nil {
		logging.Fatal("could not create CPU profile: ", err)
		f.Close()
		return nil
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		logging.Fatal("could not start CPU profile: ", err)
		f.Close()
		return nil
	}
	return f
}

// Stops CPU profiling. Call by deferring.
func ProfileCPUstop(f *os.File) {
	pprof.StopCPUProfile()
	f.Close()
}

// Executes pprof memory profiler and saves result to mem.prof in project root.
func ProfileMem() {
	f, err := os.Create("mem.prof")
	if err != nil {
		logging.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()
	runtime.GC()
	if err := pprof.WriteHeapProfile(f); err != nil {

		logging.Fatal("could not write memory profile: ", err)
	}
}
