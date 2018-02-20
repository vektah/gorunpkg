package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: package <args passed to binary>")
		os.Exit(1)
	}

	pkg, err := resolvePkg(os.Args[1])
	if err != nil {
		fmt.Println("pkg: " + err.Error())
		os.Exit(1)
	}

	tmpfile := filepath.Join(os.TempDir(), "gorunpkg_"+filepath.Base(pkg))
	if runtime.GOOS == "windows" {
		tmpfile += ".exe"
	}

	passthrough("go", "build", "-i", "-o", tmpfile, pkg)
	passthrough(tmpfile, os.Args[2:]...)
}

func passthrough(command string, args ...string) {
	var err error
	cmd := exec.Command(command, args...)
	cmd.Env = os.Environ()
	cmd.Dir, err = os.Getwd()
	if err != nil {
		fmt.Println("cwd: " + err.Error())
		os.Exit(1)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("stdout: " + err.Error())
		os.Exit(1)
	}
	defer stdout.Close()
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("stderr: " + err.Error())
		os.Exit(1)
	}
	defer stderr.Close()

	go func() {
		io.Copy(os.Stdout, stdout)
	}()

	go func() {
		io.Copy(os.Stderr, stderr)
	}()

	err = cmd.Run()

	if exitErr, ok := err.(*exec.ExitError); ok {
		if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
			os.Exit(status.ExitStatus())
		}
		os.Exit(1)
	}
	if err != nil {
		os.Exit(1)
	}
}
