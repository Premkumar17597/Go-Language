package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
)

type Details struct {
	Name        string
	Id          int
	Designation string
	email       string
	phone       string
}

var Database = map[int]Details{

	1001: {"Ramesh", 1001, "Software Developer", "samesh@gmail.com", "+91-9865321475"},
	1002: {"Suresh", 1002, "Lead Engineer", "suresh@gmail.com", "+91-9994465232"},
}

var err error
var file *os.File
var EmployeeId = []int{1001, 1002}
var RequiredInput int
var RequiredId int
var RequiredEmail string
var RequiredPhone string
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
		fmt.Println("Email : ", Database[RequiredId].email)
		fmt.Println("Phone : ", Database[RequiredId].phone)
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
		fmt.Scanln(&RequiredPhone)
		Database[RequiredId] = Details{RequiredName, RequiredId, RequiredDesignation + " " + RequiredDesignation1, RequiredEmail, RequiredPhone}
		EmployeeId = append(EmployeeId, RequiredId)
		temp, ok := Database[RequiredId]
		if ok == true {
			fmt.Printf("Details of %v Id has been added succesfully\n", RequiredId)
		} else {
			fmt.Printf("Unabele to add %v Id's details\n", temp)
		}
	}
}

func flush() {
	file, err = os.OpenFile("EmployeeData.db", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("File is failed to open")
		return
	} else {
		loadDataEncoder()
	}
}

func loadDataEncoder() {
	encoder := gob.NewEncoder(file)
	for i, _ := range EmployeeId {
		temp := EmployeeId[i]
		err = encoder.Encode(Database[temp])
		fmt.Println(temp)
		fmt.Println(Database[temp])
		if err != nil {
			fmt.Println("Encoding failed")
			return
		}
	}
	fmt.Println("Encoding completed!!!!!")
	file.Close()
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
	for {
		err := decoder.Decode(&buff)
		if err == io.EOF {
			fmt.Println("Reached the End!!!")
			break
		}
		if err != nil {
			fmt.Println("Error in Decoding the file")
			break
		}

		Data = append(Data, buff)
	}

	for i, _ := range Data {
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Println("Name : ", Data[i].Name)
		fmt.Println("Employee ID : ", Data[i].Id)
		fmt.Println("Role : ", Data[i].Designation)
		fmt.Println("Email : ", Data[i].email)
		fmt.Println("Phone : ", Data[i].phone)
	}

}
