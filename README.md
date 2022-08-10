## Go tooling

- To run program

    ```go run hello.go```

- To build and generate binary/exe 

    ```go build hello.go```

- To build and generate desired binary/exe name

    ```go build -o app.exe hello.go```

- To check work directory of go run

    ```go run --work hello.go```

- To compile and link using compiler tool.
    1. It generates object file.
    2. The linker uses that object file (hello.o) and generates exe/binary file. 

    ```go tool compile hello.go``` 

    ```go tool link -o app.exe hello.o```

- To get list of OS/Arch go supports for cross compilation

    ```go tool list dist```

- To cross compile and generate binary

    ```GOARCH=arm64 GOOS=darwin go build -o app hello.go```

- To know what is happening when run/build gofile

    ```go run -x hello.go```
- Just compile do not produce any output

    ```go build -n hello.go```

- To getch assembly language instructions

    ```go build -gcflags="-S" hello.go```

- To install go binaries in GOBIN directory

    ```go install hello.go```

- To install from remote location

    ```go install github.com/JitenPalaparthi/urllinter```

