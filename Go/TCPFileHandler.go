package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
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
			editData()

		case 3:
			decodeData()
			defer file.Close()

		case 4:
			return
		}

	}
}

func uploadData() {
	fmt.Println("Write the paragraph")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	fmt.Println("Captured : ", line)
	file, err = os.OpenFile("Note.db", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("File is failed to open")
		return
	}
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(line)
	if err != nil {
		fmt.Println("Encoding failed")
		return
	}

	fmt.Println("Encoding completed!!!!!")
	file.Close()

}

func editData() {
	fmt.Println("Edit Data fun called")

}

func decodeData() {
	file, err = os.OpenFile("Note.db", os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println("Failed to open the file")
		return
	} else {
		loadDataDecoder()
	}

}

func loadDataDecoder() {
	decoder := gob.NewDecoder(file)
	var buff string
	count := 0
	for {
		count++
		err := decoder.Decode(&buff)
		if err == io.EOF {
			//fmt.Println("Reached the End!!!")
			break
		}

		Data = append(Data, buff)
		//fmt.Println(Data)
	}
	/*err := decoder.Decode(&buff)
	if err == io.EOF {
		fmt.Println("Reached the End!!!")
	}*/
	/*fmt.Println("Captured : ")
	fmt.Println("                ", buff)
	buff = "" */

	fmt.Println("Loader count : ", count)
	fmt.Printf("length of Data %v and capacity of Data %v\n", len(Data), cap(Data))
	DisplayCount := 0
	for i, _ := range Data {
		DisplayCount++
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Println(Data[i])
		n := strings.Count(Data[i], "prem")
		fmt.Println("Number of strings matched : ", n)
		Data[i] = strings.Replace(Data[i], "prem", "arun", -1)
		fmt.Println("Modified Data : ", Data[i])

	}
	Data = nil
	fmt.Println("Data Display count : ", DisplayCount)

}
