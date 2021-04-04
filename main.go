package main

import {
	"log"
	"net"
}

func handleConnection(conn *.net.Conn) error {

}

func startServer() error {
	log.Println("Starting Server")

	ln. err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection", err)
			continue
		}
		go handleConnection(conn)
	}
	return nil
}

func main() {
	err := startServer()
	if err != nil {
		log.Fatal(err)
	}

}
