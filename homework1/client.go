package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if argsLen := len(os.Args); argsLen != 6 {
		fmt.Println("usage: client ip port str1 str2 str3")
		os.Exit(1)
	}
	fmt.Println(os.Args[1] + ":" + os.Args[2])
	conn, err := net.Dial("tcp", os.Args[1]+":"+os.Args[2])
	if err != nil {
		panic(err)
	}
	for i := 3; i < 6; i++ {
		fmt.Println(os.Args[i])
		_, err := conn.Write([]byte(os.Args[i]))
		time.Sleep(time.Microsecond * 50)
		if err != nil {
			fmt.Println(err)
		}
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	conn.Close()
	fmt.Println(string(buf[:n]))
}
