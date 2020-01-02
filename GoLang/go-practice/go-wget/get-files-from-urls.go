package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func writeToFile(fileName string, url string) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Got error while preparing request %v", err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("request err %v", err)
	}

	defer response.Body.Close()

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Got error while creating file: %v", err)
	}
	defer file.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Got error while parsing response body %v", err)
	}
	_, err = file.Write(responseData)
	if err != nil {
		log.Fatalf("Got error while writing to file: %v", err)
	}
}

type fileInfo struct {
	name string
	link string
}

func main() {
	fileNames := []fileInfo{
		{
			name: "my-bondage-and-my-freedom.txt",
			link: "https://www.gutenberg.org/files/202/202-0.txt",
		},
		{
			name: "the-iliad.txt",
			link: "https://www.gutenberg.org/ebooks/16452.txt.utf-8",
		},
		{
			name: "pride-and-prejudice.txt",
			link: "https://www.gutenberg.org/ebooks/42671.txt.utf-8",
		}, {
			name: "the-underground-railroad.txt",
			link: "https://www.gutenberg.org/ebooks/15263.txt.utf-8",
		}, {
			name: "the-republic.txt",
			link: "https://www.gutenberg.org/ebooks/55201.txt.utf-8",
		},
		{
			name: "moby-dick.txt",
			link: "https://www.gutenberg.org/ebooks/15.txt.utf-8",
		},
		{name: "meditations.txt",
			link: "https://www.gutenberg.org/ebooks/2680.txt.utf-8",
		},
		{
			name: "war-and-peace.txt",
			link: "https://www.gutenberg.org/files/2600/2600-0.txt",
		},
	}

	for _, v := range fileNames {
		writeToFile(v.name, v.link)
	}
}
