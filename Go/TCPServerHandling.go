package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"os"
)

var Data []string
var err error
var file *os.File

func main() {
	ls, err := net.Listen("tcp", "localhost:7200")
	if err != nil {
		fmt.Println("unable to listen the port id", err)
		return
	}
	fmt.Println("listening for client connections")
	conn, err := ls.Accept()
	if err != nil {
		fmt.Println("data is not receieved")
		return
	}
	ServerDecoder := gob.NewDecoder(conn)
	file, err = os.OpenFile("Note.db", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Failed to open the file")
		return
	}
	var Servbuff string
	for {
		err := ServerDecoder.Decode(&Servbuff)
		if err == io.EOF {
			break
		}
		Data = append(Data, Servbuff)
	}
	encoder := gob.NewEncoder(file)
	for i, _ := range Data {
		err = encoder.Encode(Data[i])
		if err != nil {
			fmt.Println("Encoding failed")
			return
		}
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Println(Data[i])
	}
	Data = nil
	fmt.Println("Encoding completed!!!!!")
	file.Close()
}
