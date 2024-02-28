package main

import "fmt"

func main() {
	var (
		studentPhones map[string]string

		courseEnrollment map[int]bool

		studentContacts map[string][]string

		classroomAttendance map[int]map[int]int
	)

	fmt.Printf("studentPhones        : %#v\n", studentPhones)
	fmt.Printf("courseEnrollment     : %#v\n", courseEnrollment)
	fmt.Printf("studentContacts      : %#v\n", studentContacts)
	fmt.Printf("classroomAttendance  : %#v\n", classroomAttendance)
}
