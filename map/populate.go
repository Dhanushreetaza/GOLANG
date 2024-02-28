package main

import "fmt"

func main() {
	studentPhones := map[string]string{
		"john":  "202-555-0179",
		"emma":  "03.37.77.63.06",
		"peter": "03489940240",
	}

	courseAvailability := map[int]bool{
		101: true,
		202: false,
		303: true,
	}

	studentContacts := map[string][]string{
		"john":  {"202-555-0179"},
		"emma":  {"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
		"peter": {"03489940240", "03489900120"},
	}

	schoolBasket := map[int]map[int]int{
		1001: {101: 4, 303: 2},
		1002: {303: 5, 404: 20},
		1003: {},
	}

	who, phone := "emma", "N/A"
	if v, ok := studentPhones[who]; ok {
		phone = v
	}
	fmt.Printf("%s's phone number: %s\n", who, phone)

	id, status := 202, "available"
	if !courseAvailability[id] {
		status = "not " + status
	}
	fmt.Printf("Course ID #%d is %s\n", id, status)

	who, phone = "peter", "N/A"
	if contacts := studentContacts[who]; len(contacts) >= 2 {
		phone = contacts[1]
	}
	fmt.Printf("%s's 2nd contact number: %s\n", who, phone)

	studentID, courseID := 1002, 303
	fmt.Printf("Student #%d is going to enroll in %d from Course ID #%d.\n", studentID, schoolBasket[studentID][courseID], courseID)
}
