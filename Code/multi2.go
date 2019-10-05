package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

type ClientJob struct {
	name string
	conn net.Conn
}

func generateResponse(clientJobs chan ClientJob) {
	for {
		//wait for the next job to come off the queue
		clientJob := <-clientJobs

		for start := time.Now(); time.Now().Sub(start) < time.Second; {
		}

		clientJob.conn.Write([]byte("Hello, " + clientJob.name))
	}

}

func main() {
	clientJobs := make(chan ClientJob)
	go generateResponse(clientJobs)

	ln, err := net.Listen("tcp", ":8080")
	check(err, "Server is ready.")

	for {
		conn, err := ln.Accept()
		check(err, "Accepted connection.")

		go func() {
			buf := bufio.NewReader(conn)

			for {
				name, err := buf.ReadString('\n')

				if err != nil {
					fmt.Printf("client disconnected. \n")
					break
				}

				clientJobs <- ClientJob{name, conn}
			}
		}()
	}
}