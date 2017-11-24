package main

import (
	"fmt"
	"log"
	"net"

	"strings"
	"time"
	"try/countryip"
)

const PORT = "8080"

func main() {
	fmt.Println("Server starts... on port: " + PORT)
	ln, e := net.Listen("tcp", ":"+PORT)
	if e != nil {
		log.Println(e)
	}
	ch := Conns(ln)
	for {
		go handler(<-ch)
	}
}

func Conns(ln net.Listener) chan net.Conn {
	con := make(chan net.Conn)
	go func() {
		for {
			cl, e := ln.Accept()
			if e != nil {
				log.Println("Error connection client ip:", cl.RemoteAddr())
				continue
			}
			con <- cl
		}
	}()
	return con
}

func handler(client net.Conn) {
	ip := client.RemoteAddr().String()
	i := strings.Split(ip, ":")
	if len(i) > 1 && len(i[0]) > 3 {
		country, err := GetCountry(i[0])
		if err != nil {
			client.Write([]byte("Api error"))

		}
		client.Write(country)
	} else {
		client.Write([]byte("not found"))
	}
	client.SetDeadline(time.Now())
}

func GetCountry(ip string) ([]byte, error) {
	s, e := countryip.GetCountry(ip)
	return []byte(s), e
}
