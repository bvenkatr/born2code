// Write to zip archive and read from zip archive

package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main()  {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive
	zw := zip.NewWriter(buf)

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	} {
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	for _, file := range files {
		f, err := zw.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}

		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}

	zr, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(len(buf.Bytes())))
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range zr.File {
		fmt.Printf("Contents of %s\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(os.Stdout, rc)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}
