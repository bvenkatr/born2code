package main

import (
	// "fmt"
	// "io/ioutil"
	// "log"

	// simplepb "./src/simple"
	test "./test"
	// "github.com/golang/protobuf/jsonpb"
	// "github.com/golang/protobuf/proto"
)

func main() {
	// sm := doSimple()

	// writeToFile("simple.bin", sm)

	// sm2 := &simplepb.SimpleMessage{}
	// readFromFile("simple.bin", sm2)
	// fmt.Println(sm2)

	test.OutputToStdout()
}

// func toJSON(pb proto.Message) {
// 	jsonpb.Marshaller()
// }

// func writeToFile(filename string, pb proto.Message) error {
// 	out, err := proto.Marshal(pb)
// 	if err != nil {
// 		log.Fatalln("Can't serialize to bytes", err)
// 		return err
// 	}

// 	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
// 		log.Fatalln("Can't write to file", err)
// 		return err
// 	}

// 	fmt.Println("Data has been written.")
// 	return nil
// }

// func readFromFile(filename string, pb proto.Message) error {
// 	in, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		log.Fatalln("Couldn't read from file", err)
// 		return err
// 	}

// 	if err := proto.Unmarshal(in, pb); err != nil {
// 		log.Fatalln("Can't unserialize the data", err)
// 		return err
// 	}

// 	return nil
// }

// func doSimple() *simplepb.SimpleMessage {
// 	sm := simplepb.SimpleMessage{
// 		Id:         123,
// 		IsSimple:   true,
// 		Name:       "My Simple",
// 		SampleList: []int32{1, 4, 7, 8},
// 	}
// 	return &sm
// }
