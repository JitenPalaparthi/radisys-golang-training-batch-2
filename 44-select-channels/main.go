package main

import (
	"fmt"
	"time"
)

func main() {
	s1 := server1()
	s2 := server2()
	s3 := server3()
	t1 := time.After(time.Millisecond * 10)
loop:
	select {
	case server := <-s1:
		process(server)
	case server := <-s2:
		process(server)
	case server := <-s3:
		process(server)
	case t := <-t1:
		fmt.Println("Timedout-->", t)
	default:
		fmt.Println("Default value no server selected or timedout happened")
		goto loop
	}
}

func process(server Server) {
	fmt.Println("Name:", server.Name)
	fmt.Println("Conn:", server.Conn)
}

type Server struct {
	Name, Conn string
}

func server1() <-chan Server {
	ch := make(chan Server)
	go func() {
		time.Sleep(time.Millisecond * 10)
		ch <- Server{Name: "Server1", Conn: "10.14.15.16:8080"}
		close(ch)
	}()
	return ch
}

func server2() <-chan Server {
	ch := make(chan Server)
	go func() {
		time.Sleep(time.Millisecond * 10)
		ch <- Server{Name: "Server2", Conn: "10.14.15.17:8080"}
		close(ch)
	}()
	return ch
}
func server3() <-chan Server {
	ch := make(chan Server)
	go func() {
		time.Sleep(time.Millisecond * 10)
		ch <- Server{Name: "Server3", Conn: "10.14.15.18:8080"}
		close(ch)
	}()
	return ch
}
