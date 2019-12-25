package scanning

import (
	"fmt"
)

func ScanWithSscan() {
	var d1, d2 int
	var s1 string
	fmt.Printf("Before scanning: %d, %d, %s\n", d1, d2, s1)

	if _, err := fmt.Sscan("2 5\n7", &d1, &d2, &s1); err != nil {
		fmt.Println("Got error while scanning values ", err)
		return
	}
	fmt.Printf("After scanning: %d, %d, %s\n", d1, d2, s1)
}

func ScanWithSscanf() {
	var d1, d2 int
	var s1 string
	fmt.Printf("Before scanning: %d, %d, %s\n", d1, d2, s1)

	if _, err := fmt.Sscanf("1 3 and 9 are my values", "%d %d and %s", &d1, &d2, &s1); err != nil {
		fmt.Println("Gor error while scanning values ", err)
		return
	}

	fmt.Printf("After scanning: %d, %d, %s\n", d1, d2, s1)
}

func ScanWithSscanln() {
	var d1, d2 int
	var s1 string

	fmt.Printf("Before scanning: %d, %d, %s\n", d1, d2, s1)
	if _, err := fmt.Sscanln("3 5 8\n 4 9", &d1, &d2, &s1); err != nil {
		fmt.Println("Got error while scanning values ", err)
		return
	}
	fmt.Printf("After scanning: %d, %d, %s\n", d1, d2, s1)
}

func ScanWithScan() {
	var s1, s2, s3 string
	fmt.Printf("Before scanning: %s, %s, %s\n", s1, s2, s3)

	if _, err := fmt.Scan(&s1, &s2, &s3); err != nil {
		fmt.Println("Got error while scanning values ", err)
		return
	}
	fmt.Printf("After scanning: %s, %s, %s\n", s1, s2, s3)
}
