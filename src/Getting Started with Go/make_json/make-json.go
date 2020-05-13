package main

import (
    "encoding/json"
    "fmt"
    "bufio"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    ans := make(map[string]string)
    fmt.Print("Enter your name : ")
    scanner.Scan()
    name := scanner.Text()
    fmt.Print("Enter your address : ")
    scanner.Scan()
    addr := scanner.Text()
    ans["name"] = name
    ans["address"] = addr
    fin, err := json.Marshal(ans)
    if err == nil {
        fmt.Println(string(fin))
    }
}
