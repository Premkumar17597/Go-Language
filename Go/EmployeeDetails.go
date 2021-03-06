package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"os"
	"regexp"
)

type Details struct {
	Name        string
	Id          int
	Designation string
	Email       string
	Phone       string
}

var Database = map[int]Details{}

var err error
var file *os.File
var EmployeeId []int
var RequiredInput int
var RequiredId int
var RequiredEmail string
var RequiredPhone, RequiredPhone1 string
var RequiredName string
var RequiredDesignation string
var RequiredDesignation1 string
var Data []Details

func remove(s []int, i int) []int {
	temp := s[:i-1]
	temp1 := s[i:]
	result := append(temp, temp1...)
	return result
}

func main() {
	menu()
}

func menu() {

	for {
		fmt.Println("********************************")

		fmt.Println("Enter the option")
		fmt.Println("1. To List out the Employee IDs")
		fmt.Println("2. To search Employee Details with ID")
		fmt.Println("3. To add employee Details")
		fmt.Println("4. To delete an Employee Details")
		fmt.Println("5. To commit the data")
		fmt.Println("6. To check Data")
		fmt.Println("7. To Exit")
		fmt.Scanln(&RequiredInput)

		switch RequiredInput {
		case 1:
			fmt.Println(EmployeeId)
		case 2:
			search()
		case 3:
			add()
		case 4:
			Delete()
		case 5:
			flush()
		case 6:
			openFile()
			defer file.Close()
		case 7:
			return
		default:
			fmt.Println("You have entered invalid option")
		}
	}
}

func Delete() {
	fmt.Println("Enter the Employee Id")
	fmt.Scanln(&RequiredId)
	_, ok := Database[RequiredId]
	if ok == false {
		fmt.Println("Employee details doesn't exist for entered Id")
	} else {
		delete(Database, RequiredId)
		count := 0
		for _, temp := range EmployeeId {
			if temp == RequiredId {
				fmt.Println("count :", count)
				EmployeeId = remove(EmployeeId, count+1)
			}
			count = count + 1
		}
		temp, ok := Database[RequiredId]
		if ok == false {
			fmt.Printf("Details of %v Id has been deleted succesfully\n", RequiredId)
		} else {
			fmt.Printf("Unabele to Delete %v Id's details\n", temp)
		}
	}

}

func search() {
	fmt.Println("Enter the Employee ID")
	fmt.Scanln(&RequiredId)
	_, ok := Database[RequiredId]
	if ok == false {
		fmt.Println("Employee details doesn't exist for entered Id")
	} else {
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Println("Name : ", Database[RequiredId].Name)
		fmt.Println("Employee ID : ", Database[RequiredId].Id)
		fmt.Println("Role : ", Database[RequiredId].Designation)
		fmt.Println("Email : ", Database[RequiredId].Email)
		fmt.Println("Phone : ", Database[RequiredId].Phone)
	}
}

func add() {
	fmt.Println("Enter the Employee Id :")
	fmt.Scanln(&RequiredId)
	_, ok := Database[RequiredId]
	if ok == true {
		fmt.Println("Employee details are already exist for entered Id")
	} else {
		fmt.Println("Enter the Employee Name :")
		fmt.Scanln(&RequiredName)
		fmt.Println("Enter the Employee Designation :")
		fmt.Scanln(&RequiredDesignation, &RequiredDesignation1)
		fmt.Println("Enter the Employee Email :")
		fmt.Scanln(&RequiredEmail)
		fmt.Println("Enter the Employee Phone :")
		fmt.Scanln(&RequiredPhone, &RequiredPhone1)
		RequiredPhone = RequiredPhone + " " + RequiredPhone1
		b := PhoneValidater(RequiredPhone)
		if b == false {
			fmt.Println("you have entered Invalid number")
		} else {
			Database[RequiredId] = Details{RequiredName, RequiredId, RequiredDesignation + " " + RequiredDesignation1, RequiredEmail, RequiredPhone}
			EmployeeId = append(EmployeeId, RequiredId)
			temp, ok := Database[RequiredId]
			RequiredPhone1 = ""
			if ok == true {
				fmt.Printf("Details of %v Id has been added succesfully\n", RequiredId)
			} else {
				fmt.Printf("Unabele to add %v Id's details\n", temp)
			}
		}
	}

}

func flush() {
	file, err = os.OpenFile("EmployeeData.db", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("File is failed to open")
		return
	} else {
		Conn, err := net.Dial("tcp", "localhost:7200")
		if err != nil {
			fmt.Println("Error in net Dial", err)
			return
		}
		ServEncoder := gob.NewEncoder(Conn)
		serverEnCount := 0
		for i, _ := range EmployeeId {
			serverEnCount++
			temp := EmployeeId[i]
			err = ServEncoder.Encode(Database[temp])
			fmt.Println(temp)
			fmt.Println(Database[temp])
			if err != nil {
				fmt.Println("Server Encoding failed")
				return
			}
		}
		fmt.Println("ServerEnCount : ", serverEnCount)
		fmt.Println("Server Encoding completed!!!!!")
		Conn.Close()
	}

}

func openFile() {
	file, err = os.OpenFile("EmployeeData.db", os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println("Failed to open the file")
		return
	} else {
		loadDataDecoder()
	}

}

func loadDataDecoder() {
	decoder := gob.NewDecoder(file)
	var buff Details
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
	fmt.Println("Loader count : ", count)
	DisplayCount := 0
	for i, _ := range Data {
		DisplayCount++
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Println("Name : ", Data[i].Name)
		fmt.Println("Employee ID : ", Data[i].Id)
		fmt.Println("Role : ", Data[i].Designation)
		fmt.Println("Email : ", Data[i].Email)
		fmt.Println("Phone : ", Data[i].Phone)
	}
	Data = nil
	fmt.Println("Data Display count : ", DisplayCount)

}

func PhoneValidater(s string) bool {
	if len(s) >= 10 {
		r := regexp.MustCompile(`\d{10}`)
		t := regexp.MustCompile(`\d{5} \d{5}`)
		if r.MatchString(s) == true || t.MatchString(s) == true {
			return true
		} else {
			return false
		}
	}
	return false
}
