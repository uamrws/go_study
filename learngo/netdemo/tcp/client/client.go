package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	con, err := net.Dial("tcp", "127.0.0.1:8099")
	if err != nil {
		log.Fatalf("failed connect to server, because:%v", err)
	}
	defer con.Close()
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		if _, err := con.Write(scan.Bytes()); err != nil {

			log.Printf("failed to send to server, because:%#v", err)
		}

	}
}
