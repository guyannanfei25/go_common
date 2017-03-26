package common

import (
    "testing"
    "time"
)

func TestInitPidFile(t *testing.T) {
    if err := InitPidFile("/tmp/test_pid.pid"); err != nil {
        t.Errorf("InitPidFile failed[%s]\n", err)
    }
}

func TestInitRunProcs(t *testing.T) {
    InitRunProcs(-1)
}

func TestIntervalGC(t *testing.T) {
    go IntervalGC(2, 3)

    for i := 0; i < 100; i++ {
        bigBuf := make([]string, 100000, 100000)
        bigBuf = append(bigBuf, "testdata")
    }

    time.Sleep(9 * time.Second)
}
