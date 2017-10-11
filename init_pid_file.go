package common

import (
    "os"
    "fmt"
    "strconv"
    "io/ioutil"
)

// write process pid to pidFile
func InitPidFile(pidFile string) error {
    if len(pidFile) == 0 {
        return fmt.Errorf("InitPidFile get no pidFile conf[%s]", pidFile)
    }

    if err := ioutil.WriteFile(pidFile, []byte(strconv.Itoa(os.Getpid())), 0777); err != nil {
        return err
    }

    return nil
}
