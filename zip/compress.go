package zip

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func CompressFile(filename string) error {
	// Open the original file
	originalFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func(originalFile *os.File) {
		err := originalFile.Close()
		if err != nil {
		}
	}(originalFile)

	// Create a new gzip-compressed file
	compressedFile, err := os.Create(fmt.Sprintf("%s.gz", filename))
	if err != nil {
		return err
	}
	defer func(compressedFile *os.File) {
		err := compressedFile.Close()
		if err != nil {

		}
	}(compressedFile)

	// Create a gzip writer
	gzipWriter := gzip.NewWriter(compressedFile)
	defer func(gzipWriter *gzip.Writer) {
		err := gzipWriter.Close()
		if err != nil {

		}
	}(gzipWriter)

	// Copy the original file contents to the gzip-compressed file
	_, err = io.Copy(gzipWriter, originalFile)
	if err != nil {
		return err
	}

	// Flush the gzip writer to ensure all data has been written
	err = gzipWriter.Flush()
	if err != nil {
		return err
	}

	// Close the gzip writer
	err = gzipWriter.Close()
	if err != nil {
		return err
	}

	return nil
}
