package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	PORT string
)

func main() {

	flag.StringVar(&PORT, "port", "50082", "--port 50082")
	flag.Parse()

	if port := os.Getenv("PORT"); port != "" {
		PORT = port
	}

	// fmt.Println(len(os.Args))
	// fmt.Println(os.Args)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})
	http.HandleFunc("/ping", ping)

	http.HandleFunc("/greet", greet("Hello Folks!"))

	fmt.Println("Server started and listening on port:", PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatalln(err)
	}
}

// what 0.0.0.0 -->
// 127.0.0.1 -->lo
// wls
// bluetooth
// cni
// flannel
// eth0

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func greet(msg string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, msg)
	}
}
