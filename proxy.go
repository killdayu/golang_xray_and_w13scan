package main

import (
	"io"
	"log"
	"net"
)

//src 源 dst目标
func handle(src net.Conn, target string) {
	dst, err := net.Dial("tcp", target)
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}
	defer dst.Close()
	// Run in goroutine to prevent io.Copy from blocking
	go func() { //采用并发，保证建立双向通道。
		// Copy our source's output to the destination
		if _, err := io.Copy(dst, src); err != nil { //将目标回复输出给源
			log.Fatalln(err)
		}
	}()
	// Copy our destination's output back to our source
	if _, err := io.Copy(src, dst); err != nil { //将源的请求发送给目标，这里不知道为什么和作者的注释反了，不过无所谓，负负得正。
		log.Fatalln(err)
	}
}

func proxy_two() {
	// Listen on local port 80
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn, "127.0.0.1:7777")
		go handle(conn, "127.0.0.1:7778")

	}
}
