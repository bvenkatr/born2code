package json_example

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

var obj map[string]interface{}

func UpdateJson() *map[string]interface{} {
	jsonFile, err := os.Open("./json-example/lms_default.env.json")
	if err != nil {
		log.Fatal("Got error while reading file", err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	if err := json.Unmarshal(byteValue, &obj); err != nil {
		log.Fatal("Got error while Unmarshalling json data", err)
	}
	obj["PLATFORM_NAME"] = "gurukulam.school.com"
	fmt.Println(obj["PLATFORM_NAME"])
	return &obj
}

func WriteToNewFile(data *map[string]interface{}) {
	byteValue, err := json.MarshalIndent(&data, "", "  ")
	if err != nil {
		fmt.Println("Got error while marshalling json data", err)
	}
	file, err := os.Create("./json-example/lms.env.json")
	if err != nil {
		log.Fatal("Got error while create a file", err)
	}
	n, err := file.Write(byteValue)
	if err != nil {
		log.Fatal("Got error while writing data to file", err)
	}
	fmt.Println("Successfully updated...", n)
}