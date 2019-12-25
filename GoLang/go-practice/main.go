package main

import (
	"./formatting"
	"./scanning"
	"fmt"
)

func main() {
	formatting.OutputToStdout()
	filename, err := formatting.OutputToWriter()
	fmt.Println(filename, err)

	scanning.ScanWithSscan()
	scanning.ScanWithSscanf()
	scanning.ScanWithSscanln()

	scanning.ScanWithScan()
}
