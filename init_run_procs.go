package common

import(
    "runtime"
    "fmt"
)

// TODO: whether print debug info to stdout

// Set golang program use maxProc CPUs
// -1 will use all CPUS
func InitRunProcs(maxProc int) {
    if maxProc < 1 {
        maxProc = runtime.NumCPU()
    }

    runtime.GOMAXPROCS(maxProc)
    fmt.Printf("NumCPU[%d], golang use [%d]CPUs\n", runtime.NumCPU(), maxProc)
}
