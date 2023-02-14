package delta

import (
	"github.com/gabstv/go-bsdiff/pkg/bsdiff"
	"github.com/gabstv/go-bsdiff/pkg/bspatch"
	"log"
	"smallDiff/zip"
)

// CreateDelta creates a delta file from the old file and the new file
func CreateDelta(oldFile, newFile, patchFile string, isCompressed bool) error {

	// Read the old file
	oldByte, err := readFromFile(oldFile)
	if err != nil {
		log.Panic(err)
	}

	// Read the new file
	newByte, err := readFromFile(newFile)
	if err != nil {
		log.Panic(err)
	}

	// generate a BSDIFF4 patch
	patch, err := bsdiff.Bytes(oldByte, newByte)
	if err != nil {
		log.Panic(err)
	}
	//log.Printf("patch: %v", patch)

	// Save to file
	err = saveToFile(patch, patchFile)
	if err != nil {
		log.Panic(err)
	}

	if isCompressed {
		// Compress the patch with compress package gzip
		err = zip.CompressFile(patchFile)
		if err != nil {
			return err
		}
		return err
	} else {
		return err
	}
}

// ApplyDelta applies the delta to the old file
func ApplyDelta(oldFile, patchFile, newFile string, isCompressed bool) error {

	// Decompress the delta file
	if isCompressed {
		err := zip.DecompressFile(patchFile)
		if err != nil {
			return err
		}
		patchFile = patchFile[:len(patchFile)-3]
	}

	// Read the old file
	oldByte, err := readFromFile(oldFile)
	if err != nil {
		return err
	}

	// Read the delta file
	deltaByte, err := readFromFile(patchFile)
	if err != nil {
		return err
	}

	// Apply the delta to the old file
	newByte, err := bspatch.Bytes(oldByte, deltaByte)
	if err != nil {
		return err
	}

	// Save the new file
	err = saveToFile(newByte, newFile)
	if err != nil {
		return err
	}
	return nil
}
