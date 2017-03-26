package common

import (
    // "github.com/zieckey/goini"
    "runtime"
    "runtime/debug"
    "time"
    "fmt"
)

// TODO: whether print debug info to stdout

// every check_interval seconds check alloc memory, when alloc memory larger than max_mem M,
// it will force gc process
// You can use json or ini conf, for example, ini:
//
// [gc]
// max_mem = 500
// check_interval = 20
//
// code like:
//    maxMem, _ := conf.SectionGetInt(section, "max_mem")
//    checkInterval,  _ := conf.SectionGetInt(section, "check_interval")
// use like: go IntervalGC(maxMem, checkInterval)
func IntervalGC(maxMem, checkInterval int) {
    ticker := time.NewTicker(time.Duration(checkInterval) * time.Second)
    var ms runtime.MemStats
    for {
        select {
        case <- ticker.C:
            runtime.ReadMemStats(&ms)
            alloc := ms.Alloc / 1024 / 1024 // 分配未释放多少M
            fmt.Printf("[DEBUG_GC] go routine num[%d], alloc[%d]M, head alloc[%d]M, StackInuse[%d]M, NextGC[%d]M\n",
            runtime.NumGoroutine(), alloc, ms.HeapAlloc/1024/1024, ms.StackInuse/1024/1024, ms.NextGC/1024/1024)
            if int(alloc) >= maxMem {
                debug.FreeOSMemory()
                runtime.ReadMemStats(&ms)
                fmt.Printf("[DEBUG_GC] go routine num[%d], alloc[%d]M, head alloc[%d]M, StackInuse[%d]M, NextGC[%d]M\n", 
                runtime.NumGoroutine(), alloc, ms.HeapAlloc/1024/1024, ms.StackInuse/1024/1024, ms.NextGC/1024/1024)
            }
        }
    }
}
