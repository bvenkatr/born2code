package main

import (
	// "./formatting"
	// "./scanning"
	// "fmt"
	// strings_and_bytes "./string-and-bytes"
	//"./strconv-package"
	"./json-example"
)

func main() {
	// formatting.OutputToStdout()
	// filename, err := formatting.OutputToWriter()
	// fmt.Println(filename, err)

	// scanning.ScanWithSscan()
	// scanning.ScanWithSscanf()
	// scanning.ScanWithSscanln()

	// scanning.ScanWithScan()

	// strings_and_bytes.Contains()
	// strings_and_bytes.TrimFunc()
	// strings_and_bytes.Replacer()
	// strings_and_bytes.Join()
	// strings_and_bytes.Reader()

	//strconv_examples.StringToInt()

	data := json_example.UpdateJson()
	json_example.WriteToNewFile(data)
}
