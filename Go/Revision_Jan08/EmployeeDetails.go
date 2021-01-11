package main

import (
	"bufio"
	"fmt"
	"os"
)

type Details struct {
	Name        string
	Id          int
	Designation string
	email       string
	phone       string
}

func remove(s []int, i int) []int {
	temp := s[:i-1]
	temp1 := s[i:]
	result := append(temp, temp1...)
	return result
}

func main() {
	var EmployeeId = []int{1001, 1002}
	var RequiredInput int
	var RequiredId int
	var RequiredEmail string
	var RequiredPhone string
	var Database = map[int]Details{
		1001: {"Ramesh", 1001, "Software Developer", "samesh@gmail.com", "+91-9865321475"},
		1002: {"Suresh", 1002, "Lead Engineer", "suresh@gmail.com", "+91-9994465232"},
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("********************************")

		fmt.Println("Enter the option")
		fmt.Println("1. To List out the Employee IDs")
		fmt.Println("2. To search Employee Details with ID")
		fmt.Println("3. To add employee Details")
		fmt.Println("4. To delete an Employee Details")
		fmt.Println("5. To Exit")
		fmt.Scanln(&RequiredInput)

		switch RequiredInput {
		case 1:
			fmt.Println(EmployeeId)
		case 2:
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
		case 3:
			fmt.Println("Enter the Employee Id :")
			fmt.Scanln(&RequiredId)
			_, ok := Database[RequiredId]
			if ok == true {
				fmt.Println("Employee details are already exist for entered Id")
			} else {
				fmt.Println("Enter the Employee Name :")
				RequiredName, _ := reader.ReadString('\n')
				fmt.Println("Enter the Employee Designation :")
				RequiredDesignation, _ := reader.ReadString('\n')
				fmt.Println("Enter the Employee Email :")
				fmt.Scanln(&RequiredEmail)
				fmt.Println("Enter the Employee Phone :")
				fmt.Scanln(&RequiredPhone)
				Database[RequiredId] = Details{RequiredName, RequiredId, RequiredDesignation, RequiredEmail, RequiredPhone}
				EmployeeId = append(EmployeeId, RequiredId)
				temp, ok := Database[RequiredId]
				if ok == true {
					fmt.Printf("Details of %v Id has been added succesfully\n", RequiredId)
				} else {
					fmt.Printf("Unabele to add %v Id's details\n", temp)
				}
			}
		case 4:
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

		case 5:
			return
		default:
			fmt.Println("You have entered invalid option")
		}
	}
}
