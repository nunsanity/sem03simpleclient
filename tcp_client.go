package main

import (
	"net"
	"log"
	"os"
	"github.com/nunsanity/is105sem03/mycrypt"
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.2:8080")
	if err != nil {
		log.Fatal(err)
	}
    
	message :=[]rune{}
	if len(os.Args)>1{
		message = []rune(os.Args[1])
	} else {
		log.Fatal("Ingen melding gitt")
	}
    
	kryptertMelding, err := mycrypt.Krypter(message, 4)
	if err != nil {
      log.Fatal(err)
    }

    log.Println("Kryptert melding: ", string(kryptertMelding))

    _, err = conn.Write([]byte(string(kryptertMelding)))
	if err != nil {
		log.Fatal(err)
		return
	}
  
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	} 

    kryptertRespons :=[]rune(string(buf[:n]))
	if len(kryptertRespons)>0{
	dekryptertRespons, err := mycrypt.Krypter(kryptertRespons, -4)
	if err != nil{
		log.Fatal(err)
	}
	
	log.Println("respons fra proxy: ", string(dekryptertRespons))
    } else {
	log.Println("Ingen respons fra proxy")
  }
}
