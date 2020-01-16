package main

import (
	// "./formatting"
	// "./scanning"
	// "fmt"
	// strings_and_bytes "./string-and-bytes"
	//"./strconv-package"
	//"./json-example"
	//"./regex-example"
	//"./error-examples"
	//"./web-server"
	//"./checsum_and_content_hash"
	//"./io-and-ioutil"
	"./goroutines"
	"fmt"
)

var version string

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

	//data := json_example.UpdateJson()
	//json_example.WriteToNewFile(data)

	//regex_example.Match()

	//course_backups.CourseBackups()

	//error_examples.CustomErrors()

	//web_server.WebServer()
	//web_server.CustomServeMux()

	//checsum_and_content_hash.CheckSum()
	//io_and_ioutil.StringNewReader()

	for i := 0; i < 10; i++ {
		go goroutines.Test_Go_Routines(i)
	}
	var str string
	if _, err := fmt.Scanf("%s", &str); err != nil {
		fmt.Println(err)
	}
}
