package main

import "log"

func startServer() error {
	log.Println("Starting Server")
	return nil
}

func main() {
	err := startServer()
	if err != nil {
		log.Fatal(err)
	}

}
