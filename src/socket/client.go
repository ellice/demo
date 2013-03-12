package main
    
import (
    "fmt"
    "net"
    "bufio"
)
    
func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:6110")
    if err != nil {
        panic(err)
    }
    
    fmt.Fprintf(conn, "hello server\n")
    
    data, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("%#v\n", data)
}
