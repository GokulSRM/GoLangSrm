package main

import (
	// "flag"
	"fmt"
	"strconv"
)

var student [3][11]string

// "Roll Number", "Name", "Subject 1", "Subject 2", "Subject 3", "Subject 4", "Subject 5", "Total", "Average", "Number of failed Subjects", "Grade", "Failed marks"
var quest = [10]string{"Roll Number", "Name", "Subject 1", "Subject 2", "Subject 3", "Subject 4", "Subject 5", "Total", "Average", "Number of failed Subjects"}
var rw = 0

func addstd() {
	fmt.Println("Enter 3 Students details!.")
	for i := 0; i < 3; i++ {
		sum := 0
		fail := 0
		for j := 0; j < 7; j++ {
			fmt.Printf(quest[j])
			fmt.Scanln(&student[i][j])
			if j > 1 && j < 7 {
				m, _ := strconv.Atoi(student[i][j])
				if m < 0 {
					fmt.Println("Mark should not be negative!")
					j--
				} else if m > 100 {
					fmt.Println("Mark should not be greater than 100!")
					j--
				} else if student[i][j] == "" {
					fmt.Println("Mark should not be null!")
					j--
				} else {
					// m1, _ := strconv.Atoi(student[i][j])
					sum = sum + m
					if m < 50 {
						student[i][j] = student[i][j] + " fail"
						fail++
					}
				}

			}
			t := strconv.Itoa(sum)
			student[i][7] = t
			avrg := sum / 5
			avg := strconv.Itoa(avrg)
			student[i][8] = avg
			f := strconv.Itoa(fail)
			student[i][9] = f
		}
	}
}

func viewstd() {

	if student[0][0] == "" {
		fmt.Println("No records found!")
	} else {
		fmt.Println("Student List")
		for i := 0; i < 3; i++ {
			fmt.Println("count:", i+1)
			for j := 0; j < 10; j++ {
				fmt.Println(quest[j], " : ", student[i][j])
			}
		}
	}
}

//re:=regex.MustComplie("[a-zA-Z]$")
//if re.MatchString(student[i][2])
func searchstd(a int) {
	// flag.Parse()
	lc := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			v, _ := strconv.Atoi(student[i][j])
			if v == a {
				fmt.Println(student[i])
				lc++
			}
		}
	}
	if lc == 0 {
		fmt.Println("No records found!")
	}
}

func main() {

	var choice int
	var roll int
	for {
		fmt.Println("1.Add Student  2.View Students  3.Search  4.Exit")
		fmt.Println("Enter your choice")
		fmt.Scanln(&choice)

		if choice == 1 {
			addstd()
		} else if choice == 2 {
			viewstd()
		} else if choice == 3 {
			fmt.Println("Enter the roll number")
			fmt.Scanln(&roll)
			if roll < 0 {
				fmt.Println("Roll number cannot be negative!")
			} else {
				searchstd(roll)
			}
		} else if choice == 4 {
			fmt.Println("Thanks for using student portal!")
			break
		} else {
			fmt.Println("Invalid Option!")
		}
	}

}
