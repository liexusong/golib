Daemon
------
usage:
```go
package main

import (
	"fmt"
	"github.com/liexusong/golib/daemon"
	"log"
	"time"
)

func main() {
	// Daemon() return child process id and error
	// if pid == 0 then is child process
	// if pid > 0 then is parent process
	pid, err := daemon.Daemon()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	if pid == 0 {
		for i := 0; i < 100; i++ {
			log.Println("child...")
			time.Sleep(time.Second)
		}
	} else {
		fmt.Println("parent...")
	}
}
```