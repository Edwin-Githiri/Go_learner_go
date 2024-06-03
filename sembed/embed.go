// This ia a test sample of interface embedding that is a great feature of golang
// you'll grow to like this language as you master it more.
// Your welcome:
package main

import (
	"fmt"
	"strconv"
)

type School struct {
	// faculty string
	Student
}
type Department int

// departments are coded 1-5

const (
	DeptOrganicChemistry Department = iota + 1
	DeptInorganicChemistry
	DeptAnimalScience
	DeptGeneralAnatomy
	DeptDrugTrial
)

// func (t T) String() string
// this is a stringer interface
// func (d Department) String() string {
// 	switch d {
// 	case DeptOrganicChemistry:
// 		return "Department of Organic Chemistry"
// 	case DeptInorganicChemistry:
// 		return "Department of Inorganic Chemistry"
// 	case DeptAnimalScience:
// 		return "Department of Animal Sciences"
// 	case DeptGeneralAnatomy:
// 		return "Department of General Anatomy"
// 	case DeptDrugTrial:
// 		return "Department of Drug Trial and testing"
// 	default:
// 		panic("That department doesnt exist within this school")
// 	}
// }

type Student struct {
	name      string
	Classroom string
	RegNo     int
}

func DeptFinder() {
	var adminNo int
	fmt.Println("Enter your Registration Number: ")
	fmt.Scan(&adminNo)
	sReg := strconv.Itoa(adminNo)
	if len(sReg) != 5 {
		fmt.Printf("Invalid Entry\nEnter a 5 digit number: ")
		fmt.Scan(&adminNo)
	}
	if len(sReg) == 5 {
		fmt.Println("Your entry looks good :)")
	}
	code := (sReg[0])
	strconv.Atoi(string(code))
	deptcode := code - 48
	switch deptcode {
	case 1:
		fmt.Println("You belong to the Department of Organic Chemistry")
	case 2:
		fmt.Println("You belong to the Department of Inorganic Chemistry")
	case 3:
		fmt.Println("You belong to the Department of Animal Sciences")
	case 4:
		fmt.Println("You belong to the Department of General Anatomy")
	case 5:
		fmt.Println("You belong to the Department of Drug Trial and testing")
	default:
		panic("That department doesnt exist within this school")
	}
}

func main() {
	jon := Student{
		name:      "Jonstone Merci",
		Classroom: "sc 8",
		// the regNo first digit represents the department code
		RegNo: 34521,
	}
	fmt.Println("Are you a student in this school?")
	var responce string
	fmt.Scan(&responce)
	if responce == "yes" {
		fmt.Println("What Your registration No: ")
		fmt.Scan(&jon.RegNo)
		if jon.RegNo == 34521 {
			fmt.Printf("Okay, %s.\nSorry for that.\nJust doing my job.\nGo ahead to your class %s ", jon.name, jon.Classroom)
		}
	} else {
		fmt.Println("What's your Registration Number?")
		DeptFinder()
	}
}
