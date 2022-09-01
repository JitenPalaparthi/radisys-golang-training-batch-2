package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello World!")
	// })

	// if err := http.ListenAndServe(":50082", nil); err != nil {
	// 	log.Fatalln(err)
	// }
	fs := new(FileServer)
	fs.FileName = "temp.fsd"

	fmt.Fprintln(fs, "Hello World!", time.Now())
}

type FileServer struct {
	FileName string
}

func (fs *FileServer) Write(bytes []byte) (int, error) {
	if f, err := os.Create(fs.FileName); err != nil {
		return 0, err
	} else {
		return f.WriteString(string(bytes))
	}
}
