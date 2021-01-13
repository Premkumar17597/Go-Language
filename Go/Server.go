package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"os"
)

type Details struct {
	Name        string
	Id          int
	Designation string
	Email       string
	Phone       string
}

var Data []Details
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
	file, err = os.OpenFile("EmployeeData.db", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("File is failed to open")
		return
	}
	var Servbuff Details
	ServDeCount := 0
	for {
		ServDeCount++
		err := ServerDecoder.Decode(&Servbuff)
		if err == io.EOF {
			//fmt.Println("Reached the End!!!")
			break
		}
		Data = append(Data, Servbuff)
		//fmt.Println(Data)
	}
	fmt.Println("Server decode count", ServDeCount)
	encoder := gob.NewEncoder(file)
	ServDisplayCount := 0
	for i, _ := range Data {
		ServDisplayCount++
		err = encoder.Encode(Data[i])
		if err != nil {
			fmt.Println("Encoding failed")
			return
		}
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Println("Name : ", Data[i].Name)
		fmt.Println("Employee ID : ", Data[i].Id)
		fmt.Println("Role : ", Data[i].Designation)
		fmt.Println("Email : ", Data[i].Email)
		fmt.Println("Phone : ", Data[i].Phone)
	}
	fmt.Println("Display Count", ServDisplayCount)
	fmt.Println("Encoding completed!!!!!")
	file.Close()
}
