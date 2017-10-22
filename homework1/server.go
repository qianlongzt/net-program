package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	port := 0
	var err error
	if len(os.Args) == 2 {
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("usage: " + os.Args[0] + " port")
			os.Exit(1)
		}
	} else {
		fmt.Println("usage: " + os.Args[0] + " port")
		os.Exit(1)
	}
	listener, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port))
	add := listener.Addr()
	fmt.Println("listening on " + add.Network() + " " + add.String())
	if err != nil {
		panic(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1024)
	arr := make([]string, 3)
	for i := 0; i < 3; i++ {
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		arr[i] = string(buf[:n])
		fmt.Println(string(buf[:n]))
	}
	conn.Write([]byte(arr[2] + arr[1] + arr[0]))
	conn.Close()
}
