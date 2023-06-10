package main

import (
    "os/exec"
    "os"
    "syscall"
    "fmt"
)

func my_exec(cmd_array []string) error {
    cmd := exec.Command(cmd_array[0], cmd_array[1:]...)

    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    cmd.SysProcAttr = &syscall.SysProcAttr{}

    return cmd.Run()
}

func main() {
    cmd := []string{"keepassxc-cli", "db-info", "Password.kdbx"}

    err := my_exec(cmd)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
}
