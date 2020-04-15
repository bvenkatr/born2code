package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create and add some files to the archive
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	var files = []struct {
		Name, Body string
	}{
		{
			Name: "readme.txt",
			Body: "This archive contains some text files.",
		},
		{
			Name: "gopher.txt",
			Body: "Gopher names:\nGeorge\nGeoffrey\nGonzo",
		},
		{
			Name: "todo.txt",
			Body: "Get animal handling license.",
		},
	}

	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}

		if x := tw.WriteHeader(hdr); x != nil {
			log.Fatal(x)
		}

		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	// Open and iterate through the files in the archive
	tr := tar.NewReader(&buf)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // end of archive
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}
