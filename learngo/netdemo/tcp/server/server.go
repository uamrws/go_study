package main

import (
	"io"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":8099")
	if err != nil {
		log.Fatalf("faild to listen port: 8099, because: %v", err)
		return
	}
	defer server.Close()
	for {
		con, err := server.Accept()
		if err != nil {
			log.Printf("faild to connect client, because: %v", err)
			continue
		} else {
			log.Printf("succed to connect client: %v", con)
		}
		go func(con net.Conn) {
			defer con.Close()
			for {
				buf := make([]byte, 1024)
				n, err := con.Read(buf)
				if err == io.EOF {
					return
				}
				if err != nil {
					log.Printf("faild to read from client, because: %v", err)
				}
				log.Printf("%#v", string(buf[:n]))
				log.Println(string(buf[:n]) == string("exit"))
				if string(buf[:n]) == "exit" {
					break
				}
			}

		}(con)
	}
}
