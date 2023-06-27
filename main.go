package main

import (
    "os/exec"
    "os"
    "syscall"
    "fmt"
    "github.com/keybase/go-keychain"
)

func myExec(cmd_array []string) error {
    cmd := exec.Command(cmd_array[0], cmd_array[1:]...)

    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    cmd.SysProcAttr = &syscall.SysProcAttr{}

    return cmd.Run()
}

func main() {
//    cmd := []string{"keepassxc-cli", "db-info", "Password.kdbx"}

//    err := myExec(cmd)
//    if err != nil {
//        fmt.Println("Error:", err)
//        return
//    }

    service := ""
    account := ""
    accessGroup := ""

    query := keychain.NewItem()
    query.SetSecClass(keychain.SecClassGenericPassword)
    query.SetService(service)
    query.SetAccount(account)
    query.SetAccessGroup(accessGroup)
    query.SetMatchLimit(keychain.MatchLimitAll)
    query.SetReturnAttributes(true)
    results, err := keychain.QueryItem(query)
    if err != nil {
        // Error
    } else {
        for _, r := range results {
            fmt.Printf("%#v\n", r)
        }
    }
}
