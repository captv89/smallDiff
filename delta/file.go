package delta

import (
	"io"
	"log"
	"os"
)

// saveToFile saves the data to a file
func saveToFile(data []byte, filename string) error {
	// create a file
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	//log.Printf("Filename: %v Data length: %v", filename, len(data))
	//log.Println(data)

	// write to the file
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// readFromFile reads the data from a file
func readFromFile(filename string) ([]byte, error) {
	// open the file
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	// convert the file to a byte array
	delta, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return delta, nil
}
