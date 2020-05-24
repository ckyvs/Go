package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)
 
type name struct {
    f_name string
    l_name string
}
 
func main() {
    names :=  make([]name, 0, 1)
    var temp_name name
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Enter file name : ")
    scanner.Scan()
    file := scanner.Text()
    namesFile, err := os.Open(file)
    if err == nil {
        scanFile := bufio.NewScanner(namesFile)
        for scanFile.Scan() {
            temp := strings.Split(scanFile.Text(), " ")
            temp_name.f_name = temp[0]
            temp_name.l_name = temp[1]
            names = append(names, temp_name)
        }
    }
    for _, list := range names {
        fmt.Println(list.f_name, list.l_name)
    }
}
