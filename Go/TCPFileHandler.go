package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

var Option int
var Text string
var err error
var file *os.File
var Data []string

func main() {
	for {
		fmt.Println("Enter the option")
		fmt.Println("1. Add Data to file")
		fmt.Println("2. Replace string in the file")
		fmt.Println("3. Check Data from File")
		fmt.Println("4. Exit")
		fmt.Scanln(&Option)
		switch Option {
		case 1:
			uploadData()

		case 2:
			c := decodeData()
			editData(c)

		case 3:
			b := decodeData()
			DisplayData(b)

		case 4:
			return
		}

	}
	defer file.Close()
}

func uploadData() {
	fmt.Println("Write the paragraph")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	fmt.Println("Captured : ", line)
	flush(line)
}

func editData(str []string) {
	var ReWord, ReWord2, Result string
	fmt.Println("Enter a word to replace")
	fmt.Scanln(&ReWord)
	fmt.Println("Enter a word to replace : ", ReWord)
	fmt.Scanln(&ReWord2)
	//fmt.Println(str)
	for i, _ := range str {
		n := strings.Count(str[i], ReWord)
		fmt.Println("Number of strings matched : ", n)
		str[i] = strings.Replace(str[i], ReWord, ReWord2, -1)
		Result = str[i]
		fmt.Println("Modified Data : ", str[i])
	}
	flush(Result)
	Data = nil

}

func decodeData() []string {
	file, err = os.OpenFile("Note.db", os.O_RDONLY, 0755)
	var a []string = nil
	if err != nil {
		fmt.Println("Failed to open the file")
		return a
	} else {
		a = loadDataDecoder()
		return a
	}

}

func loadDataDecoder() []string {
	decoder := gob.NewDecoder(file)
	var buff string
	for {
		err := decoder.Decode(&buff)
		if err == io.EOF {
			//fmt.Println("Reached the End!!!")
			break
		}

		Data = append(Data, buff)
	}
	return Data

}

func DisplayData(str []string) {
	fmt.Printf("length of Data %v and capacity of Data %v\n", len(str), cap(str))
	for i, _ := range str {
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Println(str[i])
	}
	Data = nil

}

func flush(s string) {

	Conn, err := net.Dial("tcp", "localhost:7200")
	if err != nil {
		fmt.Println("Error in net Dial", err)
		return
	}
	ServEncoder := gob.NewEncoder(Conn)
	err = ServEncoder.Encode(s)
	if err != nil {
		fmt.Println("Server Encoding failed")
		return
	}
	fmt.Println("Server Encoding completed!!!!!")
	Conn.Close()
}
