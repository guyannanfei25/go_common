package common

import (
    "os"
    // "logger"
    "fmt"
    "syscall"
    "strconv"
)

// write process pid to pidFile
func InitPidFile(pidFile string) error {
    if len(pidFile) == 0 {
        // logger.Errorf("InitPidFile get no pidFile conf[%s]\n", pidFile)
        return fmt.Errorf("InitPidFile get no pidFile conf[%s]", pidFile)
    }

    pidFd, err := os.OpenFile(pidFile, os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        // logger.Errorf("Create pid file err[%s]\n", err)
        return  err
    }

    if err := syscall.Flock(int(pidFd.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
        // logger.Errorf("Lock pid file[%s] err[%s]\n", pidFile, err)
        return err
    }

    pidFd.Truncate(0)
    pidFd.Write([]byte(strconv.Itoa(os.Getpid())))
    // logger.Debugf("Write pid[%d] to file[%s] Success\n", os.Getpid(), pidFile)

    return nil
}
