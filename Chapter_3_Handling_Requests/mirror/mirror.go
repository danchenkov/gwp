package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func handleConnection(c net.Conn) {

	log.Printf("Client %v connected.", c.RemoteAddr())

	// stuff to do... like read data from client, process it, write back to client
	// see what you can do with (c net.Conn) at
	// http://golang.org/pkg/net/#Conn

	buffer := make([]byte, 4096)
	w := bufio.NewWriter(os.Stdout)

	for {
		n, err := c.Read(buffer)
		if err != nil || n == 0 {
			c.Close()
			break
		}
		log.Printf("%s", buffer)
		n, err = w.Write(buffer[0:n])
		if err != nil {
			c.Close()
			break
		}
	}
	log.Printf("Connection from %v closed.", c.RemoteAddr())
}

func main() {
	ln, err := net.Listen("tcp", ":143")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server up and listening on port 143")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
