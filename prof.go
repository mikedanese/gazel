package main

import (
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/golang/glog"
)

type profiler struct {
	path string
	cf   *os.File
}

func newProfiler(path string) *profiler {
	return &profiler{path: path}
}

func (p *profiler) Start() {
	cf, err := os.Create(p.path + ".cpu.prof")
	if err != nil {
		glog.Fatal("could not create CPU profile: ", err)
	}
	if err := pprof.StartCPUProfile(cf); err != nil {
		glog.Fatal("could not start CPU profile: ", err)
	}
	p.cf = cf
}
func (p *profiler) Stop() {
	pprof.StopCPUProfile()

	mf, err := os.Create(p.path + ".mem.prof")
	if err != nil {
		glog.Fatal("could not create memory profile: ", err)
	}
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(mf); err != nil {
		glog.Fatal("could not write memory profile: ", err)
	}
	mf.Close()
}
