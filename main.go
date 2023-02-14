package main

import (
	"flag"
	"log"
	"smallDiff/delta"
)

var action string
var oldFile string
var newFile string
var patchFile string
var isCompressed bool

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	parseFlags()
}

func parseFlags() {
	flag.StringVar(&action, "action", "", "Enter action: delta or apply")
	flag.StringVar(&newFile, "new", "", "Enter the full new file name")
	flag.StringVar(&oldFile, "old", "", "Enter the full old file name")
	flag.StringVar(&patchFile, "patch", "", "Enter the full patch file name")
	flag.BoolVar(&isCompressed, "compressed", true, "Enter true if the patch file is compressed")
	flag.Parse()

	// Check if action is delta or apply
	if action == "delta" {
		log.Printf("delta")
		err := delta.CreateDelta(oldFile, newFile, patchFile, isCompressed)
		if err != nil {
			log.Panic(err)
		}
	} else if action == "apply" {
		log.Printf("apply")
		err := delta.ApplyDelta(oldFile, patchFile, newFile, isCompressed)
		if err != nil {
			log.Panic(err)
		}
	} else {
		log.Fatalf("action must be delta or apply")
	}
}
