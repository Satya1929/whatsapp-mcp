package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exeDir := filepath.Dir(exePath)

	logFilePath := filepath.Join(exeDir, "mcp_traffic.log")
	f, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	realExePath := filepath.Join(exeDir, "whatsapp-mcp-real.exe")
	cmd := exec.Command(realExePath, os.Args[1:]...)
	cmd.Dir = exeDir

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Stderr = os.Stderr // Pass stderr through

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// Forward Stdin to Process
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			f.WriteString(fmt.Sprintf("IDE -> SERVER: %s\n", line))
			fmt.Fprintln(stdin, line)
		}
		stdin.Close()
		f.WriteString("IDE stdin closed. Killing child process...\n")
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
	}()

	// Forward Process to Stdout
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			f.WriteString(fmt.Sprintf("SERVER -> IDE: %s\n", line))
			fmt.Fprintln(os.Stdout, line)
		}
	}()

	if err := cmd.Wait(); err != nil {
		f.WriteString(fmt.Sprintf("Process exited with error: %v\n", err))
	} else {
		f.WriteString("Process exited successfully\n")
	}
}
