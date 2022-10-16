package main

import (
	"flag"
	"log"
	"os"
)

func ls() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		println(f.Name())
	}
}

func main() {
	fl := flag.Bool("help", false, "Help command invocation")
	flag.Parse()

	if *fl {
		println("TODO: Implement help flag cmd")
	} else {
		println("No specific flag provided, using defaults\n")
		ls()
	}

}
