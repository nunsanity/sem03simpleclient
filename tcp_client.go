package main

import (
	"net"
	"log"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.3:8080")
	if err != nil {
		log.Fatal(err)
	}
    
	log.Println("os.Args[1] = ", os.Args[1])

 	_, err = conn.Write([]byte(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	} 
	response := string(buf[:n])
	log.Printf("reply from proxy: %s", response)
}
