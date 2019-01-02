package daemon

import (
	"errors"
	"os"
	"runtime"
	"syscall"
)

func initDaemonRuntime() {
	syscall.Setsid() // create new session

	fd, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
	if err != nil {
		return
	}

	syscall.Dup2(int(fd.Fd()), int(os.Stdin.Fd()))
	syscall.Dup2(int(fd.Fd()), int(os.Stdout.Fd()))
	syscall.Dup2(int(fd.Fd()), int(os.Stderr.Fd()))

	if fd.Fd() > os.Stderr.Fd() {
		_ = fd.Close()
	}
}

func Daemon() (int, error) {
	if runtime.GOOS == "windows" {
		return -1, errors.New("not support windows operating system")
	}

	isDaemon := false
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-daemon" {
			isDaemon = true
		}
	}

	if isDaemon { // daemon process
		initDaemonRuntime()
		return 0, nil
	}

	procPath := os.Args[0]

	args := make([]string, 0, len(os.Args)+1)

	args = append(args, os.Args...)
	args = append(args, "-daemon")

	attr := &syscall.ProcAttr{
		Env: os.Environ(),
	}

	pid, err := syscall.ForkExec(procPath, args, attr)
	if err != nil {
		return -1, err
	}

	return pid, nil
}
