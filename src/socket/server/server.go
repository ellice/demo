package main

import (
	"fmt"
	"net"
	"log"
	"bufio"
)


func handleConnction(conn net.Conn){
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {log.Fatal("get conn data err :",err)}
	fmt.Printf("%#v\n", data)
	fmt.Fprintf(conn, "hello client\n")
    conn.Close()
	
}


func main(){

ln,err := net.Listen("tcp","127.0.0.1:6110")
if err != nil {panic(err)}

for{
	conn,err := ln.Accept()
	if err != nil {log.Fatal("get client connction error ",err)}
	go handleConnction(conn)

}


}
