package daemon

import (
	"errors"
	"os"
	"runtime"
	"syscall"
)

const daemonFlag = "--daemon"

func initDaemonRuntime() error {
	// create new session
	_, err := syscall.Setsid()
	if err != nil {
		return err
	}

	err = os.Chdir("/")
	if err != nil {
		return err
	}

	// update stdin:stdout:stderr to null
	fd, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
	if err != nil {
		return err
	}

	err = syscall.Dup2(int(fd.Fd()), int(os.Stdin.Fd()))
	if err != nil {
		return err
	}

	err = syscall.Dup2(int(fd.Fd()), int(os.Stdout.Fd()))
	if err != nil {
		return err
	}

	err = syscall.Dup2(int(fd.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		return err
	}

	if fd.Fd() > os.Stderr.Fd() {
		err = fd.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func Daemon() (int, error) {
	if runtime.GOOS == "windows" {
		return -1, errors.New("unsupported windows operating system")
	}

	isDaemon := false
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == daemonFlag {
			isDaemon = true
		}
	}

	if isDaemon { // daemon process
		err := initDaemonRuntime()
		if err != nil {
			return -1, err
		}

		return 0, nil
	}

	procPath := os.Args[0]

	// add "--daemon" arg
	args := make([]string, 0, len(os.Args)+1)

	args = append(args, os.Args...)
	args = append(args, daemonFlag)

	attr := &syscall.ProcAttr{
		Env:   os.Environ(),
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
	}

	pid, err := syscall.ForkExec(procPath, args, attr)
	if err != nil {
		return -1, err
	}

	return pid, nil
}
