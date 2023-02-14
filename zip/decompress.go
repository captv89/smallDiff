package zip

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

// DecompressFile decompresses the file
func DecompressFile(filename string) error {
	// Open the gzip-compressed file
	compressedFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func(compressedFile *os.File) {
		err := compressedFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(compressedFile)

	// Create a gzip reader
	gzipReader, err := gzip.NewReader(compressedFile)
	if err != nil {
		return err
	}
	defer func(gzipReader *gzip.Reader) {
		err := gzipReader.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(gzipReader)

	// Create a new file to decompress the gzip-compressed file contents into
	decompressedFile, err := os.Create(fmt.Sprintf("%s", filename[:len(filename)-3]))
	if err != nil {
		return err
	}
	defer func(decompressedFile *os.File) {
		err := decompressedFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(decompressedFile)

	// Copy the gzip-compressed file contents to the new file
	_, err = io.Copy(decompressedFile, gzipReader)
	if err != nil {
		return err
	}

	// Flush the gzip reader to ensure all data has been written
	err = gzipReader.Close()
	if err != nil {
		return err
	}

	// Close the gzip reader
	err = gzipReader.Close()
	if err != nil {
		return err
	}

	return nil
}
